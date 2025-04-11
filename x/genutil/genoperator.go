package genutil

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	genutiltypes "github.com/dittonetwork/kepler/x/genutil/types"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"
)

func AddGenesisOperators(cdc codec.Codec, operators []restakingtypes.Validator, genesisFileURL string) error {
	appState, appGenesis, err := genutiltypes.GenesisStateFromGenFile(genesisFileURL)
	if err != nil {
		return fmt.Errorf("failed to unmarshal genesis state: %w", err)
	}

	authGenState := authtypes.GetGenesisStateFromAppState(cdc, appState)
	restakingGenState := restakingtypes.GetGenesisStateFromAppState(cdc, appState)

	accounts, err := authtypes.UnpackAccounts(authGenState.Accounts)
	if err != nil {
		return fmt.Errorf("failed to unpack accounts: %w", err)
	}

	for _, operator := range operators {
		var foundAcc bool

		// check if operator has account
		for _, acc := range accounts {
			var ok bool
			if ok, err = restakingtypes.IsOperatorAddress(acc.GetPubKey(), operator.OperatorAddress); err != nil {
				return fmt.Errorf("failed to check operator address: %w", err)
			} else if ok {
				foundAcc = true
			}
		}

		if !foundAcc {
			return fmt.Errorf("operator %s does not have an account", operator.OperatorAddress)
		}

		restakingGenState.Validators = append(restakingGenState.Validators, restakingtypes.Validator{
			OperatorAddress: operator.OperatorAddress,
			ConsensusPubkey: operator.ConsensusPubkey,
			Status:          operator.Status,
			Protocol:        operator.Protocol,
			VotingPower:     operator.VotingPower,
			IsEmergency:     operator.IsEmergency,
		})
	}

	if appState[restakingtypes.ModuleName], err = cdc.MarshalJSON(&restakingGenState); err != nil {
		return fmt.Errorf("failed to marshal restaking genesis state: %w", err)
	}

	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		return fmt.Errorf("failed to marshal app genesis state: %w", err)
	}

	appGenesis.AppState = appStateJSON
	return ExportGenesisFile(appGenesis, genesisFileURL)
}

package genutil

import (
	"encoding/json"
	"fmt"

	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/privval"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	genutiltypes "github.com/dittonetwork/kepler/x/genutil/types"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"
)

func AddGenesisOperators(cdc codec.Codec, operators []restakingtypes.Operator, genesisFileURL string) error {
	appState, appGenesis, err := genutiltypes.GenesisStateFromGenFile(genesisFileURL)
	if err != nil {
		return fmt.Errorf("failed to unmarshal genesis state: %w", err)
	}

	restakingGenState := restakingtypes.GetGenesisStateFromAppState(cdc, appState)

	for _, operator := range operators {
		restakingGenState.PendingValidators = append(restakingGenState.PendingValidators, restakingtypes.Operator{
			Address:         operator.Address,
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

func GetValidatorPubKey(config *cfg.Config) (cryptotypes.PubKey, error) {
	filePV := privval.LoadFilePV(config.PrivValidatorKeyFile(), config.PrivValidatorStateFile())

	tmValPubKey, err := filePV.GetPubKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get pubkey from filePV: %w", err)
	}

	valPubKey, err := cryptocodec.FromCmtPubKeyInterface(tmValPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to convert pubkey: %w", err)
	}

	return valPubKey, nil
}

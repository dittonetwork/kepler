package genutil

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
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

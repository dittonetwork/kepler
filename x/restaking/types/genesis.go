package types

import (
	"encoding/json"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		LastUpdate: UpdateInfo{
			EpochNum:    1,
			Timestamp:   time.Time{},
			BlockHeight: 1,
			BlockHash:   "",
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}

// GetGenesisStateFromAppState returns x/auth GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}

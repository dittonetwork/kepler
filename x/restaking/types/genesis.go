package types

import (
	"time"
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

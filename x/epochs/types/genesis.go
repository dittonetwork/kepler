package types

import (
	"errors"
	"time"
)

// this line is used by starport scaffolding # genesis/types/import.

// DefaultIndex is the default global index.
const (
	DailyEpoch  time.Duration = time.Hour * 24
	WeeklyEpoch time.Duration = time.Hour * 24 * 7
)

func NewGenesisState(epochs []EpochInfo) *GenesisState {
	return &GenesisState{Epochs: epochs, Params: Params{}}
}

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	epochs := []EpochInfo{
		NewGenesisEpochInfo("day", DailyEpoch), // alphabetical order
		NewGenesisEpochInfo("hour", time.Hour),
		NewGenesisEpochInfo("minute", time.Minute),
		NewGenesisEpochInfo("week", WeeklyEpoch),
	}
	return NewGenesisState(epochs)
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}

func (epoch EpochInfo) Validate() error {
	if epoch.Identifier == "" {
		return errors.New("epoch identifier should NOT be empty")
	}

	if epoch.Duration == 0 {
		return errors.New("epoch duration should NOT be zero")
	}

	if epoch.CurrentEpoch < 0 {
		return errors.New("current epoch should NOT be negative")
	}

	if epoch.CurrentEpochStartHeight < 0 {
		return errors.New("current epoch start height should NOT be negative")
	}

	return nil
}

func NewGenesisEpochInfo(identifier string, duration time.Duration) EpochInfo {
	return EpochInfo{
		Identifier:              identifier,
		StartTime:               time.Time{},
		Duration:                duration,
		CurrentEpoch:            0,
		CurrentEpochStartHeight: 0,
		CurrentEpochStartTime:   time.Time{},
		EpochCountingStarted:    false,
	}
}

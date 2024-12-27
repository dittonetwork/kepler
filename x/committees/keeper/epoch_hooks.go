package keeper

import (
	"context"

	epochstypes "cosmossdk.io/x/epochs/types"
)

var _ epochstypes.EpochHooks = (*Keeper)(nil)

// AfterEpochEnd is called when epoch is going to be ended, epochNumber is the number of epoch that is ending.
func (k Keeper) AfterEpochEnd(ctx context.Context, epochIdentifier string, epochNumber int64) error {
	return nil
}

// BeforeEpochStart is called when epoch is going to be started, epochNumber is the number of epoch that is starting.
func (k Keeper) BeforeEpochStart(ctx context.Context, epochIdentifier string, epochNumber int64) error {
	return nil
}

package keeper

import (
	"context"
	"kepler/x/epochs/types"
)

// Hooks returns the epoch module's epoch hooks.
func (k Keeper) Hooks() types.EpochHooks {
	if k.hooks == nil {
		// return a no-op implementation if no hooks are set
		return types.MultiEpochHooks{}
	}

	return k.hooks
}

// AfterEpochEnd is called when epoch is going to be ended, epochNumber is the number of epoch that is ending.
func (k Keeper) AfterEpochEnd(ctx context.Context, id string, number int64) error {
	return k.Hooks().AfterEpochEnd(ctx, id, number)
}

// BeforeEpochStart is called when epoch is going to be started, epochNumber is the number of epoch that is starting.
func (k Keeper) BeforeEpochStart(ctx context.Context, id string, number int64) error {
	return k.Hooks().BeforeEpochStart(ctx, id, number)
}

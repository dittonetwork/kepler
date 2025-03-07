package keeper

import (
	"context"
	"fmt"

	"github.com/dittonetwork/kepler/x/epochs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetEpochInfo returns the epoch info for the given id.
func (k Keeper) GetEpochInfo(ctx context.Context, id string) (types.EpochInfo, error) {
	return k.EpochInfo.Get(ctx, id)
}

// AddEpochInfo adds a new epoch info. Will return an error if the epoch fails validation,
// or re-uses an existing identifier.
// This method also sets the start time if left unset, and sets the epoch start height.
func (k Keeper) AddEpochInfo(ctx context.Context, epoch types.EpochInfo) error {
	if err := epoch.Validate(); err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	isExist, err := k.EpochInfo.Has(ctx, epoch.Identifier)
	if err != nil {
		return err
	}

	if isExist {
		return fmt.Errorf("epoch with identifier %s already exists", epoch.Identifier)
	}

	// Initialize empty and default epoch values
	if epoch.StartTime.IsZero() {
		epoch.StartTime = sdkCtx.HeaderInfo().Time
	}

	if epoch.CurrentEpochStartHeight == 0 {
		epoch.CurrentEpochStartHeight = sdkCtx.HeaderInfo().Height
	}

	return k.EpochInfo.Set(ctx, epoch.Identifier, epoch)
}

// AllEpochInfos iterate through epochs to return all epochs info.
func (k Keeper) AllEpochInfos(ctx context.Context) ([]types.EpochInfo, error) {
	epochs := make([]types.EpochInfo, 0)
	err := k.EpochInfo.Walk(
		ctx,
		nil,
		func(_ string, value types.EpochInfo) (bool, error) {
			epochs = append(epochs, value)
			return false, nil
		},
	)

	return epochs, err
}

// NumBlocksSinceEpochStart returns the number of blocks since the epoch started.
// if the epoch started on block N, then calling this during block N (after BeforeEpochStart)
// would return 0.
// Calling it any point in block N+1 (assuming the epoch doesn't increment) would return 1.
func (k Keeper) NumBlocksSinceEpochStart(ctx context.Context, identifier string) (int64, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	epoch, err := k.EpochInfo.Get(ctx, identifier)
	if err != nil {
		return 0, fmt.Errorf("epoch with identifier %s not found", identifier)
	}
	return sdkCtx.HeaderInfo().Height - epoch.CurrentEpochStartHeight, nil
}

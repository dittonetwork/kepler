package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) UpdateFinalizedBlock(ctx context.Context) error {
	if !k.SyncNeeded(ctx) {
		k.Logger().Debug("no need to update finalized block yet")
		return nil
	}

	finalizedBlock, err := k.getFinalizedBlock(ctx)
	if err != nil {
		return err
	}

	k.SetFinalizedBlockInfo(ctx, finalizedBlock)

	return nil
}

func (k Keeper) SyncNeeded(ctx context.Context) bool {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return (sdkCtx.BlockHeader().Height % SYNC_PERIOD) == 0
}

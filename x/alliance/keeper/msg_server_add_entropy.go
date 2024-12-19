package keeper

import (
	"context"

	"kepler/x/alliance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddEntropy(goCtx context.Context, msg *types.MsgAddEntropy) (*types.MsgAddEntropyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	prevSharedEntropy, _ := k.GetSharedEntropy(goCtx)
	k.Keeper.SetSharedEntropy(goCtx,
		types.SharedEntropy{Entropy: mergeEntropies(uint64(prevSharedEntropy.Entropy), uint64(msg.Entropy))},
	)

	return &types.MsgAddEntropyResponse{}, nil
}

func mergeEntropies(a, b uint64) uint64 {
	return a ^ b
}

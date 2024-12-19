package keeper

import (
	"context"

	"kepler/x/alliance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddEntropy(goCtx context.Context, msg *types.MsgAddEntropy) (*types.MsgAddEntropyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddEntropyResponse{}, nil
}

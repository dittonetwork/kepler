package keeper

import (
	"context"

	"kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitJobResult(goCtx context.Context, msg *types.MsgSubmitJobResult) (*types.MsgSubmitJobResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitJobResultResponse{}, nil
}

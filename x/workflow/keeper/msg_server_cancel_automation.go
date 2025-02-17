package keeper

import (
	"context"
	"kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelAutomation(
	goCtx context.Context,
	msg *types.MsgCancelAutomation,
) (*types.MsgCancelAutomationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.CancelAutomation(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.MsgCancelAutomationResponse{}, nil
}

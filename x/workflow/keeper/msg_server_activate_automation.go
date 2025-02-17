package keeper

import (
	"context"
	"kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ActivateAutomation(
	goCtx context.Context,
	msg *types.MsgActivateAutomation,
) (*types.MsgActivateAutomationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.ActivateAutomation(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.MsgActivateAutomationResponse{}, nil
}

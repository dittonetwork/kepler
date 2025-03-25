package keeper

import (
	"context"
	"fmt"

	"github.com/dittonetwork/kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) ActivateAutomation(
	goCtx context.Context,
	msg *types.MsgActivateAutomation,
) (*types.MsgActivateAutomationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	automation, err := k.GetAutomation(ctx, msg.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("failed to get automation: %s", err))
	}

	if automation.Creator != msg.Creator {
		return nil, status.Error(codes.PermissionDenied, "automation created by other address")
	}

	err = k.Keeper.ActivateAutomation(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.MsgActivateAutomationResponse{}, nil
}

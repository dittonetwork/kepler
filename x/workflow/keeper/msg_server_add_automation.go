package keeper

import (
	"context"
	"fmt"

	"github.com/dittonetwork/kepler/x/workflow/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAutomation(
	goCtx context.Context,
	msg *types.MsgAddAutomation,
) (*types.MsgAddAutomationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, err := k.GetNextAutomationID(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get next automation ID: %s", err))
	}

	automation := &types.Automation{
		Id:       id,
		Triggers: msg.GetTriggers(),
		UserOp:   msg.GetUserOp(),
		Status:   types.AutomationStatus_AUTOMATION_STATUS_ACTIVE,
		Creator:  msg.Creator,
	}

	err = k.InsertAutomation(ctx, *automation)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to set automation: %s", err))
	}

	return &types.MsgAddAutomationResponse{
		Id: id,
	}, nil
}

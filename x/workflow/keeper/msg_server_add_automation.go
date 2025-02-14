package keeper

import (
	"context"
	"fmt"
	"kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAutomation(
	goCtx context.Context,
	msg *types.MsgAddAutomation,
) (*types.MsgAddAutomationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, err := k.GetNextAutomationID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get next automation ID: %w", err)
	}

	automation := &types.Automation{
		Id:       id,
		Triggers: msg.GetTriggers(),
		Actions:  msg.GetActions(),
		ExpireAt: msg.GetExpireAt(),
		Status:   types.AutomationStatus_AUTOMATION_STATUS_ACTIVE,
	}

	err = k.InsertAutomation(ctx, *automation)
	if err != nil {
		return nil, fmt.Errorf("failed to set automation: %w", err)
	}

	return &types.MsgAddAutomationResponse{
		Id: id,
	}, nil
}

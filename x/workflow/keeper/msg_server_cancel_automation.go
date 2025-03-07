package keeper

import (
	"context"
	"fmt"

	"github.com/dittonetwork/kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) CancelAutomation(
	goCtx context.Context,
	msg *types.MsgCancelAutomation,
) (*types.MsgCancelAutomationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	automation, err := k.Keeper.GetAutomation(ctx, msg.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("failed to get automation: %s", err))
	}

	switch {
	case msg.Committee != nil:
		var signValid bool

		signValid, err = k.CommitteeKeeper.CanBeSigned(
			ctx,
			msg.Committee.CommitteeId,
			msg.Committee.ChainId,
			msg.Committee.Signs,
			msg.Committee.Payload,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to check committee signs: %w", err)
		}

		if !signValid {
			return nil, status.Error(codes.PermissionDenied, "committee signs invalid")
		}
	case msg.Creator != "":
		if automation.Creator != msg.Creator {
			return nil, status.Error(codes.PermissionDenied, "automation created by other address")
		}
	default:
		return nil, status.Error(codes.InvalidArgument, "either committee or creator must be set")
	}

	err = k.Keeper.CancelAutomation(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.MsgCancelAutomationResponse{}, nil
}

package keeper

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dittonetwork/kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitJobResult(
	goCtx context.Context,
	msg *types.MsgSubmitJobResult,
) (*types.MsgSubmitJobResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	automation, err := k.GetAutomation(ctx, msg.AutomationId)
	if err != nil {
		if errors.Is(err, types.ErrAutomationNotFound) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("automation not found: %d", msg.AutomationId))
		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get automation: %s", err))
	}

	if automation.Status != types.AutomationStatus_AUTOMATION_STATUS_ACTIVE {
		return nil, status.Error(codes.FailedPrecondition, "automation is not active")
	}

	err = k.JobKeeper.CreateJob(
		ctx,
		0,
		msg.GetCommitteeId(),
		msg.GetChainId(),
		msg.GetAutomationId(),
		msg.GetTxHash(),
		msg.GetCreator(),
		msg.GetCreatedAt(),
		msg.GetExecutedAt(),
		msg.GetSignedAt(),
		msg.GetSigns(),
		msg.GetPayload(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create job: %s", err))
	}

	return &types.MsgSubmitJobResultResponse{}, nil
}

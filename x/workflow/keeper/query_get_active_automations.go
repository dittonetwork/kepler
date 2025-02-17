package keeper

import (
	"context"

	"kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k BaseKeeper) GetActiveAutomations(
	goCtx context.Context,
	req *types.QueryGetActiveAutomationsRequest,
) (*types.QueryGetActiveAutomationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	automations, err := k.FindActiveAutomations(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetActiveAutomationsResponse{
		ActiveAutomations: automations,
	}, nil
}

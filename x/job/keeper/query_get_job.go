package keeper

import (
	"context"

	"github.com/dittonetwork/kepler/x/job/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetJob(goCtx context.Context, req *types.QueryGetJobRequest) (*types.QueryGetJobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	job, b, err := k.GetJobByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetJobResponse{
		Job: &job,
	}, nil
}

package keeper

import (
	"context"

	"kepler/x/alliance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QuorumParams(goCtx context.Context, req *types.QueryGetQuorumParamsRequest) (*types.QueryGetQuorumParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetQuorumParams(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetQuorumParamsResponse{QuorumParams: val}, nil
}

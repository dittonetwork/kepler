package keeper

import (
	"context"

	"kepler/x/beacon/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FinalizedBlockInfo(goCtx context.Context, req *types.QueryGetFinalizedBlockInfoRequest) (*types.QueryGetFinalizedBlockInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetFinalizedBlockInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFinalizedBlockInfoResponse{FinalizedBlockInfo: val}, nil
}

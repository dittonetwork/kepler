package keeper

import (
	"context"

	"kepler/x/symbiotic/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContractAddress(goCtx context.Context, req *types.QueryGetContractAddressRequest) (*types.QueryGetContractAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetContractAddress(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetContractAddressResponse{ContractAddress: val}, nil
}

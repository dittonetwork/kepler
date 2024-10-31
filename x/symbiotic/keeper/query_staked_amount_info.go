package keeper

import (
	"context"

	"kepler/x/symbiotic/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakedAmountInfoAll(ctx context.Context, req *types.QueryAllStakedAmountInfoRequest) (*types.QueryAllStakedAmountInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stakedAmountInfos []types.StakedAmountInfo

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	stakedAmountInfoStore := prefix.NewStore(store, types.KeyPrefix(types.StakedAmountInfoKeyPrefix))

	pageRes, err := query.Paginate(stakedAmountInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var stakedAmountInfo types.StakedAmountInfo
		if err := k.cdc.Unmarshal(value, &stakedAmountInfo); err != nil {
			return err
		}

		stakedAmountInfos = append(stakedAmountInfos, stakedAmountInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStakedAmountInfoResponse{StakedAmountInfo: stakedAmountInfos, Pagination: pageRes}, nil
}

func (k Keeper) StakedAmountInfo(ctx context.Context, req *types.QueryGetStakedAmountInfoRequest) (*types.QueryGetStakedAmountInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetStakedAmountInfo(
		ctx,
		req.EthereumAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStakedAmountInfoResponse{StakedAmountInfo: val}, nil
}

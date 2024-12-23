package keeper

import (
	"context"

	"kepler/x/alliance/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AlliancesTimelineAll(ctx context.Context, req *types.QueryAllAlliancesTimelineRequest) (*types.QueryAllAlliancesTimelineResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var alliancesTimelines []types.AlliancesTimeline

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	alliancesTimelineStore := prefix.NewStore(store, types.KeyPrefix(types.AlliancesTimelineKey))

	pageRes, err := query.Paginate(alliancesTimelineStore, req.Pagination, func(key []byte, value []byte) error {
		var alliancesTimeline types.AlliancesTimeline
		if err := k.cdc.Unmarshal(value, &alliancesTimeline); err != nil {
			return err
		}

		alliancesTimelines = append(alliancesTimelines, alliancesTimeline)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAlliancesTimelineResponse{AlliancesTimeline: alliancesTimelines, Pagination: pageRes}, nil
}

func (k Keeper) AlliancesTimeline(ctx context.Context, req *types.QueryGetAlliancesTimelineRequest) (*types.QueryGetAlliancesTimelineResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	alliancesTimeline, found := k.GetAlliancesTimeline(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAlliancesTimelineResponse{AlliancesTimeline: alliancesTimeline}, nil
}

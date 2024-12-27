package keeper

import (
	"context"
	"errors"

	"kepler/x/committees/types"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListCommittees(ctx context.Context, req *types.QueryAllCommitteesRequest) (*types.QueryAllCommitteesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	committeess, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Committees,
		req.Pagination,
		func(_ uint64, value types.Committees) (types.Committees, error) {
			return value, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCommitteesResponse{Committees: committeess, Pagination: pageRes}, nil
}

func (q queryServer) GetCommittees(ctx context.Context, req *types.QueryGetCommitteesRequest) (*types.QueryGetCommitteesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	committees, err := q.k.Committees.Get(ctx, req.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetCommitteesResponse{Committees: committees}, nil
}

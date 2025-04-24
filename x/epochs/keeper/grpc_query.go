package keeper

import (
	"context"
	"errors"

	"github.com/dittonetwork/kepler/x/epochs/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/epochs keeper providing gRPC method
// handlers.
type Querier struct {
	keeper Keeper
}

// NewQuerier initializes new querier.
func NewQuerier(k Keeper) Querier {
	return Querier{keeper: k}
}

func (q Querier) EpochInfos(
	ctx context.Context,
	_ *types.QueryEpochsInfoRequest,
) (*types.QueryEpochInfosResponse, error) {
	epochs, err := q.keeper.AllEpochInfos(ctx)
	return &types.QueryEpochInfosResponse{Epochs: epochs}, err
}

// CurrentEpoch provides current epoch of specified identifier.
func (q Querier) CurrentEpoch(
	ctx context.Context,
	req *types.QueryCurrentEpochRequest,
) (*types.QueryCurrentEpochResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Identifier == "" {
		return nil, status.Error(codes.InvalidArgument, "identifier is empty")
	}

	info, err := q.keeper.EpochInfo.Get(ctx, req.Identifier)
	if err != nil {
		return nil, errors.New("not available identifier")
	}

	return &types.QueryCurrentEpochResponse{
		CurrentEpoch: info.CurrentEpoch,
	}, nil
}

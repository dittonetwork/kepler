package keeper

import (
	"context"

	"github.com/dittonetwork/kepler/x/executors/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetExecutors(_ context.Context, _ *types.QueryExecutorsRequest) (*types.QueryExecutorsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) GetEmergencyExecutors(
	_ context.Context,
	_ *types.QueryEmergencyExecutorsRequest,
) (*types.QueryEmergencyExecutorsResponse, error) {
	//TODO implement me
	panic("implement me")
}

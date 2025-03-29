package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/executors/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) GetEmergencyExecutors(
	ctx context.Context,
	_ *types.QueryEmergencyExecutorsRequest,
) (*types.QueryEmergencyExecutorsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	executors, err := q.Keeper.GetEmergencyExecutors(sdkCtx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var activeExecutors []*types.Executor
	for i := range executors {
		activeExecutors = append(activeExecutors, &executors[i])
	}

	return &types.QueryEmergencyExecutorsResponse{
		Executors: activeExecutors,
	}, nil
}

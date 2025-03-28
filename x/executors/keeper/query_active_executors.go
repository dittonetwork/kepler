package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/executors/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetActiveExecutors(
	ctx context.Context,
	_ *types.QueryActiveExecutorsRequest,
) (*types.QueryActiveExecutorsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	executors, err := k.getAllExecutors(sdkCtx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var activeExecutors []*types.Executor
	for i := range executors {
		if executors[i].GetIsActive() {
			activeExecutors = append(activeExecutors, &executors[i])
		}
	}

	return &types.QueryActiveExecutorsResponse{
		Executors: activeExecutors,
	}, nil
}

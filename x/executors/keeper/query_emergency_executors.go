package keeper

import (
	"context"
	"slices"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/executors/types"
	restakingTypes "github.com/dittonetwork/kepler/x/restaking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetEmergencyExecutors(
	ctx context.Context,
	_ *types.QueryEmergencyExecutorsRequest,
) (*types.QueryEmergencyExecutorsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	executors, err := k.getAllExecutors(sdkCtx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	activeExecutors := make([]*types.Executor, 0)
	for i := range executors {
		if executors[i].GetIsActive() {
			activeExecutors = append(activeExecutors, &executors[i])
		}
	}

	emergencyValidators := k.restaking.GetActiveEmergencyValidators(sdkCtx)
	res := make([]*types.Executor, 0, len(emergencyValidators))
	for _, activeExecutor := range activeExecutors {
		if slices.ContainsFunc(emergencyValidators, func(v restakingTypes.EmergencyValidator) bool {
			return v.Address.String() == activeExecutor.GetAddress()
		}) {
			res = append(res, activeExecutor)
		}
	}

	return &types.QueryEmergencyExecutorsResponse{
		Executors: res,
	}, nil
}

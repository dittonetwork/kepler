package keeper

import (
	"slices"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/executors/types"
	restaking "github.com/dittonetwork/kepler/x/restaking/types"
)

// GetEmergencyExecutors returns all active executors that are also emergency validators.
func (k Keeper) GetEmergencyExecutors(ctx sdk.Context) ([]types.Executor, error) {
	executors, err := k.getAllExecutors(ctx)
	if err != nil {
		k.Logger(ctx).With("error", err).Error("failed to get all executors")
		return nil, err
	}

	var activeExecutors []*types.Executor
	for i := range executors {
		if executors[i].GetIsActive() {
			activeExecutors = append(activeExecutors, &executors[i])
		}
	}

	emergencyValidators := k.restaking.GetActiveEmergencyValidators(ctx)
	res := make([]types.Executor, 0, len(emergencyValidators))
	for _, activeExecutor := range activeExecutors {
		if slices.ContainsFunc(emergencyValidators, func(v restaking.EmergencyValidator) bool {
			return v.Address.String() == activeExecutor.GetAddress()
		}) {
			res = append(res, *activeExecutor)
		}
	}

	return res, nil
}

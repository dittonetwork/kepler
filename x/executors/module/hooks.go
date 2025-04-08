package executors

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	restaking "github.com/dittonetwork/kepler/x/restaking/types"
)

var _ restaking.RestakingHooks = AppModule{}

func (am AppModule) BeforeValidatorBeginUnbonding(ctx context.Context, validator restaking.Validator) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	executors, err := am.keeper.GetExecutorsByOwnerAddress(sdkCtx, validator.OperatorAddress)
	if err != nil {
		return err
	}

	for _, executor := range executors {
		if !executor.GetIsActive() {
			continue
		}

		executor.IsActive = false
		if err = am.keeper.Executors.Set(ctx, executor.Address, executor); err != nil {
			return err
		}
	}

	return nil
}

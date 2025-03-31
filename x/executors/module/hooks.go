package executors

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	restakingTypes "github.com/dittonetwork/kepler/x/restaking/types"
)

var _ restakingTypes.RestakingHooks = AppModule{}

func (am AppModule) AfterValidatorBonded(_ context.Context, _ stakingtypes.ValidatorI) error {
	// in executors module we don't need to do anything after validator is bonded, so just return nil
	return nil
}

func (am AppModule) BeforeValidatorBeginUnbonding(ctx context.Context, validator stakingtypes.ValidatorI) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	executors, err := am.keeper.GetExecutorsByOwnerAddress(sdkCtx, validator.GetOperator())
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

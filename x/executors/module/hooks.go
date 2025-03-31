package executors

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	restakingTypes "github.com/dittonetwork/kepler/x/restaking/types"
)

var _ restakingTypes.RestakingHooks = AppModule{}

func (am AppModule) AfterValidatorBonded(ctx context.Context, validator stakingtypes.ValidatorI) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return am.keeper.CheckAndSetExecutorIsActive(sdkCtx, validator.GetOperator())
}

func (am AppModule) BeforeValidatorBeginUnbonding(ctx context.Context, validator stakingtypes.ValidatorI) error {
	executor, err := am.keeper.Executors.Get(ctx, validator.GetOperator())
	if err != nil {
		return err
	}

	executor.IsActive = false
	return am.keeper.Executors.Set(ctx, executor.Address, executor)
}

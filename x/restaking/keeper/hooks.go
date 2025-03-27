package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// Hooks returns the restaking module's restaking hooks.
func (k Keeper) Hooks() types.RestakingHooks {
	if k.hooks == nil {
		return types.MultiRestakingHooks{}
	}

	return k.hooks
}

// AfterValidatorBonded is called after a validator is bonded.
func (k Keeper) AfterValidatorBonded(ctx sdk.Context, validator stakingtypes.ValidatorI) error {
	return k.Hooks().AfterValidatorBonded(ctx, validator)
}

// BeforeValidatorBeginUnbonding is called before a validator begins unbonding.
func (k Keeper) BeforeValidatorBeginUnbonding(ctx sdk.Context, validator stakingtypes.ValidatorI) error {
	return k.Hooks().BeforeValidatorBeginUnbonding(ctx, validator)
}

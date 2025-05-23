package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// Hooks returns the restaking module's restaking hooks.
func (k Keeper) Hooks() types.RestakingHooks {
	if k.hooks == nil {
		return types.MultiRestakingHooks{}
	}

	return k.hooks
}

// BeforeValidatorBeginUnbonding is called before a validator begins unbonding.
func (k Keeper) BeforeValidatorBeginUnbonding(ctx sdk.Context, validator types.Validator) error {
	return k.Hooks().BeforeValidatorBeginUnbonding(ctx, validator)
}

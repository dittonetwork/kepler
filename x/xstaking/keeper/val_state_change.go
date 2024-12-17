package keeper

import (
	"context"
	"cosmossdk.io/core/appmodule"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"kepler/x/xstaking/types"
)

// BlockValidatorUpdates calculates the ValidatorUpdates for the current block
// Called in each EndBlock
func (k Keeper) BlockValidatorUpdates(ctx context.Context) ([]appmodule.ValidatorUpdate, error) {
	// Calculate validator set changes.
	//
	// NOTE: ApplyAndReturnValidatorSetUpdates has to come before
	// UnbondAllMatureValidatorQueue.
	// This fixes a bug when the unbonding period is instant (is the case in
	// some of the tests). The test expected the validator to be completely
	// unbonded after the Endblocker (go from Bonded -> Unbonding during
	// ApplyAndReturnValidatorSetUpdates and then Unbonding -> Unbonded during
	// UnbondAllMatureValidatorQueue).
	_, err := k.ApplyAndReturnValidatorSetUpdates(ctx)
	if err != nil {
		return nil, err
	}

	// @TODO implement BlockValidatorUpdates

	return nil, nil
}

func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx context.Context) ([]appmodule.ValidatorUpdate, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	maxValidators := params.MaxValidators

	iterator, err := k.ValidatorsPowerStoreIterator(ctx)
	if err != nil {
		return nil, err
	}

	defer iterator.Close()

	var updates []appmodule.ValidatorUpdate
	for count := 0; iterator.Valid() && count < int(maxValidators); iterator.Next() {
		addr := sdk.ValAddress(iterator.Value())
		validator, err := k.GetValidator(ctx, addr)
		if err != nil {
			return nil, err
		}

		if validator.Jailed {
			return nil, errors.New("should never retrieve a jailed validator from the power store")
		}

		power, err := k.PowerReduction(ctx)
		if err != nil {
			return nil, err
		}

		// if we get to a zero-power validator (which we don't bond),
		// there are no more possible bonded validators
		if validator.PotentialConsensusPower(power) == 0 {
			break
		}

		// apply the appropriate state change if necessary
		switch {
		case validator.IsUnbonding():
			validator, err = k.unbondingToBonded(ctx, validator)
			if err != nil {
				return nil, err
			}
		case validator.IsUnbonded():
			// @TODO implement change state to bonded
		case validator.IsBonding():
			// @TODO implement change state to bonded
		case validator.IsBonded():
			// no state change
		default:
			return nil, errors.New("unexpected validator state")
		}

		//
		//updates = append(updates, appmodule.ValidatorUpdate{
		//	PubKey: validator.GetConsPubKey(),
		//	Power:  validator.GetConsensusPower(),
		//})
		//count++
	}

	// @TODO implement ApplyAndReturnValidatorSetUpdates

	return updates, nil
}

func (k Keeper) unbondingToBonded(ctx context.Context, validator types.Validator) (types.Validator, error) {
	if !validator.IsUnbonding() {
		return types.Validator{}, fmt.Errorf("bad state transition unbondingToBonded %v", validator)
	}

	return k.bondValidator(ctx, validator)
}

func (k Keeper) bondValidator(ctx context.Context, validator types.Validator) (types.Validator, error) {
	if err := k.DeleteValidatorByPowerIndex(ctx, validator); err != nil {
		return validator, err
	}

	validator.UpdateStatus(types.Bonded)

	if err := k.SetValidator(ctx, validator); err != nil {
		return validator, err
	}

	if err := k.SetValidatorByPowerIndex(ctx, validator); err != nil {
		return validator, err
	}

	// @TODO: remove from waiting queue if present

	panic("implement hooks and return validator")
}

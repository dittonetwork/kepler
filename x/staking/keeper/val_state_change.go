package keeper

import (
	"bytes"
	"context"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/math"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/cosmos/gogoproto/types"
	"kepler/x/staking/types"
	"sort"
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
	validatorUpdates, err := k.ApplyAndReturnValidatorSetUpdates(ctx)
	if err != nil {
		return nil, err
	}

	err = k.UnbondAllMatureValidators(ctx)
	if err != nil {
		return nil, err
	}

	return validatorUpdates, nil
}

// ApplyAndReturnValidatorSetUpdates applies and return accumulated updates to the bonded validator set. Also,
// * Updates the active valset as keyed by LastValidatorPowerKey.
// * Updates the total power as keyed by LastTotalPowerKey.
// * Updates relevant indices.
// It gets called once after genesis, another time maybe after genesis transactions,
// then once at every EndBlock.
//
// CONTRACT: Only validators with non-zero power or zero-power that were bonded
// at the previous block height or were removed from the validator set entirely
// are returned to CometBFT.
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx context.Context) ([]appmodule.ValidatorUpdate, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	maxValidators := params.MaxValidators
	totalPower := math.ZeroInt()

	last, err := k.getLastValidatorsByAddr(ctx)
	if err != nil {
		return nil, err
	}

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

		// fetch the old power bytes
		addrStr, err := k.validatorAddressCodec.BytesToString(addr)
		if err != nil {
			return nil, err
		}

		oldPowerBytes, found := last[addrStr]
		newPower := validator.ConsensusPower(power)
		newPowerBytes := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: newPower})

		// update the validator set if power has changed
		if !found || !bytes.Equal(oldPowerBytes, newPowerBytes) {
			updates = append(updates, validator.ModuleValidatorUpdate(power))
			if err = k.SetLastValidatorPower(ctx, addr, newPower); err != nil {
				return nil, err
			}
		}

		delete(last, addrStr)
		count++

		totalPower = totalPower.Add(math.NewInt(validator.PotentialConsensusPower(power)))
	}

	noLongerBonded, err := sortNoLongerBonded(last, k.validatorAddressCodec)
	if err != nil {
		return nil, err
	}

	for _, valAddrBytes := range noLongerBonded {
		validator, err := k.GetValidator(ctx, valAddrBytes)
		if err != nil {
			return nil, err
		}

		str, err := k.validatorAddressCodec.StringToBytes(validator.GetOperator())
		if err != nil {
			return nil, err
		}
		if err = k.DeleteLastValidatorPower(ctx, str); err != nil {
			return nil, err
		}

		updates = append(updates, validator.ModuleValidatorUpdateZero())
	}

	if len(updates) > 0 {
		if err = k.LastTotalPower.Set(ctx, totalPower); err != nil {
			return nil, err
		}
	}

	return updates, err
}

// UnbondingToUnbonded switches a validator from unbonding state to unbonded state
func (k Keeper) UnbondingToUnbonded(ctx context.Context, validator types.Validator) (types.Validator, error) {
	if !validator.IsUnbonding() {
		return types.Validator{}, fmt.Errorf("bad state transition unbondingToUnbonded, validator: %v", validator)
	}

	return k.completeUnbondingValidator(ctx, validator)
}

// perform all the store operations for when a validator status becomes unbonded
func (k Keeper) completeUnbondingValidator(ctx context.Context, validator types.Validator) (types.Validator, error) {
	validator.UpdateStatus(types.Unbonded)

	if err := k.SetValidator(ctx, validator); err != nil {
		return validator, err
	}

	return validator, nil
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

	// trigger hook
	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return validator, err
	}

	str, err := k.validatorAddressCodec.StringToBytes(validator.GetOperator())
	if err != nil {
		return validator, err
	}

	if err := k.Hooks().AfterValidatorBonded(ctx, consAddr, str); err != nil {
		return validator, err
	}

	return validator, err
}

// map of operator bech32-addresses to serialized power
// We use bech32 strings here, because we can't have slices as keys: map[[]byte][]byte
type validatorsByAddr map[string][]byte

// get the last validator set
func (k Keeper) getLastValidatorsByAddr(ctx context.Context) (validatorsByAddr, error) {
	last := make(validatorsByAddr)

	err := k.LastValidatorPower.Walk(ctx, nil, func(key []byte, value gogotypes.Int64Value) (bool, error) {
		valAddrStr, err := k.validatorAddressCodec.BytesToString(key)
		if err != nil {
			return true, err
		}

		intV := value.GetValue()
		bz := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: intV})
		last[valAddrStr] = bz
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return last, nil
}

// given a map of remaining validators to previous bonded power
// returns the list of validators to be unbonded, sorted by operator address
func sortNoLongerBonded(last validatorsByAddr, ac address.Codec) ([][]byte, error) {
	// sort the map keys for determinism
	noLongerBonded := make([][]byte, len(last))
	index := 0

	for valAddrStr := range last {
		valAddrBytes, err := ac.StringToBytes(valAddrStr)
		if err != nil {
			return nil, err
		}
		noLongerBonded[index] = valAddrBytes
		index++
	}
	// sorted by address - order doesn't matter
	sort.SliceStable(noLongerBonded, func(i, j int) bool {
		// -1 means strictly less than
		return bytes.Compare(noLongerBonded[i], noLongerBonded[j]) == -1
	})

	return noLongerBonded, nil
}

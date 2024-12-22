package keeper

import (
	"bytes"
	"context"
	"cosmossdk.io/collections"
	corestore "cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/cosmos/gogoproto/types"
	"kepler/x/staking/types"
	"time"
)

var timeBzKeySize = uint64(29) // time bytes key size is 29 by default

// ValidatorsPowerStoreIterator returns an iterator for the current validator power store
func (k Keeper) ValidatorsPowerStoreIterator(ctx context.Context) (corestore.Iterator, error) {
	store := k.KVStoreService.OpenKVStore(ctx)

	return store.ReverseIterator(
		types.ValidatorsByPowerIndexKey,
		storetypes.PrefixEndBytes(types.ValidatorsByPowerIndexKey),
	)
}

// DeleteValidatorByPowerIndex deletes a record by power index
func (k Keeper) DeleteValidatorByPowerIndex(ctx context.Context, validator types.Validator) error {
	store := k.KVStoreService.OpenKVStore(ctx)
	power, err := k.PowerReduction(ctx)
	if err != nil {
		return err
	}

	return store.Delete(types.GetValidatorsByPowerIndexKey(validator, power, k.validatorAddressCodec))
}

// SetValidatorByPowerIndex sets a validator by power index
func (k Keeper) SetValidatorByPowerIndex(ctx context.Context, validator types.Validator) error {
	if validator.Jailed {
		return nil
	}

	store := k.KVStoreService.OpenKVStore(ctx)
	str, err := k.validatorAddressCodec.StringToBytes(validator.GetOperator())
	if err != nil {
		return err
	}

	power, err := k.PowerReduction(ctx)
	if err != nil {
		return err
	}

	return store.Set(types.GetValidatorsByPowerIndexKey(validator, power, k.validatorAddressCodec), str)
}

// GetValidator gets a single validator
func (k Keeper) GetValidator(ctx context.Context, addr sdk.ValAddress) (validator types.Validator, err error) {
	validator, err = k.Validators.Get(ctx, addr)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return types.Validator{}, types.ErrNoValidatorFound
		}
		return validator, err
	}
	return validator, nil
}

// SetValidator sets the main record holding validator details
func (k Keeper) SetValidator(ctx context.Context, validator types.Validator) error {
	valBz, err := k.ValidatorAddressCodec().StringToBytes(validator.GetOperator())
	if err != nil {
		return err
	}

	return k.Validators.Set(ctx, valBz, validator)
}

// SetLastValidatorPower sets the last validator power.
func (k Keeper) SetLastValidatorPower(ctx context.Context, operator sdk.ValAddress, power int64) error {
	return k.LastValidatorPower.Set(ctx, operator, gogotypes.Int64Value{Value: power})
}

// DeleteLastValidatorPower deletes the last validator power.
func (k Keeper) DeleteLastValidatorPower(ctx context.Context, operator sdk.ValAddress) error {
	return k.LastValidatorPower.Remove(ctx, operator)
}

func (k Keeper) UnbondAllMatureValidators(ctx context.Context) error {
	headerInfo := k.HeaderService.HeaderInfo(ctx)
	blockTime := headerInfo.Time
	blockHeight := uint64(headerInfo.Height)

	rng := new(collections.Range[collections.Triple[uint64, time.Time, uint64]]).
		EndInclusive(collections.Join3(uint64(29), blockTime, blockHeight))

	iter, err := k.ValidatorQueue.Iterate(ctx, rng)
	if err != nil {
		return err
	}

	kvs, err := iter.KeyValues()
	if err != nil {
		return err
	}

	for _, kv := range kvs {
		if err := k.unbondMatureValidators(ctx, blockHeight, blockTime, kv.Key, kv.Value); err != nil {
			return err
		}
	}

	return nil
}

// GetUnbondingValidators returns a slice of mature validator addresses that
// complete their unbonding at a given time and height.
func (k Keeper) GetUnbondingValidators(ctx context.Context, endTime time.Time, endHeight int64) ([]string, error) {
	timeSize := sdk.TimeKey.Size(endTime)
	valAddrs, err := k.ValidatorQueue.Get(ctx, collections.Join3(uint64(timeSize), endTime, uint64(endHeight)))
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return []string{}, err
	}

	return valAddrs.Addresses, nil
}

// DeleteValidatorQueueTimeSlice deletes all entries in the queue indexed by a
// given height and time.
func (k Keeper) DeleteValidatorQueueTimeSlice(ctx context.Context, endTime time.Time, endHeight int64) error {
	return k.ValidatorQueue.Remove(ctx, collections.Join3(timeBzKeySize, endTime, uint64(endHeight)))
}

// InsertUnbondingValidatorQueue inserts a given unbonding validator address into
// the unbonding validator queue for a given height and time.
func (k Keeper) InsertUnbondingValidatorQueue(ctx context.Context, val types.Validator) error {
	addrs, err := k.GetUnbondingValidators(ctx, val.UnbondingTime, val.UnbondingHeight)
	if err != nil {
		return err
	}
	addrs = append(addrs, val.OperatorAddress)
	return k.SetUnbondingValidatorsQueue(ctx, val.UnbondingTime, val.UnbondingHeight, addrs)
}

// SetUnbondingValidatorsQueue sets a given slice of validator addresses into
// the unbonding validator queue by a given height and time.
func (k Keeper) SetUnbondingValidatorsQueue(ctx context.Context, endTime time.Time, endHeight int64, addrs []string) error {
	valAddrs := types.ValAddresses{Addresses: addrs}
	return k.ValidatorQueue.Set(ctx, collections.Join3(timeBzKeySize, endTime, uint64(endHeight)), valAddrs)
}

// DeleteValidatorQueue removes a validator by address from the unbonding queue
// indexed by a given height and time.
func (k Keeper) DeleteValidatorQueue(ctx context.Context, val types.Validator) error {
	addrs, err := k.GetUnbondingValidators(ctx, val.UnbondingTime, val.UnbondingHeight)
	if err != nil {
		return err
	}
	newAddrs := []string{}

	// since address string may change due to Bech32 prefix change, we parse the addresses into bytes
	// format for normalization
	deletingAddr, err := k.validatorAddressCodec.StringToBytes(val.OperatorAddress)
	if err != nil {
		return err
	}

	for _, addr := range addrs {
		storedAddr, err := k.validatorAddressCodec.StringToBytes(addr)
		if err != nil {
			// even if we don't error here, it will error in UnbondAllMatureValidators at unbond time
			return err
		}
		if !bytes.Equal(storedAddr, deletingAddr) {
			newAddrs = append(newAddrs, addr)
		}
	}

	if len(newAddrs) == 0 {
		return k.DeleteValidatorQueueTimeSlice(ctx, val.UnbondingTime, val.UnbondingHeight)
	}

	return k.SetUnbondingValidatorsQueue(ctx, val.UnbondingTime, val.UnbondingHeight, newAddrs)
}

// RemoveValidator removes the validator record and associated indexes
// except for the bonded validator index which is only handled in ApplyAndReturnTendermintUpdates
func (k Keeper) RemoveValidator(ctx context.Context, address sdk.ValAddress) error {
	// first retrieve the old validator record
	validator, err := k.GetValidator(ctx, address)
	if errors.Is(err, types.ErrNoValidatorFound) {
		return nil
	}

	if !validator.IsUnbonded() {
		return errorsmod.Wrap(types.ErrBadRemoveValidator, "cannot call RemoveValidator on bonded or unbonding validators")
	}

	valConsAddr, err := validator.GetConsAddr()
	if err != nil {
		return err
	}

	// delete the old validator record
	store := k.KVStoreService.OpenKVStore(ctx)
	if err = k.Validators.Remove(ctx, address); err != nil {
		return err
	}

	if err = k.ValidatorByConsensusAddress.Remove(ctx, valConsAddr); err != nil {
		return err
	}

	power, err := k.PowerReduction(ctx)
	if err != nil {
		return err
	}

	if err = store.Delete(types.GetValidatorsByPowerIndexKey(validator, power, k.validatorAddressCodec)); err != nil {
		return err
	}

	str, err := k.validatorAddressCodec.StringToBytes(validator.GetOperator())
	if err != nil {
		return err
	}

	if err := k.Hooks().AfterValidatorRemoved(ctx, valConsAddr, str); err != nil {
		return fmt.Errorf("error in after validator removed hook: %w", err)
	}

	return nil
}

// SetValidatorByConsAddr sets a validator by consensus address
func (k Keeper) SetValidatorByConsAddr(ctx context.Context, validator types.Validator) error {
	consPk, err := validator.GetConsAddr()
	if err != nil {
		return err
	}

	bz, err := k.validatorAddressCodec.StringToBytes(validator.GetOperator())
	if err != nil {
		return err
	}

	return k.ValidatorByConsensusAddress.Set(ctx, consPk, bz)
}

// IterateLastValidatorPowers iterates over last validator powers.
func (k Keeper) IterateLastValidatorPowers(
	ctx context.Context,
	handler func(operator sdk.ValAddress, power int64) (stop bool),
) error {
	err := k.LastValidatorPower.Walk(ctx, nil, func(key []byte, value gogotypes.Int64Value) (bool, error) {
		addr := sdk.ValAddress(key)

		if handler(addr, value.GetValue()) {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return err
	}

	return nil
}

// GetAllValidators gets the set of all validators with no limits, used during genesis dump
func (k Keeper) GetAllValidators(ctx context.Context) (validators []types.Validator, err error) {
	store := k.KVStoreService.OpenKVStore(ctx)

	iterator, err := store.Iterator(types.ValidatorsKey, storetypes.PrefixEndBytes(types.ValidatorsKey))
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validator, err := types.UnmarshalValidator(k.cdc, iterator.Value())
		if err != nil {
			return nil, err
		}
		validators = append(validators, validator)
	}

	return validators, nil
}

func (k Keeper) unbondMatureValidators(
	ctx context.Context,
	blockHeight uint64,
	blockTime time.Time,
	key collections.Triple[uint64, time.Time, uint64],
	addrs types.ValAddresses,
) error {
	keyTime, keyHeight := key.K2(), key.K3()

	// All addresses for the given key have the same unbonding height and time.
	// We only unbond if the height and time are less than the current height
	// and time.
	if keyHeight > blockHeight || keyTime.After(blockTime) {
		return nil
	}

	// finalize unbonding
	for _, valAddr := range addrs.Addresses {
		addr, err := k.validatorAddressCodec.StringToBytes(valAddr)
		if err != nil {
			return err
		}
		val, err := k.GetValidator(ctx, addr)
		if err != nil {
			return errorsmod.Wrap(err, "validator in the unbonding queue was not found")
		}

		if !val.IsUnbonding() {
			return errors.New("unexpected validator in unbonding queue; status was not unbonding")
		}

		val, err = k.UnbondingToUnbonded(ctx, val)
		if err != nil {
			return err
		}

		addr, err = k.validatorAddressCodec.StringToBytes(val.OperatorAddress)
		if err != nil {
			return err
		}

		if err := k.RemoveValidator(ctx, addr); err != nil {
			return err
		}

		// remove validator from queue
		if err = k.DeleteValidatorQueue(ctx, val); err != nil {
			return err
		}
	}

	return nil
}

package keeper

import (
	"context"
	"cosmossdk.io/collections"
	corestore "cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"kepler/x/xstaking/types"
)

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

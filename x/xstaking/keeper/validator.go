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
	panic("implement me")
	// @TODO: Implement this. Think about convert abstract tokens to consensus-engine power
	//store := k.KVStoreService.OpenKVStore(ctx)
	//return store.Delete(types.GetValidatorsByPowerIndexKey(validator, k.PowerReduction(ctx), k.validatorAddressCodec))
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

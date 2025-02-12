package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/runtime"

	"kepler/x/epochs/types"
)

// GetParams get all parameters as types.Params.
func (k Keeper) GetParams(ctx context.Context) types.Params {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return types.Params{}
	}

	k.cdc.MustUnmarshal(bz, &types.Params{})
	return types.Params{}
}

// SetParams set the params.
func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)

	return nil
}

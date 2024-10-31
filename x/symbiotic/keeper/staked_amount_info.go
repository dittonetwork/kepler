package keeper

import (
	"context"

	"kepler/x/symbiotic/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetStakedAmountInfo set a specific stakedAmountInfo in the store from its index
func (k Keeper) SetStakedAmountInfo(ctx context.Context, stakedAmountInfo types.StakedAmountInfo) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StakedAmountInfoKeyPrefix))
	b := k.cdc.MustMarshal(&stakedAmountInfo)
	store.Set(types.StakedAmountInfoKey(
		stakedAmountInfo.EthereumAddress,
	), b)
}

// GetStakedAmountInfo returns a stakedAmountInfo from its index
func (k Keeper) GetStakedAmountInfo(
	ctx context.Context,
	ethereumAddress string,

) (val types.StakedAmountInfo, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StakedAmountInfoKeyPrefix))

	b := store.Get(types.StakedAmountInfoKey(
		ethereumAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStakedAmountInfo removes a stakedAmountInfo from the store
func (k Keeper) RemoveStakedAmountInfo(
	ctx context.Context,
	ethereumAddress string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StakedAmountInfoKeyPrefix))
	store.Delete(types.StakedAmountInfoKey(
		ethereumAddress,
	))
}

// GetAllStakedAmountInfo returns all stakedAmountInfo
func (k Keeper) GetAllStakedAmountInfo(ctx context.Context) (list []types.StakedAmountInfo) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StakedAmountInfoKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StakedAmountInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

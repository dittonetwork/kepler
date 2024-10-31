package keeper

import (
	"context"

	"kepler/x/beacon/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetFinalizedBlockInfo set finalizedBlockInfo in the store
func (k Keeper) SetFinalizedBlockInfo(ctx context.Context, finalizedBlockInfo types.FinalizedBlockInfo) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FinalizedBlockInfoKey))
	b := k.cdc.MustMarshal(&finalizedBlockInfo)
	store.Set([]byte{0}, b)
}

// GetFinalizedBlockInfo returns finalizedBlockInfo
func (k Keeper) GetFinalizedBlockInfo(ctx context.Context) (val types.FinalizedBlockInfo, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FinalizedBlockInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFinalizedBlockInfo removes finalizedBlockInfo from the store
func (k Keeper) RemoveFinalizedBlockInfo(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FinalizedBlockInfoKey))
	store.Delete([]byte{0})
}

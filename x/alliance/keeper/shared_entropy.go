package keeper

import (
	"context"

	"kepler/x/alliance/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetSharedEntropy set sharedEntropy in the store
func (k Keeper) SetSharedEntropy(ctx context.Context, sharedEntropy types.SharedEntropy) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SharedEntropyKey))
	b := k.cdc.MustMarshal(&sharedEntropy)
	store.Set([]byte{0}, b)
}

// GetSharedEntropy returns sharedEntropy
func (k Keeper) GetSharedEntropy(ctx context.Context) (val types.SharedEntropy, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SharedEntropyKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSharedEntropy removes sharedEntropy from the store
func (k Keeper) RemoveSharedEntropy(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SharedEntropyKey))
	store.Delete([]byte{0})
}

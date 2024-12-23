package keeper

import (
	"context"

	"kepler/x/alliance/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetQuorumParams set quorumParams in the store
func (k Keeper) SetQuorumParams(ctx context.Context, quorumParams types.QuorumParams) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuorumParamsKey))
	b := k.cdc.MustMarshal(&quorumParams)
	store.Set([]byte{0}, b)
}

// GetQuorumParams returns quorumParams
func (k Keeper) GetQuorumParams(ctx context.Context) (val types.QuorumParams, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuorumParamsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveQuorumParams removes quorumParams from the store
func (k Keeper) RemoveQuorumParams(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuorumParamsKey))
	store.Delete([]byte{0})
}

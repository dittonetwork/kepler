package keeper

import (
	"context"

	"kepler/x/symbiotic/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetContractAddress set contractAddress in the store
func (k Keeper) SetContractAddress(ctx context.Context, contractAddress types.ContractAddress) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContractAddressKey))
	b := k.cdc.MustMarshal(&contractAddress)
	store.Set([]byte{0}, b)
}

// GetContractAddress returns contractAddress
func (k Keeper) GetContractAddress(ctx context.Context) (val types.ContractAddress, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContractAddressKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContractAddress removes contractAddress from the store
func (k Keeper) RemoveContractAddress(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContractAddressKey))
	store.Delete([]byte{0})
}

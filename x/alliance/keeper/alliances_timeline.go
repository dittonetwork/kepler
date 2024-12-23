package keeper

import (
	"context"
	"encoding/binary"

	"kepler/x/alliance/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetAlliancesTimelineCount get the total number of alliancesTimeline
func (k Keeper) GetAlliancesTimelineCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.AlliancesTimelineCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAlliancesTimelineCount set the total number of alliancesTimeline
func (k Keeper) SetAlliancesTimelineCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.AlliancesTimelineCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAlliancesTimeline appends a alliancesTimeline in the store with a new id and update the count
func (k Keeper) AppendAlliancesTimeline(
	ctx context.Context,
	alliancesTimeline types.AlliancesTimeline,
) uint64 {
	// Create the alliancesTimeline
	count := k.GetAlliancesTimelineCount(ctx)

	// Set the ID of the appended value
	alliancesTimeline.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AlliancesTimelineKey))
	appendedValue := k.cdc.MustMarshal(&alliancesTimeline)
	store.Set(GetAlliancesTimelineIDBytes(alliancesTimeline.Id), appendedValue)

	// Update alliancesTimeline count
	k.SetAlliancesTimelineCount(ctx, count+1)

	return count
}

// SetAlliancesTimeline set a specific alliancesTimeline in the store
func (k Keeper) SetAlliancesTimeline(ctx context.Context, alliancesTimeline types.AlliancesTimeline) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AlliancesTimelineKey))
	b := k.cdc.MustMarshal(&alliancesTimeline)
	store.Set(GetAlliancesTimelineIDBytes(alliancesTimeline.Id), b)
}

// GetAlliancesTimeline returns a alliancesTimeline from its id
func (k Keeper) GetAlliancesTimeline(ctx context.Context, id uint64) (val types.AlliancesTimeline, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AlliancesTimelineKey))
	b := store.Get(GetAlliancesTimelineIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAlliancesTimeline removes a alliancesTimeline from the store
func (k Keeper) RemoveAlliancesTimeline(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AlliancesTimelineKey))
	store.Delete(GetAlliancesTimelineIDBytes(id))
}

// GetAllAlliancesTimeline returns all alliancesTimeline
func (k Keeper) GetAllAlliancesTimeline(ctx context.Context) (list []types.AlliancesTimeline) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AlliancesTimelineKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AlliancesTimeline
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAlliancesTimelineIDBytes returns the byte representation of the ID
func GetAlliancesTimelineIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.AlliancesTimelineKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

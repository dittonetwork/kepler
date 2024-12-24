package keeper

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/rand"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"kepler/x/alliance/types"

	"cosmossdk.io/runtime"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdktypes "cosmossdk.io/types"
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

// FillAlliancesTimeline feels an array of future aliances useing validators list
func (k Keeper) FillAlliancesTimeline(goCtx context.Context) error {
	currentAlliances := k.GetAllAlliancesTimeline(goCtx)

	sdkCtx := sdktypes.UnwrapSDKContext(goCtx)
	currentBlockHeight := sdkCtx.BlockHeight()

	var alliancesToRemove []types.AlliancesTimeline
	for _, alliance := range currentAlliances {
		if int64(alliance.EndBlock) < currentBlockHeight {
			alliancesToRemove = append(alliancesToRemove, alliance)
		}
	}

	quorumParams, exists := k.GetQuorumParams(goCtx)
	if !exists {
		// TODO: move to errors.go
		panic(fmt.Errorf("quorum params not set"))
	}

	// clean previous alliances from the list
	if len(alliancesToRemove) != 0 {
		k.Logger().Info("removing past alliances", "num", len(alliancesToRemove))
	}
	for _, allianceToRemove := range alliancesToRemove {
		k.RemoveAlliancesTimeline(goCtx, allianceToRemove.Id)
	}

	alliancesLeft := len(currentAlliances) - len(alliancesToRemove)
	// If there are more alliances than needed, we don't remove extra ones
	if alliancesLeft >= int(quorumParams.GetLifetimeInBlocks()) {
		k.Logger().Info("alliances timeline doesn't require modification")
		return nil
	}

	var lastEndBlock = currentBlockHeight
	var lastId uint
	if len(currentAlliances) != 0 {
		lastEndBlock = uint(currentAlliances[len(currentAlliances)-1].EndBlock)
		lastId = uint(currentAlliances[len(currentAlliances)-1].Id)
	}

	sharedEntropy, _ := k.GetSharedEntropy(goCtx)
	seed := int64(sharedEntropy.Entropy) ^ sdkCtx.BlockHash
	fmt.Printf("block hash: %s\n", sdkCtx.BlockHash) // TODO: remove after check that BlockHash is not empty
	for i := 0; i < int(quorumParams.GetLifetimeInBlocks())-alliancesLeft; i++ {
		startBlock := lastEndBlock
		endBlock := startBlock + quorumParams.LifetimeInBlocks
		k.Logger().Debug("selecting validators", "startBlkNum", startBlock, "endBlkNum", endBlock)
		validators, err := k.selectValidatorsForAlliance(goCtx, uint(quorumParams.MaxParticipants), seed+i)
		if err != nil {
			return err
		}
		validatorsAddresses := make([]string, len(validators))
		for j, validator := range validators {
			validatorsAddresses[j] = validator.OperatorAddress
		}
		alliancesTimeline := types.AlliancesTimeline{
			Id:           uint64(lastId + uint(i)),
			Participants: validatorsAddresses,
			StartBlock:   startBlock,
			EndBlock:     endBlock,
		}
		k.AppendAlliancesTimeline(goCtx, alliancesTimeline)
		lastEndBlock = endBlock
	}

	return nil
}

// generateRandomIndices generates k unique random indices from 0 to min(n-1, k).
func generateRandomIndices(n uint, k uint, seed uint) []int {
	rand := rand.New(rand.NewSource(int64(seed)))

	if k > n {
		k = n
	}
	indices := rand.Perm(int(n))[:k]
	return indices
}

// selectValidatorsForAlliance returns validators from pool to form new alliance
func (k Keeper) selectValidatorsForAlliance(goCtx context.Context, maxParticipants uint, seed uint) ([]stakingtypes.Validator, error) {
	allValidators, err := k.stakingKeeper.GetAllValidators(goCtx)
	if err != nil {
		return nil, err
	}

	randomIndices := generateRandomIndices(uint(len(allValidators)), maxParticipants, seed)

	var generatedAliance []stakingtypes.Validator
	for i, randomIndex := range randomIndices {
		if uint(i) == maxParticipants {
			break
		}
		generatedAliance = append(generatedAliance, allValidators[randomIndex])
	}

	return generatedAliance, nil
}

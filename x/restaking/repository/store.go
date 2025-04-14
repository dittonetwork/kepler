package repository

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

type RestakingRepository struct {
	pending    collections.Map[string, types.Operator]
	validators *collections.IndexedMap[string, types.Validator, Idx]
	lastUpdate collections.Item[types.UpdateInfo]

	cdc codec.BinaryCodec
}

func New(storeService store.KVStoreService, cdc codec.BinaryCodec) *RestakingRepository {
	sb := collections.NewSchemaBuilder(storeService)

	return &RestakingRepository{
		pending: collections.NewMap(
			sb,
			types.KeyPrefixPendingOperators,
			"pending",
			collections.StringKey,
			codec.CollValue[types.Operator](cdc),
		),
		validators: collections.NewIndexedMap(
			sb,
			types.KeyPrefixValidators,
			"validators",
			collections.StringKey,
			codec.CollValue[types.Validator](cdc),
			NewIndexes(sb),
		),
		lastUpdate: collections.NewItem(
			sb,
			types.KeyPrefixLastUpdate,
			"last_update",
			codec.CollValue[types.UpdateInfo](cdc),
		),

		cdc: cdc,
	}
}

// GetLastUpdate retrieves the last update information from the store.
func (s *RestakingRepository) GetLastUpdate(ctx sdk.Context) (types.UpdateInfo, error) {
	lastUpdate, err := s.lastUpdate.Get(ctx)
	if err != nil {
		return types.UpdateInfo{}, err
	}

	return lastUpdate, nil
}

// SetLastUpdate sets the last update information in the store.
func (s *RestakingRepository) SetLastUpdate(ctx sdk.Context, lastUpdate types.UpdateInfo) error {
	return s.lastUpdate.Set(ctx, lastUpdate)
}

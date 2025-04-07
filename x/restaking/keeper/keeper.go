package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/dittonetwork/kepler/x/restaking/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger
		hooks        types.RestakingHooks

		pendingValidators collections.Map[string, types.Validator]
		validators        *collections.IndexedMap[string, types.Validator, Idx]
		lastUpdate        collections.Item[types.UpdateInfo]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
) *Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	return &Keeper{
		cdc:          cdc,
		storeService: storeService,
		logger:       logger,

		pendingValidators: collections.NewMap(
			sb,
			types.KeyPrefixPendingValidators,
			"pending",
			collections.StringKey,
			codec.CollValue[types.Validator](cdc),
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
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetHooks set the gamm hooks.
func (k *Keeper) SetHooks(rh types.RestakingHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set hooks twice")
	}

	k.hooks = rh

	return k
}

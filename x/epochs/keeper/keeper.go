package keeper

import (
	"fmt"

	"github.com/dittonetwork/kepler/x/epochs/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
)

type (
	Keeper struct {
		cdc    codec.BinaryCodec
		logger log.Logger
		hooks  types.EpochHooks

		Schema    collections.Schema
		EpochInfo collections.Map[string, types.EpochInfo]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
) *Keeper {
	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:    cdc,
		logger: logger,
		EpochInfo: collections.NewMap(
			sb,
			types.KeyPrefixEpoch,
			"epoch_info",
			collections.StringKey,
			codec.CollValue[types.EpochInfo](cdc),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return &k
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetHooks set the gamm hooks.
func (k *Keeper) SetHooks(eh types.EpochHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set epochs hooks twice")
	}

	k.hooks = eh

	return k
}

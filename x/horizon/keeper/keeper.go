package keeper

import (
	"cosmossdk.io/orm/model/ormdb"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/codec"

	"kepler/x/horizon/types"
	"kepler/x/horizon/types/state"
)

type Keeper struct {
	appmodule.Environment

	addressCodec address.Codec
	authority    []byte
	cdc          codec.BinaryCodec
	state        state.StateStore

	Schema collections.Schema
	Params collections.Item[types.Params]
	// this line is used by starport scaffolding # collection/type

}

func NewKeeper(
	addressCodec address.Codec,
	authority []byte,
	cdc codec.BinaryCodec,
	env appmodule.Environment,

) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	modDb, err := ormdb.NewModuleDB(StateSchema, ormdb.ModuleDBOptions{KVStoreService: env.KVStoreService})
	if err != nil {
		panic(err)
	}

	db, err := state.NewStateStore(modDb)
	if err != nil {
		panic(err)
	}

	sb := collections.NewSchemaBuilder(env.KVStoreService)

	k := Keeper{
		addressCodec: addressCodec,
		authority:    authority,
		cdc:          cdc,
		Environment:  env,
		state:        db,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		// this line is used by starport scaffolding # collection/instantiate
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}

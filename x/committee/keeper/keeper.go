package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kepler/x/committee/types"
)

type CommitteeKeeper interface {
	// IsCommitteeExists returns true if the committee exists
	IsCommitteeExists(ctx sdk.Context, committeeID string) (bool, error)

	// GetAuthority returns the module's authority.
	GetAuthority() string

	// SetParams updates the committee module's parameters.
	SetParams(ctx context.Context, params types.Params) error
}

type committeeIndexes struct {
	// Key: ChainID | Value: CommitteeID | Type: MultiKey
	ChainID *indexes.Multi[string, string, types.Committee]

	// Key: ChainID | Value: CommitteeID | Type: Unique
	Active *indexes.Unique[string, string, types.Committee]
}

func (i committeeIndexes) IndexesList() []collections.Index[string, types.Committee] {
	return []collections.Index[string, types.Committee]{
		i.ChainID,
		i.Active,
	}
}

func newCommitteeIndexes(sb *collections.SchemaBuilder) committeeIndexes {
	return committeeIndexes{
		ChainID: indexes.NewMulti(
			sb,
			types.ChainIDStoreKeyPrefix,
			"committee_by_chain_id",
			collections.StringKey,
			collections.StringKey,
			func(_ string, value types.Committee) (string, error) {
				return value.ChainId, nil
			},
		),
		Active: indexes.NewUnique(
			sb,
			types.ActiveCommitteeStoreKeyPrefix,
			"active_committee",
			collections.StringKey,
			collections.StringKey,
			func(_ string, value types.Committee) (string, error) {
				return value.ChainId, nil
			},
		),
	}
}

type Keeper struct {
	cdc    codec.BinaryCodec
	logger log.Logger

	Schema     collections.Schema
	Committees *collections.IndexedMap[string, types.Committee, committeeIndexes]

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:       cdc,
		authority: authority,
		logger:    logger,

		Committees: collections.NewIndexedMap(
			sb,
			types.CommitteeStoreKeyPrefix,
			"committees",
			collections.StringKey,
			codec.CollValue[types.Committee](cdc),
			newCommitteeIndexes(sb),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

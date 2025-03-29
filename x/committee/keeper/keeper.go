package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dittonetwork/kepler/x/committee/types"
)

type CommitteeKeeper interface {
	// IsCommitteeExists returns true if the committee exists
	// Deprecated: CanBeSigned returns true if the message can be signed.
	IsCommitteeExists(ctx sdk.Context, committeeID string) (bool, error)

	// CreateCommittee creates a new committee for the given epoch.
	CreateCommittee(ctx sdk.Context, epoch uint32) (types.Committee, error)

	// GetAuthority returns the module's authority.
	GetAuthority() string

	// SetParams updates the committee module's parameters.
	SetParams(ctx context.Context, params types.Params) error
}

type Keeper struct {
	cdc codec.BinaryCodec

	Schema     collections.Schema
	Committees *collections.IndexedMap[uint32, types.Committee, Idx]
	LastEpoch  collections.Item[uint32]

	executors types.ExecutorsKeeper
	restaking types.RestakingKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	authority string,
	executors types.ExecutorsKeeper,
	restaking types.RestakingKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:       cdc,
		authority: authority,
		executors: executors,
		restaking: restaking,

		Committees: collections.NewIndexedMap(
			sb,
			types.CommitteesStoreKeyPrefix,
			"committees",
			collections.Uint32Key,
			codec.CollValue[types.Committee](cdc),
			NewIndexes(sb),
		),
		LastEpoch: collections.NewItem(
			sb,
			types.LatestEpochStorePrefix,
			"last_epoch",
			collections.Uint32Value,
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// Deprecated: This method is deprecated and will be reworked.
// CanBeSigned returns true if the message can be signed.
func (k Keeper) CanBeSigned(
	_ sdk.Context,
	_ string,
	_ string,
	_ [][]byte,
	_ []byte,
) (bool, error) {
	return true, nil
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	return sdkCtx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

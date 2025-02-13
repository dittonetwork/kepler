package keeper

import (
	"fmt"

	"kepler/x/committee/types"

	"cosmossdk.io/collections"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger
		epochKeeper  types.EpochsKeeper

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
		Schema    collections.Schema
		// Randao Commits key:chainID+epochID+validatorAddress
		RandaoCommitments collections.Map[collections.Triple[string, uint64, sdk.ValAddress], types.CommitRandao]
		// Randao Reveals key:chainID+epochID+validatorAddress
		RandaoReveals collections.Map[collections.Triple[string, uint64, sdk.ValAddress], types.RevealRandao]

		validatorAddressCodec address.Codec
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	valAddrCodec address.Codec,
	ek types.EpochsKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:                   cdc,
		storeService:          storeService,
		authority:             authority,
		logger:                logger,
		validatorAddressCodec: valAddrCodec,
		epochKeeper:           ek,
		RandaoCommitments: collections.NewMap(
			sb,
			types.RandaoCommitmentsKey,
			"randao_commits",
			collections.TripleKeyCodec(
				collections.StringKey,
				collections.Uint64Key,
				sdk.ValAddressKey,
			),
			codec.CollValue[types.CommitRandao](cdc),
		),
		RandaoReveals: collections.NewMap(
			sb,
			types.RandaoRevealsKey,
			"randao_reveals",
			collections.TripleKeyCodec(
				collections.StringKey,
				collections.Uint64Key,
				sdk.ValAddressKey,
			),
			codec.CollValue[types.RevealRandao](cdc),
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

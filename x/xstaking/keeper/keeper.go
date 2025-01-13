package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	"cosmossdk.io/collections"
	addresscodec "cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	bankkeeper "cosmossdk.io/x/bank/keeper"
	stakingkeeper "cosmossdk.io/x/staking/keeper"

	"kepler/x/xstaking/types"
)

type Keeper struct {
	appmodule.Environment

	cdc                   codec.BinaryCodec
	bankKeeper            bankkeeper.Keeper
	consensusKeeper       types.ConsensusKeeper
	stakingKeeper         *stakingkeeper.Keeper
	authority             []byte
	validatorAddressCodec addresscodec.Codec

	Schema collections.Schema
	Params collections.Item[types.Params]
}

func NewKeeper(
	env appmodule.Environment,
	cdc codec.BinaryCodec,
	bankKeeper bankkeeper.Keeper,
	consensusKeeper types.ConsensusKeeper,
	stakingKeeper *stakingkeeper.Keeper,
	authority []byte,
	validatorAddressCodec addresscodec.Codec,

) Keeper {
	if _, err := validatorAddressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(env.KVStoreService)

	k := Keeper{
		Environment:           env,
		cdc:                   cdc,
		bankKeeper:            bankKeeper,
		consensusKeeper:       consensusKeeper,
		stakingKeeper:         stakingKeeper,
		authority:             authority,
		validatorAddressCodec: validatorAddressCodec,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
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

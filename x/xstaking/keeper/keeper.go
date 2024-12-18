package keeper

import (
	"cosmossdk.io/collections"
	collcodec "cosmossdk.io/collections/codec"
	"cosmossdk.io/core/address"
	addresscodec "cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/math"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/cosmos/gogoproto/types"
	"kepler/x/xstaking/types"
	"time"
)

type Keeper struct {
	appmodule.Environment

	cdc                   codec.BinaryCodec
	addressCodec          address.Codec
	hooks                 types.StakingHooks
	validatorAddressCodec addresscodec.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema collections.Schema
	Params collections.Item[types.Params]

	// LastTotalPower value: LastTotalPower
	LastTotalPower collections.Item[math.Int]
	// Validators key: valAddr, value: Validator
	Validators collections.Map[[]byte, types.Validator]
	// LastValidatorPower key: valAddr | value: power(gogotypes.Int64Value())
	LastValidatorPower collections.Map[[]byte, gogotypes.Int64Value]
	// ValidatorQueue key: len(timestamp bytes)+timestamp+height | value: ValAddresses
	ValidatorQueue collections.Map[collections.Triple[uint64, time.Time, uint64], types.ValAddresses]
	// ValidatorByConsensusAddress key: consAddr, value: valAddr
	ValidatorByConsensusAddress collections.Map[sdk.ConsAddress, sdk.ValAddress]
}

func NewKeeper(
	env appmodule.Environment,
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	authority []byte,
	validatorAddressCodec addresscodec.Codec,
) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	if validatorAddressCodec == nil {
		panic("validator address codec is required")
	}

	sb := collections.NewSchemaBuilder(env.KVStoreService)

	k := Keeper{
		Environment:  env,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		hooks:        nil,

		validatorAddressCodec: validatorAddressCodec,

		LastTotalPower: collections.NewItem(
			sb, types.LastTotalPowerKey,
			"last_total_power", sdk.IntValue,
		),
		Validators: collections.NewMap(
			sb,
			types.ValidatorsKey,
			"validators",
			sdk.LengthPrefixedBytesKey.WithName("validator_address"),
			codec.CollValue[types.Validator](cdc),
		),
		ValidatorByConsensusAddress: collections.NewMap(
			sb, types.ValidatorsByConsAddrKey,
			"validators_by_cons_addr",
			sdk.LengthPrefixedAddressKey(sdk.ConsAddressKey).WithName("cons_address"),
			collcodec.KeyToValueCodec(sdk.ValAddressKey),
		),

		LastValidatorPower: collections.NewMap(
			sb, types.LastValidatorPowerKey,
			"last_validator_power",
			sdk.LengthPrefixedBytesKey.
				WithName("validator_address"), codec.CollValue[gogotypes.Int64Value](cdc),
		), // sdk.LengthPrefixedBytesKey is needed to retain state compatibility

		// key format is: 67 | length(timestamp Bytes) | timestamp | height
		// Note: We use 3 keys here because we prefixed time bytes with its length previously and to retain state compatibility we remain to use the same
		ValidatorQueue: collections.NewMap(
			sb, types.ValidatorQueueKey,
			"validator_queue",
			collections.NamedTripleKeyCodec(
				"ts_length",
				collections.Uint64Key,
				"timestamp",
				sdk.TimeKey,
				"height",
				collections.Uint64Key,
			),
			codec.CollValue[types.ValAddresses](cdc),
		),

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

// ValidatorAddressCodec returns the validator address codec.
func (k Keeper) ValidatorAddressCodec() addresscodec.Codec {
	return k.validatorAddressCodec
}

func (k Keeper) Hooks() types.StakingHooks {
	if k.hooks == nil {
		return types.MultiStakingHooks{}
	}

	return k.hooks
}

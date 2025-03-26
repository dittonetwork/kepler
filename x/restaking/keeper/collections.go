package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

type Idx struct {
	// Emergency index by emergency status
	Emergency *indexes.Multi[bool, string, types.Validator]
}

func NewIndexes(sb *collections.SchemaBuilder) Idx {
	return Idx{
		Emergency: indexes.NewMulti(
			sb,
			types.KeyPrefixValidator,
			types.CollectionIndexValidatorsByEmergency,
			collections.BoolKey,
			collections.StringKey,
			func(_ string, val types.Validator) (bool, error) {
				return val.IsEmergency, nil
			}),
	}
}

func (a Idx) IndexesList() []collections.Index[string, types.Validator] {
	return []collections.Index[string, types.Validator]{
		a.Emergency,
	}
}

// SetValidator stores a validator in KVStore.
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) error {
	return k.ValidatorsMap.Set(ctx, validator.CosmosAddress, validator)
}

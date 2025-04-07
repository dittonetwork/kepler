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

	// Bonded index by bonded status
	Bonded *indexes.Multi[bool, string, types.Validator]
}

func NewIndexes(sb *collections.SchemaBuilder) Idx {
	return Idx{
		Emergency: indexes.NewMulti(
			sb,
			types.KeyPrefixValidators,
			"validators_by_emergency",
			collections.BoolKey,
			collections.StringKey,
			func(_ string, val types.Validator) (bool, error) {
				return val.IsEmergency, nil
			},
		),
		Bonded: indexes.NewMulti(
			sb,
			types.KeyPrefixValidators,
			"validators_by_bonded",
			collections.BoolKey,
			collections.StringKey,
			func(_ string, val types.Validator) (bool, error) {
				return val.Status == types.Bonded, nil
			},
		),
	}
}

func (a Idx) IndexesList() []collections.Index[string, types.Validator] {
	return []collections.Index[string, types.Validator]{
		a.Emergency,
		a.Bonded,
	}
}

// getAllValidators retrieves all validators from the store.
func (k Keeper) getAllValidators(ctx sdk.Context) ([]*types.Validator, error) {
	iter, err := k.validators.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	var validators []*types.Validator

	for ; iter.Valid(); iter.Next() {
		var validator types.Validator

		validator, err = iter.Value()
		if err != nil {
			return nil, err
		}

		validators = append(validators, &validator)
	}

	return validators, nil
}

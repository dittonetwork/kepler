package repository

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
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
			types.KeyPrefixEmergencyValidators,
			"validators_by_emergency",
			collections.BoolKey,
			collections.StringKey,
			func(_ string, val types.Validator) (bool, error) {
				return val.IsEmergency, nil
			},
		),
		Bonded: indexes.NewMulti(
			sb,
			types.KeyPrefixBondedValidators,
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

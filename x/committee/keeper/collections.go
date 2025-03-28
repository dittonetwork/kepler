package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	"github.com/dittonetwork/kepler/x/committee/types"
)

type Idx struct {
	// Emergency index by emergency status
	Emergency *indexes.Multi[bool, uint32, types.Committee]
}

func NewIndexes(sb *collections.SchemaBuilder) Idx {
	return Idx{
		Emergency: indexes.NewMulti(
			sb,
			types.CommitteesEmergencyIdxPrefix,
			"committees_by_emergency",
			collections.BoolKey,
			collections.Uint32Key,
			func(_ uint32, committee types.Committee) (bool, error) {
				return committee.IsEmergency, nil
			},
		),
	}
}

func (a Idx) IndexesList() []collections.Index[uint32, types.Committee] {
	return []collections.Index[uint32, types.Committee]{
		a.Emergency,
	}
}

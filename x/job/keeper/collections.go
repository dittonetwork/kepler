package keeper

import (
	"github.com/dittonetwork/kepler/x/job/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
)

type Idx struct {
	// AutomationID index by id of automation
	AutomationID *indexes.Multi[uint64, uint64, types.Job]
}

func NewJobIndexes(sb *collections.SchemaBuilder) Idx {
	return Idx{
		AutomationID: indexes.NewMulti(
			sb,
			types.JobsPrefix,
			collectionIndexJobByAutomationID,
			collections.Uint64Key,
			collections.Uint64Key,
			func(_ uint64, val types.Job) (uint64, error) {
				return val.AutomationId, nil
			}),
	}
}

func (i Idx) IndexesList() []collections.Index[uint64, types.Job] {
	return []collections.Index[uint64, types.Job]{
		i.AutomationID,
	}
}

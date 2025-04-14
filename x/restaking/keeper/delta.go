package keeper

import (
	"github.com/dittonetwork/kepler/x/restaking/types"
)

type validatorChanges struct {
	Created []*types.Operator
	Updated []*operatorUpdate
	Deleted []*types.Operator
}

type operatorUpdate struct {
	Before, After *types.Operator
}

// calculateOperatorsDelta compares two slices of validators and returns the changes.
func calculateOperatorsDelta(oldVals, newVals []*types.Operator) validatorChanges {
	changes := validatorChanges{
		Created: make([]*types.Operator, 0),
		Updated: make([]*operatorUpdate, 0),
		Deleted: make([]*types.Operator, 0),
	}

	// create a map of oldVals validators for easy lookup
	oldMap := make(map[string]*types.Operator)
	for _, operator := range oldVals {
		oldMap[operator.Address] = operator
	}

	// create a map of newVals validators for easy lookup
	newMap := make(map[string]*types.Operator)
	for _, operator := range newVals {
		newMap[operator.Address] = operator

		// check if the validator is in the oldVals map
		if oldOperator, exists := oldMap[operator.Address]; exists {
			// if the validator is in both maps, check for changes
			if !oldOperator.Equal(operator) {
				changes.Updated = append(changes.Updated, &operatorUpdate{
					Before: oldOperator,
					After:  operator,
				})
			}
		} else {
			// if the validator is not in the oldVals map, it is a newVals validator
			changes.Created = append(changes.Created, operator)
		}
	}

	// find deleted validators
	for address, operator := range oldMap {
		if _, exists := newMap[address]; !exists {
			changes.Deleted = append(changes.Deleted, operator)
		}
	}

	return changes
}

package keeper

import (
	"github.com/dittonetwork/kepler/x/restaking/types"
)

type validatorChanges struct {
	Created []*types.Validator
	Updated []*validatorUpdate
	Deleted []*types.Validator
}

type validatorUpdate struct {
	Before, After *types.Validator
}

// calculateValidatorDelta compares two slices of validators and returns the changes.
func calculateValidatorDelta(oldVals, newVals []*types.Validator) validatorChanges {
	changes := validatorChanges{
		Created: make([]*types.Validator, 0),
		Updated: make([]*validatorUpdate, 0),
		Deleted: make([]*types.Validator, 0),
	}

	// create a map of oldVals validators for easy lookup
	oldMap := make(map[string]*types.Validator)
	for _, validator := range oldVals {
		oldMap[validator.OperatorAddress] = validator
	}

	// create a map of newVals validators for easy lookup
	newMap := make(map[string]*types.Validator)
	for _, validator := range newVals {
		newMap[validator.OperatorAddress] = validator

		// check if the validator is in the oldVals map
		if oldValidator, exists := oldMap[validator.OperatorAddress]; exists {
			// if the validator is in both maps, check for changes
			if !oldValidator.Equal(validator) {
				changes.Updated = append(changes.Updated, &validatorUpdate{
					Before: oldValidator,
					After:  validator,
				})
			}
		} else {
			// if the validator is not in the oldVals map, it is a newVals validator
			changes.Created = append(changes.Created, validator)
		}
	}

	// find deleted validators
	for address, validator := range oldMap {
		if _, exists := newMap[address]; !exists {
			changes.Deleted = append(changes.Deleted, validator)
		}
	}

	return changes
}

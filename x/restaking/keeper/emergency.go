package keeper

import (
	"cosmossdk.io/collections/indexes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// GetActiveEmergencyValidators returns the list of active (bonded) emergency validators.
// @TODO: return an error.
func (k Keeper) GetActiveEmergencyValidators(ctx sdk.Context) []types.Validator {
	iter, err := k.validators.Indexes.Emergency.Iterate(ctx, nil)
	if err != nil {
		k.Logger().With("error", err).Error("failed to get emergency validators iterator")
		return []types.Validator{}
	}
	defer iter.Close()

	validators, err := indexes.CollectValues(ctx, k.validators, iter)
	if err != nil {
		k.Logger().With("error", err).Error("failed to collect emergency validators")
		return []types.Validator{}
	}

	activeEmergencyValidators := make([]types.Validator, 0, len(validators))
	for _, val := range validators {
		if val.Status != types.Bonded {
			continue
		}

		activeEmergencyValidators = append(activeEmergencyValidators, val)
	}

	return activeEmergencyValidators
}

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// GetActiveEmergencyValidators returns the list of active (bonded) emergency validators.
func (k Keeper) GetActiveEmergencyValidators(ctx sdk.Context) ([]types.Validator, error) {
	validators, err := k.repository.GetEmergencyValidators(ctx)
	if err != nil {
		return nil, err
	}

	activeEmergencyValidators := make([]types.Validator, 0, len(validators))
	for _, val := range validators {
		if val.Status != types.Bonded {
			continue
		}

		activeEmergencyValidators = append(activeEmergencyValidators, val)
	}

	return activeEmergencyValidators, nil
}

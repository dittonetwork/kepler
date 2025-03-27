package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

const (
	activeEmergencyValidatorsCapacity = 16
)

// GetActiveEmergencyValidators returns the list of active (bonded) emergency validators.
// Returns an empty list if any error occurs during iteration.
func (k Keeper) GetActiveEmergencyValidators(ctx sdk.Context) []types.EmergencyValidator {
	iter, err := k.ValidatorsMap.Indexes.Emergency.Iterate(ctx, nil)
	if err != nil {
		k.Logger().With("error", err).Error("failed to get emergency validators iterator")
		return []types.EmergencyValidator{}
	}
	defer iter.Close()

	// Pre-allocate with small capacity to avoid reallocations
	validators := make([]types.EmergencyValidator, 0, activeEmergencyValidatorsCapacity)

	for ; iter.Valid(); iter.Next() {
		var addr string
		addr, err = iter.PrimaryKey()
		if err != nil {
			k.Logger().With("error", err).Error("error iterating emergency validator keys")
			continue
		}

		var valAddr sdk.ValAddress
		valAddr, err = sdk.ValAddressFromBech32(addr)
		if err != nil {
			k.Logger().With("error", err, "addr", addr).Error("invalid validator address format")
			continue
		}

		var validator stakingtypes.Validator
		validator, err = k.staking.GetValidator(ctx, valAddr)
		if err != nil {
			k.Logger().With("error", err, "addr", addr).Error("failed to get validator")
			continue
		}

		// Only include active validators (with Bonded status)
		if validator.Status != stakingtypes.Bonded {
			continue
		}

		validators = append(validators, types.EmergencyValidator{
			Address:     valAddr,
			VotingPower: validator.GetConsensusPower(sdk.DefaultPowerReduction),
		})
	}

	return validators
}

package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	blockHashLength = 66
)

// NeedValidatorsUpdate is a helper function to check if the validators need to be updated for TaskManager module.
func (k Keeper) NeedValidatorsUpdate(ctx sdk.Context, epoch int64) (bool, error) {
	// Get the last epoch number
	lastUpdate, err := k.lastUpdate.Get(ctx)
	if err != nil {
		return false, err
	}

	return lastUpdate.EpochNum < epoch, nil
}

// UpdateValidatorSet updates the validator set based on the provided updates.
func (k Keeper) UpdateValidatorSet(ctx sdk.Context, update types.ValidatorsUpdate) error {
	if err := k.validateUpdateValidatorSet(ctx, update); err != nil {
		return err
	}

	currentValidators, err := k.getAllValidators(ctx)
	if err != nil {
		return sdkerrors.Wrap(types.ErrUpdateValidator, "failed to get all validators")
	}

	var newValidators []*types.Validator
	for i := range update.Operators {
		newValidators = append(newValidators, &update.Operators[i])
	}

	delta := calculateValidatorDelta(currentValidators, newValidators)

	for _, validator := range delta.Created {
		// update or create the validator in pending pool
		if err = k.pendingValidators.Set(ctx, validator.OperatorAddress, *validator); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to set pending validator")
		}
	}

	for _, validator := range delta.Deleted {
		// delete the validator from pending pool
		if err = k.pendingValidators.Remove(ctx, validator.OperatorAddress); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to remove pending validator")
		}

		// delete the validator from the store
		if err = k.validators.Remove(ctx, validator.OperatorAddress); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to remove validator")
		}
	}

	for _, validator := range delta.Updated {
		// update the validator in the store
		if err = k.validators.Set(ctx, validator.After.OperatorAddress, *validator.After); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to set updated validator")
		}

		// call the AfterValidatorBonded hook if the validator is bonded
		if !validator.Before.IsBonded() && validator.After.IsBonded() {
			if err = k.hooks.AfterValidatorBonded(ctx, *validator.After); err != nil {
				return sdkerrors.Wrap(types.ErrUpdateValidator, "error in AfterValidatorBonded hook")
			}
		}

		// call the BeforeValidatorBeginUnbonding hook if the validator is unbonding
		if validator.Before.IsBonded() && validator.After.IsUnbonding() {
			if err = k.hooks.BeforeValidatorBeginUnbonding(ctx, *validator.After); err != nil {
				return sdkerrors.Wrap(types.ErrUpdateValidator, "error in BeforeValidatorBeginUnbonding hook")
			}
		}
	}

	return k.lastUpdate.Set(ctx, update.Info)
}

// validateUpdateValidatorSet validates the parameters for updating the validator set.
func (k Keeper) validateUpdateValidatorSet(ctx sdk.Context, update types.ValidatorsUpdate) error {
	// Check if the block height is higher than the last update
	lastUpdate, err := k.lastUpdate.Get(ctx)
	if err != nil {
		return err
	}

	if lastUpdate.BlockHeight >= update.Info.BlockHeight {
		return sdkerrors.Wrap(types.ErrUpdateValidator, "block height is lower than last update")
	}

	// Check if the block hash is valid
	if len(update.Info.BlockHash) != blockHashLength {
		return sdkerrors.Wrap(types.ErrUpdateValidator, "invalid block hash")
	}

	// Check if the epoch number is valid
	if update.Info.EpochNum <= 0 {
		return sdkerrors.Wrap(types.ErrUpdateValidator, "invalid epoch number")
	}

	if lastUpdate.EpochNum >= update.Info.EpochNum {
		return sdkerrors.Wrap(types.ErrUpdateValidator, "epoch number is lower than last update")
	}

	// Check if the validator are valid
	for _, validator := range update.Operators {
		if len(validator.OperatorAddress) == 0 {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "operator address is empty")
		}

		if !common.IsHexAddress(validator.OperatorAddress) {
			return sdkerrors.Wrap(
				types.ErrUpdateValidator,
				"operator address is not a valid Ethereum address",
			)
		}

		if validator.ConsensusPubkey == nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "consensus public key is empty")
		}
	}

	return nil
}

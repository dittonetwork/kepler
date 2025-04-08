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
	lastUpdate, err := k.repository.GetLastUpdate(ctx)
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

	delta, err := k.makeDeltaUpdates(ctx, update)
	if err != nil {
		return err
	}

	if err = k.processCreatedValidators(ctx, delta.Created); err != nil {
		return err
	}

	if err = k.processDeletedValidators(ctx, delta.Deleted); err != nil {
		return err
	}

	if err = k.processUpdatedValidators(ctx, delta.Updated); err != nil {
		return err
	}

	return k.repository.SetLastUpdate(ctx, update.Info)
}

// processCreatedValidators handles all newly created validators.
func (k Keeper) processCreatedValidators(ctx sdk.Context, validators []*types.Validator) error {
	for _, validator := range validators {
		if err := k.repository.SetPendingValidator(ctx, validator.OperatorAddress, *validator); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to set pending validator")
		}
	}
	return nil
}

// processDeletedValidators handles all validators that need to be deleted.
func (k Keeper) processDeletedValidators(ctx sdk.Context, validators []*types.Validator) error {
	for _, validator := range validators {
		if err := k.repository.RemovePendingValidator(ctx, validator.OperatorAddress); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to remove pending validator")
		}

		if err := k.repository.RemoveValidator(ctx, validator.OperatorAddress); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to remove validator")
		}
	}
	return nil
}

// processUpdatedValidators handles all validators that have been updated.
func (k Keeper) processUpdatedValidators(ctx sdk.Context, updates []*validatorUpdate) error {
	for _, update := range updates {
		if err := k.repository.SetValidator(ctx, update.After.OperatorAddress, *update.After); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to set updated validator")
		}

		if err := k.handleValidatorStatusChange(ctx, update); err != nil {
			return err
		}
	}
	return nil
}

// handleValidatorStatusChange triggers appropriate hooks based on validator status changes.
func (k Keeper) handleValidatorStatusChange(ctx sdk.Context, update *validatorUpdate) error {
	// Validator became bonded
	if !update.Before.IsBonded() && update.After.IsBonded() {
		if err := k.hooks.AfterValidatorBonded(ctx, *update.After); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "error in AfterValidatorBonded hook")
		}
	}

	// Validator began unbonding
	if update.Before.IsBonded() && update.After.IsUnbonding() {
		if err := k.hooks.BeforeValidatorBeginUnbonding(ctx, *update.After); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "error in BeforeValidatorBeginUnbonding hook")
		}
	}

	return nil
}

// makeDeltaUpdates retrieves current validators and calculates the delta with new validators.
func (k Keeper) makeDeltaUpdates(ctx sdk.Context, update types.ValidatorsUpdate) (validatorChanges, error) {
	allValidators, err := k.repository.GetAllValidators(ctx)
	if err != nil {
		return validatorChanges{}, sdkerrors.Wrap(types.ErrUpdateValidator, "failed to get all validators")
	}

	var newValidators []*types.Validator
	for i := range update.Operators {
		newValidators = append(newValidators, &update.Operators[i])
	}

	var currentValidators []*types.Validator

	for i := range allValidators {
		currentValidators = append(currentValidators, &allValidators[i])
	}

	return calculateValidatorDelta(currentValidators, newValidators), nil
}

// validateUpdateValidatorSet validates the parameters for updating the validator set.
func (k Keeper) validateUpdateValidatorSet(ctx sdk.Context, update types.ValidatorsUpdate) error {
	// Check if the block height is higher than the last update
	lastUpdate, err := k.repository.GetLastUpdate(ctx)
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

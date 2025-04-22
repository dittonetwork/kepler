package keeper

import (
	"errors"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
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

// processCreatedOperators handles all newly created validators.
func (k Keeper) processCreatedOperators(ctx sdk.Context, operators []*types.Operator) error {
	for _, operator := range operators {
		if err := k.repository.SetPendingOperator(ctx, operator.Address, *operator); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to set pending validator")
		}
	}
	return nil
}

// processDeletedOperators handles all validators that need to be deleted.
func (k Keeper) processDeletedOperators(ctx sdk.Context, operators []*types.Operator) error {
	for _, operator := range operators {
		if err := k.repository.RemovePendingOperator(ctx, operator.Address); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to remove pending operator")
		}

		validator, err := k.repository.GetValidatorByEvmAddr(ctx, operator.Address)
		if err != nil {
			if errors.Is(err, collections.ErrNotFound) {
				return nil
			}

			return sdkerrors.Wrap(types.ErrUpdateValidator, "failed to get validator by EVM address")
		}

		if err = k.repository.AddValidatorsChange(ctx, validator, types.ValidatorChangeTypeDelete); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to add validator change")
		}
	}
	return nil
}

// processUpdatedOperators handles all validators that have been updated.
func (k Keeper) processUpdatedOperators(ctx sdk.Context, updates []*operatorUpdate) error {
	for _, update := range updates {
		validator, err := k.repository.GetValidatorByEvmAddr(ctx, update.Before.Address)
		if err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "failed to get validator by EVM address")
		}

		if err = k.repository.AddValidatorsChange(ctx, validator, types.ValidatorChangeTypeUpdate); err != nil {
			return sdkerrors.Wrap(types.ErrUpdateValidator, "unable to add validator change")
		}
	}
	return nil
}

// makeDeltaUpdates retrieves current validators and calculates the delta with new validators.
func (k Keeper) makeDeltaUpdates(ctx sdk.Context, operators []types.Operator) (validatorChanges, error) {
	allValidators, err := k.repository.GetAllValidators(ctx)
	if err != nil {
		return validatorChanges{}, sdkerrors.Wrap(types.ErrUpdateValidator, "failed to get all validators")
	}

	var newOperators []*types.Operator
	for i := range operators {
		newOperators = append(newOperators, &operators[i])
	}

	var currentValidators []*types.Operator

	for i := range allValidators {
		currentValidators = append(currentValidators, allValidators[i].ConvertToOperator())
	}

	return calculateOperatorsDelta(currentValidators, newOperators), nil
}

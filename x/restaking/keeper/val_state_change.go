package keeper

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// ApplyAndReturnValidatorSetUpdates applies and return accumulated updates to the bonded validator set.
// Notes: we always return the full set of validators, even if they are not updated.
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	changes, err := k.repository.GetValidatorsChanges(sdkCtx)
	if err != nil {
		k.Logger().With("error", err).Error("failed to collect bonded validators")
		return nil, err
	}

	updates := make([]abci.ValidatorUpdate, 0, changes.Len())

	updates, err = k.handleCreatedValidators(sdkCtx, changes.Created, updates)
	if err != nil {
		return nil, err
	}

	updates, err = k.handleUpdatedValidators(sdkCtx, changes.Updated, updates)
	if err != nil {
		return nil, err
	}

	updates, err = k.handleDeletedValidators(sdkCtx, changes.Deleted, updates)
	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		k.logger.Info("no validator updates")
		return nil, nil
	}

	k.logger.With("list", updates).Info("updated validators")
	return updates, k.repository.PruneValidatorsChanges(sdkCtx)
}

// handleCreatedValidators processes created validators and appends updates.
func (k Keeper) handleCreatedValidators(
	ctx sdk.Context,
	validators []types.Validator,
	updates []abci.ValidatorUpdate,
) ([]abci.ValidatorUpdate, error) {
	for _, validator := range validators {
		tmProtoPk, err := validator.GetConsensusPubkey(k.cdc)
		if err != nil {
			k.Logger().With("error", err, "validator", validator.OperatorAddress).
				Error("failed to get consensus public key when creating validator")
			continue
		}

		update := abci.ValidatorUpdate{
			PubKey: tmProtoPk,
			Power:  validator.VotingPower,
		}
		updates = append(updates, update)

		err = k.repository.SetValidator(ctx, sdk.ValAddress(validator.OperatorAddress), validator)
		if err != nil {
			k.Logger().With("validator", validator.OperatorAddress, "err", err).
				Error("failed to set validator")
			return nil, err
		}
	}
	return updates, nil
}

// handleUpdatedValidators processes updated validators and appends updates.
func (k Keeper) handleUpdatedValidators(
	ctx sdk.Context,
	validators []types.Validator,
	updates []abci.ValidatorUpdate,
) ([]abci.ValidatorUpdate, error) {
	for _, validator := range validators {
		oldValidator, err := k.repository.GetValidator(ctx, sdk.ValAddress(validator.OperatorAddress))
		if err != nil {
			k.Logger().With("validator", validator.OperatorAddress, "err", err).
				Error("validator not found")
			continue
		}

		// Set voting power to 0 for non-bonded validators
		if validator.Status != types.Bonded {
			validator.VotingPower = 0
		}

		tmProtoPk, err := validator.GetConsensusPubkey(k.cdc)
		if err != nil {
			k.Logger().With("error", err, "validator", validator.OperatorAddress).
				Error("failed to get consensus public key when updating validator")
			continue
		}

		// Check if validator parameters changed
		if oldValidator.HasConsensusParamsChanges(&validator) {
			k.Logger().With("validator", validator.OperatorAddress).
				Info("validator has changed consensus params")

			update := abci.ValidatorUpdate{
				PubKey: tmProtoPk,
				Power:  validator.VotingPower,
			}
			updates = append(updates, update)
		}

		// Save validator to state
		err = k.repository.SetValidator(ctx, sdk.ValAddress(validator.OperatorAddress), validator)
		if err != nil {
			k.Logger().With("validator", validator.OperatorAddress, "err", err).
				Error("failed to set validator")
			return nil, err
		}
	}
	return updates, nil
}

// handleDeletedValidators processes deleted validators and appends updates.
func (k Keeper) handleDeletedValidators(
	ctx sdk.Context,
	validators []types.Validator,
	updates []abci.ValidatorUpdate,
) ([]abci.ValidatorUpdate, error) {
	for _, validator := range validators {
		tmProtoPk, err := validator.GetConsensusPubkey(k.cdc)
		if err != nil {
			k.Logger().With("error", err, "validator", validator.OperatorAddress).
				Error("failed to get consensus public key when deleting validator")
			continue
		}

		// Create validator update with zero power
		update := abci.ValidatorUpdate{
			PubKey: tmProtoPk,
			Power:  0,
		}
		updates = append(updates, update)

		// Remove validator from state
		err = k.repository.RemoveValidator(ctx, validator.OperatorAddress)
		if err != nil {
			k.Logger().With("validator", validator.OperatorAddress, "err", err).
				Error("failed to remove validator")
			return nil, err
		}
	}
	return updates, nil
}

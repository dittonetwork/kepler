package keeper

import (
	"errors"
	"math"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/log"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// NeedValidatorsUpdate is a helper function to check if the validators need to be updated for TaskManager module.
func (k Keeper) NeedValidatorsUpdate(ctx sdk.Context, epoch int64) (bool, error) {
	// Get the last epoch number
	lastUpdate, err := k.LastUpdate.Get(ctx)
	if err != nil {
		return false, err
	}

	return lastUpdate.Epoch < epoch, nil
}

// UpdateValidatorSet updates the validator set in the staking module and keeps a local copy
// of validators with additional metadata in the restaking module's store.
func (k Keeper) UpdateValidatorSet(ctx sdk.Context, params types.UpdateValidatorSetParams) error {
	ok, err := k.NeedValidatorsUpdate(ctx, params.EpochNumber)
	if err != nil {
		return err
	}

	if !ok {
		return sdkerrors.Wrap(types.ErrUpdateValidator, "update not needed")
	}

	lastUpdate, err := k.LastUpdate.Get(ctx)
	if err != nil {
		return err
	}

	// Ensure the block height is higher than the last update
	if lastUpdate.BlockHeight >= params.BlockHeight {
		return sdkerrors.Wrap(types.ErrUpdateValidator, "block height is lower than last update")
	}

	for _, operator := range params.Operators {
		logger := k.Logger().With(
			"operator", operator.Address,
			"status", operator.Status,
			"tokens", operator.Tokens,
		)

		// Convert operator address to cosmos address
		var cosmosAddr sdk.AccAddress
		cosmosAddr, err = sdk.AccAddressFromBech32(operator.Address)
		if err != nil {
			// Skip invalid addresses
			logger.Error("failed to convert operator address")
			continue
		}

		// Convert account address to validator address
		valAddr := sdk.ValAddress(cosmosAddr)

		// Process validator update
		err = k.processValidatorUpdate(ctx, operator, cosmosAddr, valAddr, logger)
		if err != nil {
			return err
		}

		logger.Info("validator updated")
	}

	// Update the last epoch number
	return k.LastUpdate.Set(ctx, types.LastUpdate{
		Epoch:       params.EpochNumber,
		Timestamp:   ctx.BlockTime(),
		BlockHeight: params.BlockHeight,
		BlockHash:   params.BlockHash,
	})
}

// processValidatorUpdate handles updating or creating a validator based on operator information.
func (k Keeper) processValidatorUpdate(
	ctx sdk.Context,
	operator types.Operator,
	cosmosAddr sdk.AccAddress,
	valAddr sdk.ValAddress,
	logger log.Logger,
) error {
	// Flag to track if this is a new validator
	isNewValidator := false

	// Get existing validator if any
	validator, err := k.staking.GetValidator(ctx, valAddr)
	if err != nil {
		// Handle validator not found
		if !errors.Is(err, stakingtypes.ErrNoValidatorFound) {
			return sdkerrors.Wrapf(types.ErrUpdateValidator, "failed to get validator: %s", err)
		}

		// Create new validator
		validator, isNewValidator = k.createNewValidator(operator, valAddr, logger)
		if validator.Status == stakingtypes.Unspecified {
			// Skip if validator creation failed
			return nil
		}

		logger.
			With("status", validator.Status, "isNewValidator", isNewValidator).
			Info("created new validator")
	} else {
		// Update public key
		err = k.updateValidatorPubKey(ctx, validator, operator, logger)
		if err != nil {
			return err
		}

		logger.Info("updated validator public key")
	}

	// Update validator status and tokens
	updateErr := k.updateValidatorStatusAndTokens(ctx, validator, operator, logger)
	if updateErr != nil {
		return updateErr
	}

	// Save validator to local store
	return k.saveValidatorToStore(ctx, operator, cosmosAddr, validator, isNewValidator, logger)
}

// createNewValidator creates a new validator from operator information.
func (k Keeper) createNewValidator(
	operator types.Operator,
	valAddr sdk.ValAddress,
	logger log.Logger,
) (stakingtypes.Validator, bool) {
	accountPubKeyBytes, err := sdk.GetFromBech32(operator.PublicKey, sdk.GetConfig().GetBech32AccountPubPrefix())
	if err != nil {
		// Skip invalid public keys
		logger.With("error", err).Error("failed to convert public key")
		return stakingtypes.Validator{Status: stakingtypes.Unspecified}, false
	}

	// Create public key from bytes
	pubKey := &ed25519.PubKey{Key: accountPubKeyBytes}

	// Use operator's public key to create the validator
	validator, err := stakingtypes.NewValidator(valAddr.String(), pubKey, stakingtypes.Description{})
	if err != nil {
		// Skip if validator creation fails
		logger.With("error", err).Error("failed to create validator")
		return stakingtypes.Validator{Status: stakingtypes.Unspecified}, false
	}

	// Initialize validator status
	validator.Status = stakingtypes.Unspecified
	logger.With("validator", validator).Debug("new validator created")
	return validator, true
}

// updateValidatorPubKey updates the public key of an existing validator.
func (k Keeper) updateValidatorPubKey(
	ctx sdk.Context,
	validator stakingtypes.Validator,
	operator types.Operator,
	logger log.Logger,
) error {
	config := sdk.GetConfig()
	validatorPubKeyBytes, err := sdk.GetFromBech32(
		operator.PublicKey,
		config.GetBech32ValidatorPubPrefix(),
	)
	if err != nil {
		logger.With("error", err).Error("failed to convert public key for existing validator")
		return nil
	}

	pubKey := &ed25519.PubKey{Key: validatorPubKeyBytes}

	// Serialize public key to Any
	anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		logger.With("error", err).Error("failed to encode public key to Any")
		return nil
	}

	// Update the validator's public key
	validator.ConsensusPubkey = anyPubKey
	logger.Debug("updating validator public key")

	// Save the validator with the updated public key
	return k.staking.SetValidator(ctx, validator)
}

// updateValidatorStatusAndTokens updates the status and tokens of a validator.
func (k Keeper) updateValidatorStatusAndTokens(
	ctx sdk.Context,
	validator stakingtypes.Validator,
	operator types.Operator,
	logger log.Logger,
) error {
	// Map operator status to Cosmos validator status
	cosmosStatus := operator.Status.ToStakingBondStatus()

	// For new validators, if they're marked as bonded, start with Unbonded status first
	if validator.Status == stakingtypes.Unspecified && cosmosStatus == stakingtypes.Bonded {
		validator.Status = stakingtypes.Unbonded
	} else {
		validator.Status = cosmosStatus
	}

	// Update validator tokens - safely convert to int64, checking for overflow
	if operator.Tokens > uint64(math.MaxInt64) {
		logger.Error("token amount too large for int64 conversion", "tokens", operator.Tokens)
		return nil
	}

	tokenAmount := sdk.TokensFromConsensusPower(int64(operator.Tokens), sdk.DefaultPowerReduction)
	validator.Tokens = tokenAmount

	if validator.Status == stakingtypes.Unbonding {
		// Set the unbonding completion time
		err := k.BeforeValidatorBeginUnbonding(ctx, validator)
		if err != nil {
			logger.With("error", err).Error("failed to set unbonding hook")
			return err
		}
	}

	// Set the updated validator in the staking module
	err := k.staking.SetValidator(ctx, validator)
	if err != nil {
		logger.With("error", err).Error("failed to update validator")
		return err
	}

	// Dispatch hook for validator update if the validator is bonded
	if validator.Status == stakingtypes.Bonded {
		return k.AfterValidatorBonded(ctx, validator)
	}

	return nil
}

// saveValidatorToStore saves the validator to the restaking module's store.
func (k Keeper) saveValidatorToStore(
	ctx sdk.Context,
	operator types.Operator,
	cosmosAddr sdk.AccAddress,
	validator stakingtypes.Validator,
	isNewValidator bool,
	logger log.Logger,
) error {
	power := sdk.TokensToConsensusPower(validator.Tokens, sdk.DefaultPowerReduction)

	// Ensure power can be safely converted to uint64
	var votingPower uint64
	if power >= 0 {
		votingPower = uint64(power)
	} else {
		votingPower = 0
		logger.Error("negative power value converted to zero", "power", power)
	}

	// Save validator to local store with the appropriate restaking validator status
	return k.SetValidator(ctx, types.Validator{
		Address:       operator.Address,
		CosmosAddress: cosmosAddr.String(),
		IsEmergency:   operator.IsEmergency,
		VotingPower:   votingPower,
		Status:        operator.Status.ToRestakingValidatorStatus(isNewValidator),
	})
}

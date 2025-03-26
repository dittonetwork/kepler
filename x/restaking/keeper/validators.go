package keeper

import (
	"errors"

	sdkerrors "cosmossdk.io/errors"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// UpdateValidatorSet updates the validator set in the staking module
func (k Keeper) UpdateValidatorSet(ctx sdk.Context, params types.UpdateValidatorSetParams) error {

	for _, operator := range params.Operators {
		log := k.Logger().With(
			"operator", operator.Address,
			"status", operator.Status,
			"tokens", operator.Tokens,
		)

		// Convert operator address to cosmos address
		cosmosAddr, err := sdk.AccAddressFromBech32(operator.Address)
		if err != nil {
			// Skip invalid addresses
			log.Error("failed to convert operator address")
			continue
		}

		// Convert account address to validator address
		valAddr := sdk.ValAddress(cosmosAddr)

		// Get existing validator if any
		validator, err := k.staking.GetValidator(ctx, valAddr)
		if err != nil {
			if errors.Is(err, stakingtypes.ErrNoValidatorFound) {
				// Create new validator if not found
				pubKeyBytes, err := sdk.GetFromBech32(operator.PublicKey, sdk.GetConfig().GetBech32AccountPubPrefix())
				if err != nil {
					// Skip invalid public keys
					log.Error("failed to convert public key", "error", err)
					continue
				}

				// Create public key from bytes
				pubKey := &ed25519.PubKey{Key: pubKeyBytes}

				// Use operator's public key to create the validator
				validator, err = stakingtypes.NewValidator(valAddr.String(), pubKey, stakingtypes.Description{})
				if err != nil {
					// Skip if validator creation fails
					log.Error("failed to create validator", "error", err)
					continue
				}

				validator.Status = stakingtypes.Unspecified
			} else {
				return sdkerrors.Wrapf(types.ErrUpdateValidator, "failed to get validator: %s", err)
			}
		}

		// If validator already exists, update its public key if needed
		pubKeyBytes, err := sdk.GetFromBech32(operator.PublicKey, sdk.GetConfig().GetBech32ValidatorPubPrefix())
		if err != nil {
			log.Error("failed to convert public key for existing validator", "error", err)
		} else {
			pubKey := &ed25519.PubKey{Key: pubKeyBytes}

			// Serialize public key to Any
			anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
			if err != nil {
				log.Error("failed to encode public key to Any", "error", err)
			} else {
				validator.ConsensusPubkey = anyPubKey
			}
		}

		// Update validator status based on operator status
		switch operator.Status {
		case types.OperatorStatusBonded:
			// Check if the validator is already bonded on kepler
			// If not, set the status to Unbonded
			if validator.Status == stakingtypes.Unspecified {
				validator.Status = stakingtypes.Unbonded
			} else {
				validator.Status = stakingtypes.Bonded
			}
		case types.OperatorStatusUnbonded:
			validator.Status = stakingtypes.Unbonded
		case types.OperatorStatusUnbonding:
			validator.Status = stakingtypes.Unbonding
		}

		// Update validator tokens
		tokenAmount := sdk.TokensFromConsensusPower(int64(operator.Tokens), sdk.DefaultPowerReduction)
		validator.Tokens = tokenAmount

		// Set the updated validator
		err = k.staking.SetValidator(ctx, validator)
		if err != nil {
			return err
		}

		log.Info("validator updated")
	}

	return nil
}

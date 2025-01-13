package keeper

import (
	"context"
	"slices"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	stakingtypes "cosmossdk.io/x/staking/types"
)

func (k Keeper) minCommissionRate(ctx context.Context) (math.LegacyDec, error) {
	params, err := k.stakingKeeper.Params.Get(ctx)
	return params.MinCommissionRate, err
}

func validatePubKey(pubKey cryptotypes.PubKey, knownPubKeyTypes []string) error {
	pubKeyType := pubKey.Type()

	if !slices.Contains(knownPubKeyTypes, pubKeyType) {
		return errorsmod.Wrapf(
			stakingtypes.ErrValidatorPubKeyTypeNotSupported,
			"got: %s, expected: %s", pubKey.Type(), knownPubKeyTypes,
		)
	}

	if pubKeyType == sdk.PubKeyEd25519Type {
		if len(pubKey.Bytes()) != ed25519.PubKeySize {
			return errorsmod.Wrapf(
				stakingtypes.ErrConsensusPubKeyLenInvalid,
				"invalid Ed25519 pubkey size: got %d, expected %d", len(pubKey.Bytes()), ed25519.PubKeySize,
			)
		}
	}

	return nil
}

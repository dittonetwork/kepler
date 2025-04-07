package types

import (
	"cosmossdk.io/errors"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IsBonded helper function to check if the validator is bonded.
func (v Validator) IsBonded() bool {
	return v.Status == Bonded
}

// IsUnbonding helper function to check if the validator is unbonding.
func (v Validator) IsUnbonding() bool {
	return v.Status == Unbonding
}

// ConsPubKey returns the validator's PubKey as cryptotypes.PubKey.
func (v Validator) ConsPubKey() (cryptotypes.PubKey, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrapf(
			sdkerrors.ErrInvalidType,
			"expecting cryptotypes.PubKey, got %T", v.ConsensusPubkey,
		)
	}

	return pk, nil
}

// CmtConsPublicKey casts Validator.ConsPublicKey to cmtprotocrypto.PublicKey.
func (v Validator) CmtConsPublicKey() (cmtprotocrypto.PublicKey, error) {
	pk, err := v.ConsPubKey()
	if err != nil {
		return cmtprotocrypto.PublicKey{}, err
	}

	tmPk, err := cryptocodec.ToCmtProtoPublicKey(pk)
	if err != nil {
		return cmtprotocrypto.PublicKey{}, err
	}

	return tmPk, nil
}

type EmergencyValidator struct {
	Address     sdk.ValAddress
	VotingPower int64
}

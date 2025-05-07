package types

import (
	"cosmossdk.io/errors"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MaxMonikerLength         = 70
	MaxIdentityLength        = 3000
	MaxWebsiteLength         = 140
	MaxSecurityContactLength = 140
	MaxDetailsLength         = 280
)

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
		return cmtprotocrypto.PublicKey{}, errors.Wrapf(err, "failed to get consensus public key")
	}

	tmPk, err := cryptocodec.ToCmtProtoPublicKey(pk)
	if err != nil {
		return cmtprotocrypto.PublicKey{}, err
	}

	return tmPk, nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces.
func (v Validator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pk cryptotypes.PubKey
	return unpacker.UnpackAny(v.ConsensusPubkey, &pk)
}

// ConvertToOperator converts a Validator to an Operator.
func (v Validator) ConvertToOperator() *Operator {
	return &Operator{
		Address:         v.EvmOperatorAddress,
		ConsensusPubkey: v.ConsensusPubkey,
		IsEmergency:     v.IsEmergency,
		Status:          v.Status,
		VotingPower:     v.VotingPower,
		Protocol:        v.Protocol,
	}
}

func (v Validator) GetConsensusPubkey(unpacker codectypes.AnyUnpacker) (cmtprotocrypto.PublicKey, error) {
	var pk cryptotypes.PubKey
	if err := unpacker.UnpackAny(v.ConsensusPubkey, &pk); err != nil {
		return cmtprotocrypto.PublicKey{}, errors.Wrapf(err, "failed to unpack consensus public key")
	}

	return v.CmtConsPublicKey()
}

// Len represents the changes to be applied to validators.
func (c ValidatorsChanges) Len() int {
	return len(c.Created) + len(c.Updated) + len(c.Deleted)
}

// GetCreatedAndUpdated returns a slice of validators that were either created or updated.
func (c ValidatorsChanges) GetCreatedAndUpdated() []Validator {
	createdAndUpdated := make([]Validator, 0, len(c.Created)+len(c.Updated))
	createdAndUpdated = append(createdAndUpdated, c.Created...)
	createdAndUpdated = append(createdAndUpdated, c.Updated...)

	return createdAndUpdated
}

// DoNotModifyDesc constant used in flags to indicate that description field should not be updated.
const DoNotModifyDesc = "[do-not-modify]"

func NewDescription(moniker, identity, website, securityContact, details string) Description {
	return Description{
		Moniker:         moniker,
		Identity:        identity,
		Website:         website,
		SecurityContact: securityContact,
		Details:         details,
	}
}

// UpdateDescription updates the fields of a given description. An error is
// returned if the resulting description contains an invalid length.
func (d Description) UpdateDescription(d2 Description) (Description, error) {
	if d2.Moniker == DoNotModifyDesc {
		d2.Moniker = d.Moniker
	}

	if d2.Identity == DoNotModifyDesc {
		d2.Identity = d.Identity
	}

	if d2.Website == DoNotModifyDesc {
		d2.Website = d.Website
	}

	if d2.SecurityContact == DoNotModifyDesc {
		d2.SecurityContact = d.SecurityContact
	}

	if d2.Details == DoNotModifyDesc {
		d2.Details = d.Details
	}

	return NewDescription(
		d2.Moniker,
		d2.Identity,
		d2.Website,
		d2.SecurityContact,
		d2.Details,
	).EnsureLength()
}

// EnsureLength ensures the length of a validator's description.
func (d Description) EnsureLength() (Description, error) {
	if len(d.Moniker) > MaxMonikerLength {
		return d, errors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid moniker length; got: %d, max: %d", len(d.Moniker), MaxMonikerLength,
		)
	}

	if len(d.Identity) > MaxIdentityLength {
		return d, errors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid identity length; got: %d, max: %d", len(d.Identity), MaxIdentityLength,
		)
	}

	if len(d.Website) > MaxWebsiteLength {
		return d, errors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid website length; got: %d, max: %d", len(d.Website), MaxWebsiteLength,
		)
	}

	if len(d.SecurityContact) > MaxSecurityContactLength {
		return d, errors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid security contact length; got: %d, max: %d", len(d.SecurityContact), MaxSecurityContactLength,
		)
	}

	if len(d.Details) > MaxDetailsLength {
		return d, errors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid details length; got: %d, max: %d", len(d.Details), MaxDetailsLength,
		)
	}

	return d, nil
}

// HasConsensusParamsChanges checks if the operator has any changes in consensus parameters.
func (v Validator) HasConsensusParamsChanges(change *Validator) bool {
	if !v.ConsensusPubkey.Equal(change.ConsensusPubkey) {
		return true
	}

	if v.VotingPower != change.VotingPower {
		return true
	}

	return false
}

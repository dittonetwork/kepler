package types

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"net/url"
	"time"
)

const (
	// TODO: Why can't we just have one string description which can be JSON by convention
	MaxMonikerLength         = 70
	MaxIdentityLength        = 3000
	MaxWebsiteLength         = 140
	MaxSecurityContactLength = 140
	MaxDetailsLength         = 280

	DoNotModifyDesc = "[do-not-modify]"
)

func NewValidator(operator string, pubKey cryptotypes.PubKey, description Description) (Validator, error) {
	pkAny, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		return Validator{}, err
	}

	return Validator{
		OperatorAddress: operator,
		ConsensusPubkey: pkAny,
		Jailed:          false,
		Status:          Unbonded,
		Tokens:          math.ZeroInt(),
		UnbondingHeight: int64(0),
		UnbondingTime:   time.Unix(0, 0).UTC(),
		Description:     description,
	}, nil
}

func NewDescription(moniker, identity, website, securityContact, details string, metadata *Metadata) Description {
	return Description{
		Moniker:         moniker,
		Identity:        identity,
		Website:         website,
		SecurityContact: securityContact,
		Details:         details,
		Metadata:        metadata,
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

	if d2.Metadata != nil {
		if d2.Metadata.ProfilePicUri == DoNotModifyDesc {
			d2.Metadata.ProfilePicUri = d.Metadata.ProfilePicUri
		}
	}

	return NewDescription(
		d2.Moniker,
		d2.Identity,
		d2.Website,
		d2.SecurityContact,
		d2.Details,
		d.Metadata,
	).Validate()
}

func (v Validator) IsUnbonding() bool {
	return v.Status == Unbonding
}

func (v Validator) IsUnbonded() bool {
	return v.Status == Unbonded
}

func (v Validator) IsBonded() bool {
	return v.Status == Bonded
}

func (v Validator) IsBonding() bool {
	return v.Status == Bonding
}

func (v Validator) IsJailed() bool {
	return v.Jailed
}

func (v Validator) GetTokens() math.Int {
	return v.Tokens
}

func (v Validator) GetMoniker() string {
	return v.Description.Moniker
}

func (v Validator) GetStatus() sdk.BondStatus {
	switch v.Status {
	case Unbonded:
		return sdk.Unbonded
	case Unbonding:
		return sdk.Unbonding
	case Bonded:
		return sdk.Bonded
	case Bonding:
		return sdk.Unbonded
	default:
		return sdk.Unspecified
	}
}

func (v Validator) GetBondedTokens() math.Int {
	if v.IsBonded() {
		return v.Tokens
	}

	return math.ZeroInt()
}

func (v Validator) GetConsensusPower(r math.Int) int64 {
	return v.ConsensusPower(r)
}

// GetCommission returns the commission of the validator
// Deprecated: use for interface compatibility with sdk.ValidatorI
func (v Validator) GetCommission() math.LegacyDec {
	return math.ZeroInt().ToLegacyDec()
}

// GetMinSelfDelegation returns the minimum self delegation of the validator
// Deprecated: use for interface compatibility with sdk.ValidatorI
func (v Validator) GetMinSelfDelegation() math.Int {
	return math.ZeroInt()
}

// GetDelegatorShares returns the total amount of delegator shares
// Deprecated: use for interface compatibility with sdk.ValidatorI
func (v Validator) GetDelegatorShares() math.LegacyDec {
	return v.Tokens.ToLegacyDec()
}

// TokensFromShares returns the amount of tokens a delegation of the given shares
// Deprecated: use for interface compatibility with sdk.ValidatorI
func (v Validator) TokensFromShares(shares math.LegacyDec) math.LegacyDec {
	return shares
}

// TokensFromSharesTruncated returns the amount of tokens a delegation of the given shares
// Deprecated: use for interface compatibility with sdk.ValidatorI
func (v Validator) TokensFromSharesTruncated(dec math.LegacyDec) math.LegacyDec {
	return dec
}

// TokensFromSharesRoundUp returns the amount of tokens a delegation of the given shares
// Deprecated: use for interface compatibility with sdk.ValidatorI
func (v Validator) TokensFromSharesRoundUp(dec math.LegacyDec) math.LegacyDec {
	return dec
}

// SharesFromTokens returns the amount of shares a delegation of the given tokens
// Deprecated: use for interface compatibility with sdk.ValidatorI
func (v Validator) SharesFromTokens(tokens math.Int) (math.LegacyDec, error) {
	return v.Tokens.ToLegacyDec(), nil
}

// SharesFromTokensTruncated returns the amount of shares a delegation of the given tokens
// Deprecated: use for interface compatibility with sdk.ValidatorI
func (v Validator) SharesFromTokensTruncated(tokens math.Int) (math.LegacyDec, error) {
	return tokens.ToLegacyDec(), nil
}

func (v Validator) GetOperator() string {
	return v.OperatorAddress
}

// PotentialConsensusPower returns the potential consensus power of the validator
func (v Validator) PotentialConsensusPower(r math.Int) int64 {
	return sdk.TokensToConsensusPower(v.Tokens, r)
}

// UpdateStatus updates the location of the shares within a validator
// to reflect the new status
func (v *Validator) UpdateStatus(newStatus BondStatus) {
	v.Status = newStatus
}

// GetConsAddr extract Consensus key address
func (v Validator) GetConsAddr() ([]byte, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return pk.Address().Bytes(), nil
}

// ConsensusPower gets the consensus-engine power. Aa reduction of 10^6 from
// validator tokens is applied
func (v Validator) ConsensusPower(r math.Int) int64 {
	if v.IsBonded() {
		return v.PotentialConsensusPower(r)
	}

	return 0
}

// ConsPubKey returns the validator PubKey as a cryptotypes.PubKey.
func (v Validator) ConsPubKey() (cryptotypes.PubKey, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return pk, nil
}

// ModuleValidatorUpdate returns a appmodule.ValidatorUpdate from a staking validator type
// with the full validator power.
// It replaces the previous ABCIValidatorUpdate function.
func (v Validator) ModuleValidatorUpdate(r math.Int) appmodule.ValidatorUpdate {
	consPk, err := v.ConsPubKey()
	if err != nil {
		panic(err)
	}

	return appmodule.ValidatorUpdate{
		PubKey:     consPk.Bytes(),
		PubKeyType: consPk.Type(),
		Power:      v.ConsensusPower(r),
	}
}

// ModuleValidatorUpdateZero returns a appmodule.ValidatorUpdate from a staking validator type
// with zero power used for validator updates.
// It replaces the previous ABCIValidatorUpdateZero function.
func (v Validator) ModuleValidatorUpdateZero() appmodule.ValidatorUpdate {
	consPk, err := v.ConsPubKey()
	if err != nil {
		panic(err)
	}

	return appmodule.ValidatorUpdate{
		PubKey:     consPk.Bytes(),
		PubKeyType: consPk.Type(),
		Power:      0,
	}
}

// EnsureLength ensures the length of a validator's description.
func (d Description) EnsureLength() (Description, error) {
	if len(d.Moniker) > MaxMonikerLength {
		return d, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid moniker length; got: %d, max: %d", len(d.Moniker), MaxMonikerLength)
	}

	if len(d.Identity) > MaxIdentityLength {
		return d, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid identity length; got: %d, max: %d", len(d.Identity), MaxIdentityLength)
	}

	if len(d.Website) > MaxWebsiteLength {
		return d, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid website length; got: %d, max: %d", len(d.Website), MaxWebsiteLength)
	}

	if len(d.SecurityContact) > MaxSecurityContactLength {
		return d, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid security contact length; got: %d, max: %d", len(d.SecurityContact), MaxSecurityContactLength)
	}

	if len(d.Details) > MaxDetailsLength {
		return d, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid details length; got: %d, max: %d", len(d.Details), MaxDetailsLength)
	}

	return d, nil
}

func (d Description) IsEmpty() bool {
	return d.Moniker == "" && d.Details == "" && d.Identity == "" && d.Website == "" && d.SecurityContact == "" &&
		(d.Metadata == nil || d.Metadata.ProfilePicUri == "" && len(d.Metadata.SocialHandleUris) == 0)
}

// Validate calls metadata.Validate() description.EnsureLength()
func (d Description) Validate() (Description, error) {
	if d.Metadata != nil {
		if err := d.Metadata.Validate(); err != nil {
			return d, err
		}
	}

	return d.EnsureLength()
}

// Validate checks that the metadata fields are valid. For the ProfilePicUri, checks if a valid URI.
func (m Metadata) Validate() error {
	if m.ProfilePicUri != "" {
		_, err := url.ParseRequestURI(m.ProfilePicUri)
		if err != nil {
			return errors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"invalid profile_pic_uri format: %s, err: %s", m.ProfilePicUri, err,
			)
		}
	}

	if m.SocialHandleUris != nil {
		for _, socialHandleUri := range m.SocialHandleUris {
			_, err := url.ParseRequestURI(socialHandleUri)
			if err != nil {
				return errors.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"invalid social_handle_uri: %s, err: %s", socialHandleUri, err,
				)
			}
		}
	}
	return nil
}

// unmarshal a redelegation from a store value
func UnmarshalValidator(cdc codec.BinaryCodec, value []byte) (v Validator, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}

package types

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (v *Validator) IsUnbonding() bool {
	return v.Status == Unbonding
}

func (v *Validator) IsUnbonded() bool {
	return v.Status == Unbonded
}

func (v *Validator) IsBonded() bool {
	return v.Status == Bonded
}

func (v *Validator) IsBonding() bool {
	return v.Status == Bonding
}

func (v *Validator) IsJailed() bool {
	return v.Jailed
}

func (v *Validator) GetTokens() math.Int {
	return v.Tokens
}

func (v *Validator) GetOperator() string {
	return v.OperatorAddress
}

// PotentialConsensusPower returns the potential consensus power of the validator
func (v *Validator) PotentialConsensusPower(r math.Int) int64 {
	return sdk.TokensToConsensusPower(v.Tokens, r)
}

// UpdateStatus updates the location of the shares within a validator
// to reflect the new status
func (v *Validator) UpdateStatus(newStatus BondStatus) {
	v.Status = newStatus
}

// GetConsAddr extract Consensus key address
func (v *Validator) GetConsAddr() ([]byte, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return pk.Address().Bytes(), nil
}

// ConsensusPower gets the consensus-engine power. Aa reduction of 10^6 from
// validator tokens is applied
func (v *Validator) ConsensusPower(r math.Int) int64 {
	if v.IsBonded() {
		return v.PotentialConsensusPower(r)
	}

	return 0
}

// ConsPubKey returns the validator PubKey as a cryptotypes.PubKey.
func (v *Validator) ConsPubKey() (cryptotypes.PubKey, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return pk, nil
}

// ModuleValidatorUpdate returns a appmodule.ValidatorUpdate from a staking validator type
// with the full validator power.
// It replaces the previous ABCIValidatorUpdate function.
func (v *Validator) ModuleValidatorUpdate(r math.Int) appmodule.ValidatorUpdate {
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
func (v *Validator) ModuleValidatorUpdateZero() appmodule.ValidatorUpdate {
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

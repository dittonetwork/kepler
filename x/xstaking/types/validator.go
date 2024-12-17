package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

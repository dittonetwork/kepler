package types

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

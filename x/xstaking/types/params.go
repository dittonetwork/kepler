package types

import "time"

const (
	// DefaultUnbondingTime is the default time duration for unbonding
	// unbonding time
	DefaultUnbondingTime time.Duration = time.Hour * 24 * 7 * 3

	// DefaultMaxValidators is the default maximum number of validators
	DefaultMaxValidators uint32 = 100
)

// NewParams creates a new Params instance.
func NewParams(unbondingTime time.Duration, maxValidators uint32) Params {
	return Params{
		UnbondingTime: unbondingTime,
		MaxValidators: maxValidators,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultUnbondingTime,
		DefaultMaxValidators,
	)
}

// Validate validates the set of params.
func (p Params) Validate() error {

	return nil
}

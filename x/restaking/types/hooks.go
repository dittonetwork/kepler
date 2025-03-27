package types

import (
	"context"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type RestakingHooks interface {
	// AfterValidatorBonded is called after a validator is bonded.
	AfterValidatorBonded(ctx context.Context, validator stakingtypes.ValidatorI) error

	// BeforeValidatorBeginUnbonding is called before a validator begins unbonding.
	BeforeValidatorBeginUnbonding(ctx context.Context, validator stakingtypes.ValidatorI) error
}

var _ RestakingHooks = MultiRestakingHooks{}

type MultiRestakingHooks []RestakingHooks

// NewMultiRestakingHooks creates a new MultiRestakingHooks object.
func NewMultiRestakingHooks(hooks ...RestakingHooks) MultiRestakingHooks {
	return hooks
}

// AfterValidatorBonded is called after a validator is bonded.
func (m MultiRestakingHooks) AfterValidatorBonded(ctx context.Context, validator stakingtypes.ValidatorI) error {
	for i := range m {
		if err := m[i].AfterValidatorBonded(ctx, validator); err != nil {
			return err
		}
	}
	return nil
}

// BeforeValidatorBeginUnbonding is called before a validator begins unbonding.
func (m MultiRestakingHooks) BeforeValidatorBeginUnbonding(
	ctx context.Context,
	validator stakingtypes.ValidatorI,
) error {
	for i := range m {
		if err := m[i].BeforeValidatorBeginUnbonding(ctx, validator); err != nil {
			return err
		}
	}
	return nil
}

// RestakingHooksWrapper is a wrapper for modules to inject RestakingHooks using depinject.
type RestakingHooksWrapper struct{ RestakingHooks }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (RestakingHooksWrapper) IsOnePerModuleType() {}

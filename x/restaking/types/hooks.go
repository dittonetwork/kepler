package types

import (
	"context"
)

type RestakingHooks interface {
	// BeforeValidatorBeginUnbonding is called before a validator begins unbonding.
	BeforeValidatorBeginUnbonding(ctx context.Context, validator Validator) error
}

var _ RestakingHooks = MultiRestakingHooks{}

type MultiRestakingHooks []RestakingHooks

// NewMultiRestakingHooks creates a new MultiRestakingHooks object.
func NewMultiRestakingHooks(hooks ...RestakingHooks) MultiRestakingHooks {
	return hooks
}

// BeforeValidatorBeginUnbonding is called before a validator begins unbonding.
func (m MultiRestakingHooks) BeforeValidatorBeginUnbonding(
	ctx context.Context,
	validator Validator,
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

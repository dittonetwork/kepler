package types

import (
	"context"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MultiStakingHooks []StakingHooks

func NewMultiStakingHooks(hooks ...StakingHooks) MultiStakingHooks {
	return hooks
}

func (h MultiStakingHooks) AfterValidatorCreated(ctx context.Context, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterValidatorCreated(ctx, valAddr); err != nil {
			return err
		}
	}

	return nil
}

func (h MultiStakingHooks) BeforeValidatorModified(ctx context.Context, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].BeforeValidatorModified(ctx, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) AfterValidatorRemoved(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterValidatorRemoved(ctx, consAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) AfterValidatorBonded(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterValidatorBonded(ctx, consAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) AfterValidatorBeginUnbonding(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterValidatorBeginUnbonding(ctx, consAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) BeforeValidatorSlashed(ctx context.Context, valAddr sdk.ValAddress, fraction math.LegacyDec) error {
	for i := range h {
		if err := h[i].BeforeValidatorSlashed(ctx, valAddr, fraction); err != nil {
			return err
		}
	}
	return nil
}

package types

import (
	"context"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event Hooks
// These can be utilized to communicate between a staking keeper and another
// keeper which must take particular actions when validators/delegators change
// state. The second keeper must implement this interface, which then the
// staking keeper can call.

// StakingHooks event hooks for staking validator object (noalias)
type StakingHooks interface {
	// AfterValidatorCreated - must be called when a validator is created
	AfterValidatorCreated(ctx context.Context, valAddr sdk.ValAddress) error
	// BeforeValidatorModified - must be called when a validator's state changes
	BeforeValidatorModified(ctx context.Context, valAddr sdk.ValAddress) error
	// AfterValidatorRemoved - must be called when a validator is deleted
	AfterValidatorRemoved(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error

	// AfterValidatorBonded - must be called when a validator is bonded
	AfterValidatorBonded(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error
	// AfterValidatorBeginUnbonding - must be called when a validator begins unbonding
	AfterValidatorBeginUnbonding(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error

	// BeforeValidatorSlashed - must be called when a validator is slashed
	BeforeValidatorSlashed(ctx context.Context, valAddr sdk.ValAddress, fraction math.LegacyDec) error
}

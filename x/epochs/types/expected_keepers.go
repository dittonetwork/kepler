package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

// EpochHooks is event hooks.
// These can be utilized to communicate between the epochs keeper and another
// keeper which must take particular actions.
type EpochHooks interface {
	// AfterEpochEnd the first block whose timestamp is after the duration is counted as the end of the epoch.
	AfterEpochEnd(ctx context.Context, epochID string, epochNumber int64) error

	// BeforeEpochStart new epoch is next block of epoch end block.
	BeforeEpochStart(ctx context.Context, epochID string, epochNumber int64) error
}

// EpochHooksWrapper is a wrapper for modules to inject EpochHooks using depinject.
type EpochHooksWrapper struct{ EpochHooks }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (EpochHooksWrapper) IsOnePerModuleType() {}

package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	exectypes "github.com/dittonetwork/kepler/x/executors/types"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"
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

// ExecutorsKeeper defines the expected interface for the ExecutorsKeeper module.
type ExecutorsKeeper interface {
	GetEmergencyExecutors(ctx sdk.Context) ([]exectypes.Executor, error)
}

// RestakingKeeper defines the expected interface for the Validator module.
type RestakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (restakingtypes.Validator, error)
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

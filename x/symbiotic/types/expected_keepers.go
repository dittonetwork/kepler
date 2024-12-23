package types

import (
	"context"
	"kepler/x/beacon/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BeaconKeeper interface {
	// Add methods imported from beacon should be defined here
	SyncNeeded(ctx context.Context) bool
	GetFinalizedBlockInfo(ctx context.Context) (val types.FinalizedBlockInfo, found bool)
}

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

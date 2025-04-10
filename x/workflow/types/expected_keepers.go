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

type CommitteeKeeper interface {
	IsCommitteeExists(ctx sdk.Context, committeeID string) (bool, error)
	CanBeSigned(
		ctx sdk.Context,
		committeeID string,
		chainID string,
		signatures [][]byte,
		payload []byte,
	) (bool, error)
}

type JobKeeper interface {
	CreateJob(
		ctx sdk.Context,
		status any,
		committeeID string,
		chainID string,
		automationID uint64,
		txHash string,
		executorAddress string,
		createdAt uint64,
		executedAt uint64,
		signedAt uint64,
		signs [][]byte,
		payload []byte,
	) error
	GetLastSuccessfulJobByAutomation(ctx sdk.Context, automationID uint64) (any, error)
}

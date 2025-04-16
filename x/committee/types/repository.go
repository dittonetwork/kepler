package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Repository interface {
	GetLastEpoch(ctx sdk.Context) (int64, error)
	SetLastEpoch(ctx sdk.Context, epoch int64) error

	GetCommittee(ctx sdk.Context, epoch int64) (Committee, error)
	GetLastCommittee(ctx sdk.Context) (Committee, error)
	SetCommittee(ctx sdk.Context, epoch int64, committee Committee) error
	HasCommittee(ctx sdk.Context, epoch int64) (bool, error)
	IterateCommittees(ctx sdk.Context, fn func(committee Committee) error) error
}

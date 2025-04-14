package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Repository interface {
	GetLastEpoch(ctx sdk.Context) (uint32, error)
	SetLastEpoch(ctx sdk.Context, epoch uint32) error

	GetCommittee(ctx sdk.Context, epoch uint32) (Committee, error)
	GetLastCommittee(ctx sdk.Context) (Committee, error)
	SetCommittee(ctx sdk.Context, epoch uint32, committee Committee) error
	HasCommittee(ctx sdk.Context, epoch uint32) (bool, error)
	IterateCommittees(ctx sdk.Context, fn func(committee Committee) error) error
}

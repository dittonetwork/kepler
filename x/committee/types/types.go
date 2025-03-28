package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ExecutorI interface {
	GetAddress() sdk.AccAddress
	GetVotingPower() uint32
}

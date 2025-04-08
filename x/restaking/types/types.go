package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EmergencyValidator struct {
	Address     sdk.ValAddress
	VotingPower int64
}

package types

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"

	cosmosmath "cosmossdk.io/math"
	stakingtypes "cosmossdk.io/x/staking/types"
)

type DecreaseStakeProps struct {
	Amount          types.Coin
	OperatorAddress string
}

type IncreaseStakeProps struct {
	Amount          types.Coin
	OperatorAddress string
}

type RegisterOperatorProps struct {
	Commission        stakingtypes.CommissionRates
	Description       stakingtypes.Description
	MinSelfDelegation cosmosmath.Int
	OperatorAddress   string
	PubKey            cryptotypes.PubKey
	StakeValue        types.Coin
}

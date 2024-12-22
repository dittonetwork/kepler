package staking

import (
	"context"

	"kepler/x/staking/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WriteValidators returns a slice of bonded genesis validators.
func WriteValidators(ctx context.Context, keeper *keeper.Keeper) (vals []sdk.GenesisValidator, returnErr error) {
	panic("implement me")
}

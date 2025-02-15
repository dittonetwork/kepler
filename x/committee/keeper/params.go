package keeper

import (
	"context"

	"kepler/x/committee/types"
)

// GetParams get all parameters as types.Params.
func (k Keeper) GetParams(_ context.Context) types.Params {
	params := types.DefaultParams()

	return params
}

// SetParams set the params.
func (k Keeper) SetParams(_ context.Context, _ types.Params) error {
	return nil
}

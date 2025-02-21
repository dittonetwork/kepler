package keeper

import (
	"context"
	"kepler/x/committee/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetCommittee(
	_ context.Context,
	_ *types.QueryGetCommitteeRequest,
) (*types.QueryGetCommitteeResponse, error) {
	// TODO implement me
	panic("implement me")
}

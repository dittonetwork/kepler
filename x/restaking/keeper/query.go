package keeper

import (
	"context"

	"github.com/dittonetwork/kepler/x/restaking/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{Keeper: keeper}
}

func (q queryServer) PendingValidators(
	_ context.Context,
	_ *types.QueryPendingValidatorsRequest,
) (*types.QueryPendingValidatorsResponse, error) {
	// @TODO https://github.com/dittonetwork/kepler/issues/175
	panic("implement me")
}

// Validators returns the list of all validators.
func (q queryServer) Validators(
	_ context.Context,
	_ *types.QueryValidatorsRequest,
) (*types.QueryValidatorsResponse, error) {
	// TODO github.com/dittonetwork/kepler/issues/177
	panic("implement me")
}

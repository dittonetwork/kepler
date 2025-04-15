package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (q queryServer) NeedValidatorsUpdate(
	ctx context.Context,
	_ *types.QueryNeedValidatorsUpdateRequest) (*types.QueryNeedValidatorsUpdateResponse, error) {
	lastUpdate, err := q.repository.GetLastUpdate(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		return nil, err
	}

	epoch, err := q.epochs.GetEpochInfo(ctx, q.mainEpochID)
	if err != nil {
		return nil, err
	}

	return &types.QueryNeedValidatorsUpdateResponse{
		Result: lastUpdate.EpochNum <= epoch.CurrentEpoch,
	}, nil
}

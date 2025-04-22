package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{Keeper: keeper}
}

// PendingOperators returns the list of pending operators.
func (q queryServer) PendingOperators(
	ctx context.Context,
	_ *types.QueryPendingOperatorsRequest,
) (*types.QueryPendingOperatorsResponse, error) {
	operators, err := q.repository.GetPendingOperators(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		return nil, err
	}

	return &types.QueryPendingOperatorsResponse{
		Pending: operators,
	}, nil
}

// Validators returns the list of all validators.
func (q queryServer) Validators(
	ctx context.Context,
	req *types.QueryValidatorsRequest,
) (*types.QueryValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	validators, pagination, err := q.repository.GetPaginatedValidatorsByStatus(
		sdk.UnwrapSDKContext(ctx),
		req.Pagination,
		req.Status,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryValidatorsResponse{
		Validators: validators,
		Pagination: pagination,
	}, nil
}

func (q queryServer) NeedValidatorsUpdate(
	ctx context.Context,
	_ *types.QueryNeedValidatorsUpdateRequest,
) (*types.QueryNeedValidatorsUpdateResponse, error) {
	lastUpdate, err := q.repository.GetLastUpdate(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		return nil, err
	}

	epoch, err := q.epochs.GetEpochInfo(ctx, q.mainEpochID)
	if err != nil {
		return nil, err
	}

	return &types.QueryNeedValidatorsUpdateResponse{
		Result: lastUpdate.EpochNum < epoch.CurrentEpoch,
	}, nil
}

// OperatorStatus returns the status of an operator by its EVM address.
func (q queryServer) OperatorStatus(
	ctx context.Context,
	req *types.QueryOperatorStatusRequest,
) (*types.QueryOperatorStatusResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// check in pending pool
	operator, err := q.repository.GetPendingOperator(sdkCtx, req.EvmAddress)
	if err == nil {
		// operator found in pending pool
		return &types.QueryOperatorStatusResponse{
			Status: types.PendingOperatorStatus,
			Info:   operator,
		}, nil
	}

	if !errors.Is(err, collections.ErrNotFound) {
		return nil, err
	}

	// check in validators pool
	validator, err := q.repository.GetValidatorByEvmAddr(sdkCtx, req.EvmAddress)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrNotFoundValidator.Wrapf("operator %s not found", req.EvmAddress)
		}

		return nil, err
	}

	operator = *validator.ConvertToOperator()

	switch validator.Status {
	case types.Bonded:
		return &types.QueryOperatorStatusResponse{
			Status: types.ActiveOperatorStatus,
			Info:   operator,
		}, nil
	case types.Unbonding, types.Unbonded:
		return &types.QueryOperatorStatusResponse{
			Status: types.InactiveOperatorStatus,
			Info:   operator,
		}, nil

	case types.UnspecifiedStatus:
		return &types.QueryOperatorStatusResponse{
			Status: types.UnknownOperatorStatus,
			Info:   operator,
		}, nil
	}

	return &types.QueryOperatorStatusResponse{
		Status: types.UnknownOperatorStatus,
		Info:   operator,
	}, nil
}

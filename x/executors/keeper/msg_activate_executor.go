package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/executors/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) ActivateExecutor(
	ctx context.Context,
	msg *types.MsgActivateExecutor,
) (*types.MsgActivateExecutorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	executor, err := k.Executors.Get(sdkCtx, msg.GetCreator())
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "executor not found")
		}
	}

	ownersExecutors, err := k.GetExecutorsByOwnerAddress(sdkCtx, executor.OwnerAddress)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var hasActiveExecutors bool
	for _, e := range ownersExecutors {
		if e.IsActive {
			hasActiveExecutors = true
			break
		}
	}

	if hasActiveExecutors {
		return nil, status.Error(codes.AlreadyExists, "owner can have only one active executor")
	}

	executor.IsActive = true
	if err = k.Executors.Set(sdkCtx, executor.Address, executor); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.MsgActivateExecutorResponse{}, nil
}

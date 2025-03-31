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

func (k msgServer) DeactivateExecutor(
	ctx context.Context,
	msg *types.MsgDeactivateExecutor,
) (*types.MsgDeactivateExecutorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	executor, err := k.Executors.Get(sdkCtx, msg.GetCreator())
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "executor not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	executor.IsActive = false
	if err = k.Executors.Set(sdkCtx, executor.Address, executor); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.MsgDeactivateExecutorResponse{}, nil
}

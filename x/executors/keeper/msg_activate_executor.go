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

	if err := k.CheckAndSetExecutorIsActive(sdkCtx, msg.GetCreator()); err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "executor not found")
		}

		if errors.Is(err, types.ErrAlreadyHasActiveExecutors) {
			return nil, status.Error(codes.AlreadyExists, "owner can have only one active executor")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.MsgActivateExecutorResponse{}, nil
}

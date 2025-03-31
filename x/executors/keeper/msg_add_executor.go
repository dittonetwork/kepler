package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/executors/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) AddExecutor(
	ctx context.Context,
	msg *types.MsgAddExecutor,
) (*types.MsgAddExecutorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	newExecutor := types.Executor{
		Address:      msg.GetCreator(),
		OwnerAddress: msg.GetOwnerAddress(),
		PublicKey:    msg.GetPublicKey(),
	}

	insertedExecutor, err := k.addExecutor(sdkCtx, newExecutor)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.MsgAddExecutorResponse{
		Executor: insertedExecutor,
	}, nil
}

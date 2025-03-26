package keeper

import (
	"context"

	"github.com/dittonetwork/kepler/x/executors/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) AddExecutor(
	_ context.Context,
	_ *types.MsgAddExecutor,
) (*types.MsgAddExecutorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) ActivateExecutor(
	_ context.Context,
	_ *types.MsgActivateExecutor,
) (*types.MsgActivateExecutorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) DeactivateExecutor(
	_ context.Context,
	_ *types.MsgDeactivateExecutor,
) (*types.MsgDeactivateExecutorResponse, error) {
	//TODO implement me
	panic("implement me")
}

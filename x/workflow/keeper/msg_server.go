package keeper

import (
	"kepler/x/workflow/types"
)

type msgServer struct {
	Keeper
	CommitteeKeeper types.CommitteeKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, cmtKeeper types.CommitteeKeeper) types.MsgServer {
	return &msgServer{Keeper: keeper, CommitteeKeeper: cmtKeeper}
}

var _ types.MsgServer = msgServer{}

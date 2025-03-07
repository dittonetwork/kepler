package keeper

import (
	"github.com/dittonetwork/kepler/x/committee/types"
)

type msgServer struct {
	Keeper CommitteeKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper CommitteeKeeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

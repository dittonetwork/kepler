package keeper

import (
	"github.com/dittonetwork/kepler/x/workflow/types"
)

type msgServer struct {
	Keeper
	CommitteeKeeper types.CommitteeKeeper
	JobKeeper       types.JobKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(
	keeper Keeper,
	cmtKeeper types.CommitteeKeeper,
	jobKeeper types.JobKeeper,
) types.MsgServer {
	return &msgServer{
		Keeper:          keeper,
		CommitteeKeeper: cmtKeeper,
		JobKeeper:       jobKeeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k BaseKeeper) mustEmbedUnimplementedQueryServer() {}

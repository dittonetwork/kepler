package keeper

import (
	"context"

	"github.com/dittonetwork/kepler/x/restaking/types"
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

// BondValidator complete the bonding process for a validator recognized in Bonding status.
func (k msgServer) BondValidator(
	_ context.Context,
	_ *types.BondValidatorRequest,
) (*types.BondValidatorResponse, error) {
	// TODO github.com/dittonetwork/kepler/issues/175
	panic("implement me")
}

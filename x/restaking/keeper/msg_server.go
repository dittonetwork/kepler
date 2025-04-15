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
func (s msgServer) BondValidator(
	_ context.Context,
	_ *types.MsgBondValidator,
) (*types.MsgBondValidatorResponse, error) {
	// TODO github.com/dittonetwork/kepler/issues/175
	panic("implement me")
}

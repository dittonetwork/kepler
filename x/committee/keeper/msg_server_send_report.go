package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

func (k msgServer) SendReport(goCtx context.Context, _ *types.MsgSendReport) (*types.MsgSendReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendReportResponse{}, nil
}

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

func (k msgServer) SendReport(goCtx context.Context, msg *types.MsgSendReport) (*types.MsgSendReportResponse, error) {
	err := k.Keeper.HandleReport(sdk.UnwrapSDKContext(goCtx), msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendReportResponse{}, nil
}

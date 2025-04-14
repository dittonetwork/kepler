package keeper

import (
	"errors"
	"strconv"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

func (k Keeper) HandleReport(ctx sdk.Context, msg *types.MsgSendReport) error {
	lastCommittee, err := k.repository.GetLastCommittee(ctx)
	if err != nil {
		return err
	}

	if lastCommittee.GetAddress() != msg.GetCreator() {
		return sdkerrors.Wrapf(errors.New("invalid committee"), "invalid committee")
	}

	if msg.GetReport().GetCommitteeId() != lastCommittee.GetId() {
		return sdkerrors.Wrapf(errors.New("invalid committee id"), "invalid committee id")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventKeyReportGot,
			sdk.NewAttribute("committee_id", lastCommittee.GetId()),
			sdk.NewAttribute("creator", msg.GetCreator()),
			sdk.NewAttribute("context", msg.GetReport().GetExecutionContext().String()),
			sdk.NewAttribute("report_count", strconv.Itoa(len(msg.GetReport().GetMessages()))),
		),
	)

	// Route the message to the correct handler
	for _, msg := range msg.GetReport().GetMessages() {
		var sdkMsg sdk.Msg
		err = k.cdc.UnpackAny(msg, &sdkMsg)
		if err != nil {
			return sdkerrors.Wrapf(err, "cannot unpack message")
		}

		handler := k.router.Handler(sdkMsg)
		if handler == nil {
			return sdkerrors.Wrapf(errors.New("no handler for route"),
				"no handler for %s", sdkMsg.String())
		}

		_, err = handler(ctx, sdkMsg)
		if err != nil {
			return err
		}
	}
	return nil
}

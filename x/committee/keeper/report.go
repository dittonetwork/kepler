package keeper

import (
	"errors"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

func (k Keeper) HandleReport(ctx sdk.Context, msg *types.MsgSendReport) error {
	err := k.checkCommitteeValidity(ctx, msg.GetCreator(), msg.GetEpochId())
	if err != nil {
		return sdkerrors.Wrapf(err, "invalid committee")
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventReportReceived{
			Creator:     msg.Creator,
			EpochId:     msg.GetEpochId(),
			ReportCount: int64(len(msg.GetReport().GetMessages())),
		},
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

func (k Keeper) checkCommitteeValidity(ctx sdk.Context, address string, epochID uint32) error {
	lastCommittee, err := k.repository.GetLastCommittee(ctx)
	if err != nil {
		return err
	}

	if lastCommittee.GetAddress() != address {
		return k.epochIsValid(ctx, address, epochID, lastCommittee)
	}

	return nil
}

func (k Keeper) epochIsValid(ctx sdk.Context, address string, epochID uint32, lastCommittee types.Committee) error {
	epochCommittee, err := k.repository.GetCommittee(ctx, epochID)
	if err != nil {
		return err
	}

	if lastCommittee.GetEpoch()-epochID > 1 {
		return sdkerrors.Wrapf(errors.New("epoch is too old"), "epoch is too old")
	}

	if epochCommittee.GetAddress() != address {
		return sdkerrors.Wrapf(errors.New("invalid committee"), "invalid committee")
	}

	return nil
}

package keeper

import (
	"context"
	"fmt"

	"kepler/x/alliance/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAlliancesTimeline(goCtx context.Context, msg *types.MsgCreateAlliancesTimeline) (*types.MsgCreateAlliancesTimelineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var alliancesTimeline = types.AlliancesTimeline{
		Creator:      msg.Creator,
		Participants: msg.Participants,
		StartBlock:   msg.StartBlock,
		EndBlock:     msg.EndBlock,
	}

	id := k.AppendAlliancesTimeline(
		ctx,
		alliancesTimeline,
	)

	return &types.MsgCreateAlliancesTimelineResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAlliancesTimeline(goCtx context.Context, msg *types.MsgUpdateAlliancesTimeline) (*types.MsgUpdateAlliancesTimelineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var alliancesTimeline = types.AlliancesTimeline{
		Creator:      msg.Creator,
		Id:           msg.Id,
		Participants: msg.Participants,
		StartBlock:   msg.StartBlock,
		EndBlock:     msg.EndBlock,
	}

	// Checks that the element exists
	val, found := k.GetAlliancesTimeline(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAlliancesTimeline(ctx, alliancesTimeline)

	return &types.MsgUpdateAlliancesTimelineResponse{}, nil
}

func (k msgServer) DeleteAlliancesTimeline(goCtx context.Context, msg *types.MsgDeleteAlliancesTimeline) (*types.MsgDeleteAlliancesTimelineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAlliancesTimeline(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAlliancesTimeline(ctx, msg.Id)

	return &types.MsgDeleteAlliancesTimelineResponse{}, nil
}

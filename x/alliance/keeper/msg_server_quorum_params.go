package keeper

import (
	"context"

	"kepler/x/alliance/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateQuorumParams(goCtx context.Context, msg *types.MsgCreateQuorumParams) (*types.MsgCreateQuorumParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetQuorumParams(ctx)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var quorumParams = types.QuorumParams{
		Creator:          msg.Creator,
		MaxParticipants:  msg.MaxParticipants,
		ThresholdPercent: msg.ThresholdPercent,
		LifetimeInBlocks: msg.LifetimeInBlocks,
	}

	k.SetQuorumParams(
		ctx,
		quorumParams,
	)
	return &types.MsgCreateQuorumParamsResponse{}, nil
}

func (k msgServer) UpdateQuorumParams(goCtx context.Context, msg *types.MsgUpdateQuorumParams) (*types.MsgUpdateQuorumParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetQuorumParams(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var quorumParams = types.QuorumParams{
		Creator:          msg.Creator,
		MaxParticipants:  msg.MaxParticipants,
		ThresholdPercent: msg.ThresholdPercent,
		LifetimeInBlocks: msg.LifetimeInBlocks,
	}

	k.SetQuorumParams(ctx, quorumParams)

	return &types.MsgUpdateQuorumParamsResponse{}, nil
}

func (k msgServer) DeleteQuorumParams(goCtx context.Context, msg *types.MsgDeleteQuorumParams) (*types.MsgDeleteQuorumParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetQuorumParams(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveQuorumParams(ctx)

	return &types.MsgDeleteQuorumParamsResponse{}, nil
}

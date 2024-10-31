package keeper

import (
	"context"

	"kepler/x/beacon/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateFinalizedBlockInfo(goCtx context.Context, msg *types.MsgCreateFinalizedBlockInfo) (*types.MsgCreateFinalizedBlockInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetFinalizedBlockInfo(ctx)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var finalizedBlockInfo = types.FinalizedBlockInfo{
		Creator:        msg.Creator,
		SlotNum:        msg.SlotNum,
		BlockTimestamp: msg.BlockTimestamp,
		BlockNum:       msg.BlockNum,
		BlockHash:      msg.BlockHash,
	}

	k.SetFinalizedBlockInfo(
		ctx,
		finalizedBlockInfo,
	)
	return &types.MsgCreateFinalizedBlockInfoResponse{}, nil
}

func (k msgServer) UpdateFinalizedBlockInfo(goCtx context.Context, msg *types.MsgUpdateFinalizedBlockInfo) (*types.MsgUpdateFinalizedBlockInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFinalizedBlockInfo(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var finalizedBlockInfo = types.FinalizedBlockInfo{
		Creator:        msg.Creator,
		SlotNum:        msg.SlotNum,
		BlockTimestamp: msg.BlockTimestamp,
		BlockNum:       msg.BlockNum,
		BlockHash:      msg.BlockHash,
	}

	k.SetFinalizedBlockInfo(ctx, finalizedBlockInfo)

	return &types.MsgUpdateFinalizedBlockInfoResponse{}, nil
}

func (k msgServer) DeleteFinalizedBlockInfo(goCtx context.Context, msg *types.MsgDeleteFinalizedBlockInfo) (*types.MsgDeleteFinalizedBlockInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFinalizedBlockInfo(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveFinalizedBlockInfo(ctx)

	return &types.MsgDeleteFinalizedBlockInfoResponse{}, nil
}

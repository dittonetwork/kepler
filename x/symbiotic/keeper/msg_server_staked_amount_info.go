package keeper

import (
	"context"

	"kepler/x/symbiotic/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStakedAmountInfo(goCtx context.Context, msg *types.MsgCreateStakedAmountInfo) (*types.MsgCreateStakedAmountInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetStakedAmountInfo(
		ctx,
		msg.EthereumAddress,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var stakedAmountInfo = types.StakedAmountInfo{
		Creator:         msg.Creator,
		EthereumAddress: msg.EthereumAddress,
		StakedAmount:    msg.StakedAmount,
		LastUpdated:     msg.LastUpdated,
	}

	k.SetStakedAmountInfo(
		ctx,
		stakedAmountInfo,
	)
	return &types.MsgCreateStakedAmountInfoResponse{}, nil
}

func (k msgServer) UpdateStakedAmountInfo(goCtx context.Context, msg *types.MsgUpdateStakedAmountInfo) (*types.MsgUpdateStakedAmountInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStakedAmountInfo(
		ctx,
		msg.EthereumAddress,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var stakedAmountInfo = types.StakedAmountInfo{
		Creator:         msg.Creator,
		EthereumAddress: msg.EthereumAddress,
		StakedAmount:    msg.StakedAmount,
		LastUpdated:     msg.LastUpdated,
	}

	k.SetStakedAmountInfo(ctx, stakedAmountInfo)

	return &types.MsgUpdateStakedAmountInfoResponse{}, nil
}

func (k msgServer) DeleteStakedAmountInfo(goCtx context.Context, msg *types.MsgDeleteStakedAmountInfo) (*types.MsgDeleteStakedAmountInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStakedAmountInfo(
		ctx,
		msg.EthereumAddress,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveStakedAmountInfo(
		ctx,
		msg.EthereumAddress,
	)

	return &types.MsgDeleteStakedAmountInfoResponse{}, nil
}

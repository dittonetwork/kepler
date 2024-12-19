package keeper

import (
	"context"

	"kepler/x/alliance/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSharedEntropy(goCtx context.Context, msg *types.MsgCreateSharedEntropy) (*types.MsgCreateSharedEntropyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSharedEntropy(ctx)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var sharedEntropy = types.SharedEntropy{
		Creator: msg.Creator,
		Entropy: msg.Entropy,
	}

	k.SetSharedEntropy(
		ctx,
		sharedEntropy,
	)
	return &types.MsgCreateSharedEntropyResponse{}, nil
}

func (k msgServer) UpdateSharedEntropy(goCtx context.Context, msg *types.MsgUpdateSharedEntropy) (*types.MsgUpdateSharedEntropyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSharedEntropy(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var sharedEntropy = types.SharedEntropy{
		Creator: msg.Creator,
		Entropy: msg.Entropy,
	}

	k.SetSharedEntropy(ctx, sharedEntropy)

	return &types.MsgUpdateSharedEntropyResponse{}, nil
}

func (k msgServer) DeleteSharedEntropy(goCtx context.Context, msg *types.MsgDeleteSharedEntropy) (*types.MsgDeleteSharedEntropyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSharedEntropy(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSharedEntropy(ctx)

	return &types.MsgDeleteSharedEntropyResponse{}, nil
}

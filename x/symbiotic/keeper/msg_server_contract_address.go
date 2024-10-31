package keeper

import (
	"context"

	"kepler/x/symbiotic/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateContractAddress(goCtx context.Context, msg *types.MsgCreateContractAddress) (*types.MsgCreateContractAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetContractAddress(ctx)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var contractAddress = types.ContractAddress{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetContractAddress(
		ctx,
		contractAddress,
	)
	return &types.MsgCreateContractAddressResponse{}, nil
}

func (k msgServer) UpdateContractAddress(goCtx context.Context, msg *types.MsgUpdateContractAddress) (*types.MsgUpdateContractAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetContractAddress(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var contractAddress = types.ContractAddress{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetContractAddress(ctx, contractAddress)

	return &types.MsgUpdateContractAddressResponse{}, nil
}

func (k msgServer) DeleteContractAddress(goCtx context.Context, msg *types.MsgDeleteContractAddress) (*types.MsgDeleteContractAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetContractAddress(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveContractAddress(ctx)

	return &types.MsgDeleteContractAddressResponse{}, nil
}

package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateContractAddress{}

func NewMsgCreateContractAddress(creator string, address string) *MsgCreateContractAddress {
	return &MsgCreateContractAddress{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateContractAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateContractAddress{}

func NewMsgUpdateContractAddress(creator string, address string) *MsgUpdateContractAddress {
	return &MsgUpdateContractAddress{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateContractAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteContractAddress{}

func NewMsgDeleteContractAddress(creator string) *MsgDeleteContractAddress {
	return &MsgDeleteContractAddress{
		Creator: creator,
	}
}

func (msg *MsgDeleteContractAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

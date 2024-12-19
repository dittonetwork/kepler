package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSharedEntropy{}

func NewMsgCreateSharedEntropy(creator string, entropy uint64) *MsgCreateSharedEntropy {
	return &MsgCreateSharedEntropy{
		Creator: creator,
		Entropy: entropy,
	}
}

func (msg *MsgCreateSharedEntropy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSharedEntropy{}

func NewMsgUpdateSharedEntropy(creator string, entropy uint64) *MsgUpdateSharedEntropy {
	return &MsgUpdateSharedEntropy{
		Creator: creator,
		Entropy: entropy,
	}
}

func (msg *MsgUpdateSharedEntropy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSharedEntropy{}

func NewMsgDeleteSharedEntropy(creator string) *MsgDeleteSharedEntropy {
	return &MsgDeleteSharedEntropy{
		Creator: creator,
	}
}

func (msg *MsgDeleteSharedEntropy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

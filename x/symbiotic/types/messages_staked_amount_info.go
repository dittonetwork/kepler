package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStakedAmountInfo{}

func NewMsgCreateStakedAmountInfo(
	creator string,
	ethereumAddress string,
	stakedAmount string,
	lastUpdateTs uint64,

) *MsgCreateStakedAmountInfo {
	return &MsgCreateStakedAmountInfo{
		Creator:         creator,
		EthereumAddress: ethereumAddress,
		StakedAmount:    stakedAmount,
		LastUpdateTs:    lastUpdateTs,
	}
}

func (msg *MsgCreateStakedAmountInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateStakedAmountInfo{}

func NewMsgUpdateStakedAmountInfo(
	creator string,
	ethereumAddress string,
	stakedAmount string,
	lastUpdateTs uint64,

) *MsgUpdateStakedAmountInfo {
	return &MsgUpdateStakedAmountInfo{
		Creator:         creator,
		EthereumAddress: ethereumAddress,
		StakedAmount:    stakedAmount,
		LastUpdateTs:    lastUpdateTs,
	}
}

func (msg *MsgUpdateStakedAmountInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteStakedAmountInfo{}

func NewMsgDeleteStakedAmountInfo(
	creator string,
	ethereumAddress string,

) *MsgDeleteStakedAmountInfo {
	return &MsgDeleteStakedAmountInfo{
		Creator:         creator,
		EthereumAddress: ethereumAddress,
	}
}

func (msg *MsgDeleteStakedAmountInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFinalizedBlockInfo{}

func NewMsgCreateFinalizedBlockInfo(creator string, slotNum uint64, blockTimestamp uint64, blockNum uint64, blockHash string) *MsgCreateFinalizedBlockInfo {
	return &MsgCreateFinalizedBlockInfo{
		Creator:        creator,
		SlotNum:        slotNum,
		BlockTimestamp: blockTimestamp,
		BlockNum:       blockNum,
		BlockHash:      blockHash,
	}
}

func (msg *MsgCreateFinalizedBlockInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateFinalizedBlockInfo{}

func NewMsgUpdateFinalizedBlockInfo(creator string, slotNum uint64, blockTimestamp uint64, blockNum uint64, blockHash string) *MsgUpdateFinalizedBlockInfo {
	return &MsgUpdateFinalizedBlockInfo{
		Creator:        creator,
		SlotNum:        slotNum,
		BlockTimestamp: blockTimestamp,
		BlockNum:       blockNum,
		BlockHash:      blockHash,
	}
}

func (msg *MsgUpdateFinalizedBlockInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteFinalizedBlockInfo{}

func NewMsgDeleteFinalizedBlockInfo(creator string) *MsgDeleteFinalizedBlockInfo {
	return &MsgDeleteFinalizedBlockInfo{
		Creator: creator,
	}
}

func (msg *MsgDeleteFinalizedBlockInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

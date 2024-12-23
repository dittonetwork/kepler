package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateQuorumParams{}

func NewMsgCreateQuorumParams(creator string, maxParticipants uint64, thresholdPercent uint64, lifetimeInBlocks uint64) *MsgCreateQuorumParams {
	return &MsgCreateQuorumParams{
		Creator:          creator,
		MaxParticipants:  maxParticipants,
		ThresholdPercent: thresholdPercent,
		LifetimeInBlocks: lifetimeInBlocks,
	}
}

func (msg *MsgCreateQuorumParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateQuorumParams{}

func NewMsgUpdateQuorumParams(creator string, maxParticipants uint64, thresholdPercent uint64, lifetimeInBlocks uint64) *MsgUpdateQuorumParams {
	return &MsgUpdateQuorumParams{
		Creator:          creator,
		MaxParticipants:  maxParticipants,
		ThresholdPercent: thresholdPercent,
		LifetimeInBlocks: lifetimeInBlocks,
	}
}

func (msg *MsgUpdateQuorumParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteQuorumParams{}

func NewMsgDeleteQuorumParams(creator string) *MsgDeleteQuorumParams {
	return &MsgDeleteQuorumParams{
		Creator: creator,
	}
}

func (msg *MsgDeleteQuorumParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

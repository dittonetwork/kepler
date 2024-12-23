package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAlliancesTimeline{}

func NewMsgCreateAlliancesTimeline(creator string, participants []string, startBlock uint64, endBlock uint64) *MsgCreateAlliancesTimeline {
	return &MsgCreateAlliancesTimeline{
		Creator:      creator,
		Participants: participants,
		StartBlock:   startBlock,
		EndBlock:     endBlock,
	}
}

func (msg *MsgCreateAlliancesTimeline) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAlliancesTimeline{}

func NewMsgUpdateAlliancesTimeline(creator string, id uint64, participants []string, startBlock uint64, endBlock uint64) *MsgUpdateAlliancesTimeline {
	return &MsgUpdateAlliancesTimeline{
		Id:           id,
		Creator:      creator,
		Participants: participants,
		StartBlock:   startBlock,
		EndBlock:     endBlock,
	}
}

func (msg *MsgUpdateAlliancesTimeline) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAlliancesTimeline{}

func NewMsgDeleteAlliancesTimeline(creator string, id uint64) *MsgDeleteAlliancesTimeline {
	return &MsgDeleteAlliancesTimeline{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteAlliancesTimeline) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddEntropy{}

func NewMsgAddEntropy(contributor string, entropy uint64) *MsgAddEntropy {
	return &MsgAddEntropy{
		Contributor: contributor,
		Entropy:     entropy,
	}
}

func (msg *MsgAddEntropy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Contributor)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contributor address (%s)", err)
	}
	return nil
}

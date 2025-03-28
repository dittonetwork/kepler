package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddExecutor{}

func (msg *MsgAddExecutor) ValidateBasic() error {
	creator, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	owner, err := sdk.AccAddressFromBech32(msg.GetOwnerAddress())
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if creator.Equals(owner) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "creator and owner cannot be the same")
	}

	return nil
}

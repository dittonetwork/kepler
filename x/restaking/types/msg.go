package types

import (
	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgBondValidator{}
)

// Validate validates the MsgBondValidator sdk msg.
func (msg MsgBondValidator) Validate(ac address.Codec) error {
	_, err := ac.StringToBytes(msg.Owner)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid validator address: %s", err)
	}

	return nil
}

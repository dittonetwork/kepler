package types

import (
	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgBondValidator{}
)

// NewMsgBondValidator defines a message to bond a validator.
// It is used to create a new validator or update an existing one.
func NewMsgBondValidator(valAddr string, description Description) *MsgBondValidator {
	return &MsgBondValidator{
		Owner:       valAddr,
		Description: description,
	}
}

// Validate validates the MsgBondValidator sdk msg.
func (msg MsgBondValidator) Validate(ac address.Codec) error {
	_, err := ac.StringToBytes(msg.Owner)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid validator address: %s", err)
	}

	return nil
}

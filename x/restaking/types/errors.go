package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeInvalidSigner   = 1100
	ErrCodeValidatorUpdate = 1101
)

var (
	// ErrInvalidSigner defines an error when an account is not authorized to perform an action.
	ErrInvalidSigner = sdkerrors.Register(ModuleName, ErrCodeInvalidSigner, "invalid signer")

	// ErrUpdateValidator defines an error when a validator cannot be updated.
	ErrUpdateValidator = sdkerrors.Register(ModuleName, ErrCodeValidatorUpdate, "validator update failed")
)

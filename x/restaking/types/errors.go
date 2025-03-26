package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeInvalidSigner = 1100
	ErrCodeSample        = 1101
)

var (
	// ErrInvalidSigner defines an error when an account is not authorized to perform an action.
	ErrInvalidSigner = sdkerrors.Register(ModuleName, ErrCodeInvalidSigner, "invalid signer")

	// ErrSample defines a sample error.
	ErrSample = sdkerrors.Register(ModuleName, ErrCodeSample, "sample error")
)

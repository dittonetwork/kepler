package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeInvalidSigner = 1100
	ErrCodeSample        = 1101
)

// x/executors module sentinel errors.
var (
	ErrInvalidSigner = sdkerrors.Register(
		ModuleName,
		ErrCodeInvalidSigner,
		"expected gov account as only signer for proposal message",
	)
	ErrSample = sdkerrors.Register(ModuleName, ErrCodeSample, "sample error")
)

package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeInvalidSigner = 1100
	ErrCodeSample        = 1101
	ErrCodeAlreadyExists = 1102
	ErrCodeNotFound      = 1103
)

// x/workflow module sentinel errors.
var (
	ErrInvalidSigner = sdkerrors.Register(
		ModuleName,
		ErrCodeInvalidSigner,
		"expected gov account as only signer for proposal message",
	)
	ErrSample                  = sdkerrors.Register(ModuleName, ErrCodeSample, "sample error")
	ErrAutomationAlreadyExists = sdkerrors.Register(
		ModuleName,
		ErrCodeAlreadyExists,
		"automation already exists",
	)
	ErrAutomationNotFound = sdkerrors.Register(ModuleName, ErrCodeNotFound, "automation not found")
)

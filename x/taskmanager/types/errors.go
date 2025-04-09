package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeInvalidSigner   = 1100
	ErrCodeValidatorUpdate = 1101
)

// x/taskmanager module sentinel errors.
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, ErrCodeInvalidSigner, "invalid signer")
	ErrSample        = sdkerrors.Register(ModuleName, ErrCodeValidatorUpdate, "sample error")
)

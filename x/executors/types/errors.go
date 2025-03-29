package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeSample = 1101
)

// x/executors module sentinel errors.
var (
	ErrSample = sdkerrors.Register(ModuleName, ErrCodeSample, "sample error")
)

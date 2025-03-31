package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeSample                    = 1101
	ErrCodeAlreadyHasActiveExecutors = 1102
)

// x/executors module sentinel errors.
var (
	ErrSample                    = sdkerrors.Register(ModuleName, ErrCodeSample, "sample error")
	ErrAlreadyHasActiveExecutors = sdkerrors.Register(
		ModuleName,
		ErrCodeAlreadyHasActiveExecutors,
		"owner already has active executors",
	)
)

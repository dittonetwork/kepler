package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeAlreadyHasActiveExecutor = 1102
)

var ErrAlreadyHasActiveExecutor = sdkerrors.Register(
	ModuleName,
	ErrCodeAlreadyHasActiveExecutor,
	"owner already has active executor",
)

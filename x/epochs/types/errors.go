package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrInvalidSignerCode = 1100
	ErrSampleCode        = 1101
)

// x/epochs module sentinel errors.
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName,
		ErrInvalidSignerCode,
		"expected gov account as only signer for proposal message",
	)
	ErrSample = sdkerrors.Register(ModuleName,
		ErrSampleCode,
		"sample error",
	)
)

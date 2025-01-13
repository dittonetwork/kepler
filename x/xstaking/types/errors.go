package types

// DONTCOVER

import (
	"cosmossdk.io/errors/v2"
)

var (
	ErrInvalidSigner = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
)

package types

// DONTCOVER

import (
	"cosmossdk.io/errors/v2"
)

// x/xstaking module sentinel errors
var (
	ErrInvalidSigner      = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrEmptyValidatorAddr = errors.Register(ModuleName, 2, "empty validator address")
	ErrNoValidatorFound   = errors.Register(ModuleName, 3, "validator does not exist")
)

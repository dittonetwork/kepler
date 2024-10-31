package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/beacon module sentinel errors
var (
	ErrInvalidSigner    = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample           = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrBeaconNotFound   = sdkerrors.Register(ModuleName, 1102, "no such block in beacon chain")
	ErrNoFinalizedBlock = sdkerrors.Register(ModuleName, 1103, "found block is not finalized")
)

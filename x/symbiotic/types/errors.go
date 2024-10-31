package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/symbiotic module sentinel errors
var (
	ErrInvalidSigner                  = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample                         = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrContractNotInStorage           = sdkerrors.Register(ModuleName, 1102, "no contract address in storage")
	ErrNoFinalizedBlockInBeaconModule = sdkerrors.Register(ModuleName, 1103, "no finalized block stored in beacon module keeper")
)

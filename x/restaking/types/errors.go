package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeInvalidSigner     = 1100
	ErrCodeValidatorUpdate   = 1101
	ErrCodeGenesisInit       = 1102
	ErrCodeGenesisExport     = 1103
	ErrCodeNoPendingOperator = 1104
	ErrCodeNotFoundValidator = 1105
)

var (
	// ErrInvalidSigner defines an error when an account is not authorized to perform an action.
	ErrInvalidSigner = sdkerrors.Register(ModuleName, ErrCodeInvalidSigner, "invalid signer")

	// ErrUpdateValidator defines an error when a validator cannot be updated.
	ErrUpdateValidator = sdkerrors.Register(ModuleName, ErrCodeValidatorUpdate, "validator update failed")

	// ErrGenesisInit defines an error when the genesis state cannot be initialized.
	ErrGenesisInit = sdkerrors.Register(ModuleName, ErrCodeGenesisInit, "genesis init failed")

	// ErrGenesisExport defines an error when the genesis state cannot be exported.
	ErrGenesisExport = sdkerrors.Register(ModuleName, ErrCodeGenesisExport, "genesis export failed")

	// ErrNoPendingOperator defines an error when the no pending operator is found.
	ErrNoPendingOperator = sdkerrors.Register(ModuleName, ErrCodeNoPendingOperator, "no pending operator")

	// ErrNotFoundValidator defines an error when the validator is not found.
	ErrNotFoundValidator = sdkerrors.Register(ModuleName, ErrCodeNotFoundValidator, "validator not found")
)

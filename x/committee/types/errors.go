package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeInvalidSigner          = 1100
	ErrCodeCommitteeAlreadyExists = 1102
	ErrCodeCommitteeInvalidEpoch  = 1103
)

// x/committee module sentinel errors.
var (
	ErrInvalidSigner = sdkerrors.Register(
		ModuleName,
		ErrCodeInvalidSigner,
		"expected gov account as only signer for proposal message",
	)
	ErrCommitteeAlreadyExists = sdkerrors.Register(
		ModuleName,
		ErrCodeCommitteeAlreadyExists,
		"committee already exists",
	)

	ErrInvalidEpoch = sdkerrors.Register(
		ModuleName,
		ErrCodeCommitteeInvalidEpoch,
		"invalid epoch",
	)
)

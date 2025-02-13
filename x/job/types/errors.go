package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

const (
	ErrCodeJobSignsIsNil        = 1200
	ErrCodeJobAlreadyExists     = 1201
	ErrCodeCommitteeDoesntExist = 1202
)

// x/job module sentinel errors.
var (
	ErrJobSignsIsNil         = sdkerrors.Register(ModuleName, ErrCodeJobSignsIsNil, "job signs is nil")
	ErrJobAlreadyExists      = sdkerrors.Register(ModuleName, ErrCodeJobAlreadyExists, "job already exists")
	ErrCommitteeDoesntExists = sdkerrors.Register(ModuleName, ErrCodeCommitteeDoesntExist, "committee does not exist")
)

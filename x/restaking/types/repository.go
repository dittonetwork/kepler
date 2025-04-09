package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Repository interface {
	GetLastUpdate(ctx sdk.Context) (UpdateInfo, error)
	SetLastUpdate(ctx sdk.Context, info UpdateInfo) error

	GetValidator(ctx sdk.Context, operatorAddr string) (Validator, error)
	SetValidator(ctx sdk.Context, operatorAddr string, validator Validator) error
	GetAllValidators(ctx sdk.Context) ([]Validator, error)
	GetBondedValidators(ctx sdk.Context) ([]Validator, error)
	GetEmergencyValidators(ctx sdk.Context) ([]Validator, error)
	RemoveValidator(ctx sdk.Context, operatorAddr string) error

	GetPendingValidators(ctx sdk.Context) ([]Validator, error)
	SetPendingValidator(ctx sdk.Context, operatorAddr string, validator Validator) error
	RemovePendingValidator(ctx sdk.Context, operatorAddr string) error
}

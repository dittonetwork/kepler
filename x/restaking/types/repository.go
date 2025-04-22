package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ValidatorChangeType int

const (
	ValidatorChangeTypeUnknown ValidatorChangeType = iota
	ValidatorChangeTypeCreate
	ValidatorChangeTypeUpdate
	ValidatorChangeTypeDelete
)

type Repository interface {
	GetLastUpdate(ctx sdk.Context) (UpdateInfo, error)
	SetLastUpdate(ctx sdk.Context, info UpdateInfo) error

	GetPendingOperator(ctx sdk.Context, addr string) (Operator, error)
	GetPendingOperators(ctx sdk.Context) ([]Operator, error)
	SetPendingOperator(ctx sdk.Context, operatorAddr string, operator Operator) error
	RemovePendingOperator(ctx sdk.Context, operatorAddr string) error

	SetValidator(ctx sdk.Context, addr sdk.ValAddress, validator Validator) error
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (Validator, error)
	GetAllValidators(ctx sdk.Context) ([]Validator, error)
	GetBondedValidators(ctx sdk.Context) ([]Validator, error)
	GetEmergencyValidators(ctx sdk.Context) ([]Validator, error)
	GetValidatorByEvmAddr(ctx sdk.Context, addr string) (Validator, error)
	RemoveValidator(ctx sdk.Context, addr string) error

	AddValidatorsChange(ctx sdk.Context, validator Validator, ctype ValidatorChangeType) error
	GetValidatorsChanges(ctx sdk.Context) (ValidatorsChanges, error)
	PruneValidatorsChanges(ctx sdk.Context) error
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	RemoveValidatorByOperatorAddr(ctx sdk.Context, addr string) error
}

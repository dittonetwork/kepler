package repository

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// GetPendingOperators returns all pending operators.
func (s *RestakingRepository) GetPendingOperators(_ sdk.Context) ([]types.Operator, error) {
	// @TODO https://github.com/dittonetwork/kepler/issues/175
	panic("not implemented")
}

// SetPendingOperator sets the pending operator in the repository.
func (s *RestakingRepository) SetPendingOperator(
	ctx sdk.Context,
	operatorAddr string,
	validator types.Operator,
) error {
	return s.pending.Set(ctx, operatorAddr, validator)
}

// RemovePendingOperator removes the pending operator from the repository.
func (s *RestakingRepository) RemovePendingOperator(ctx sdk.Context, operatorAddr string) error {
	return s.pending.Remove(ctx, operatorAddr)
}

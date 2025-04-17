package repository

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// GetPendingOperators returns all pending operators.
func (s *RestakingRepository) GetPendingOperators(ctx sdk.Context) ([]types.Operator, error) {
	iter, err := s.pending.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	return iter.Values()
}

// SetPendingOperator sets the pending operator in the repository.
func (s *RestakingRepository) SetPendingOperator(
	ctx sdk.Context,
	addr string,
	validator types.Operator,
) error {
	return s.pending.Set(ctx, addr, validator)
}

// RemovePendingOperator removes the pending operator from the repository.
func (s *RestakingRepository) RemovePendingOperator(ctx sdk.Context, operatorAddr string) error {
	return s.pending.Remove(ctx, operatorAddr)
}

// GetPendingOperator retrieves a pending operator from the repository.
func (s *RestakingRepository) GetPendingOperator(ctx sdk.Context, addr string) (types.Operator, error) {
	return s.pending.Get(ctx, addr)
}

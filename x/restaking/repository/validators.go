package repository

import (
	"cosmossdk.io/collections/indexes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// GetPendingValidators returns all pending validators.
func (s *RestakingRepository) GetPendingValidators(_ sdk.Context) ([]types.Validator, error) {
	panic("not implemented")
}

// SetPendingValidator sets the pending validator in the repository.
func (s *RestakingRepository) SetPendingValidator(
	ctx sdk.Context,
	operatorAddr string,
	validator types.Validator,
) error {
	return s.pendingValidators.Set(ctx, operatorAddr, validator)
}

// RemovePendingValidator removes the pending validator from the repository.
func (s *RestakingRepository) RemovePendingValidator(ctx sdk.Context, operatorAddr string) error {
	return s.pendingValidators.Remove(ctx, operatorAddr)
}

// SetValidator sets the validator in the repository.
func (s *RestakingRepository) SetValidator(ctx sdk.Context, operatorAddr string, validator types.Validator) error {
	return s.validators.Set(ctx, operatorAddr, validator)
}

// GetAllValidators returns all validators.
func (s *RestakingRepository) GetAllValidators(ctx sdk.Context) ([]types.Validator, error) {
	iter, err := s.validators.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	return iter.Values()
}

// GetBondedValidators returns all bonded validators.
func (s *RestakingRepository) GetBondedValidators(ctx sdk.Context) ([]types.Validator, error) {
	iter, err := s.validators.Indexes.Bonded.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	return indexes.CollectValues(ctx, s.validators, iter)
}

// GetEmergencyValidators returns all active emergency validators.
func (s *RestakingRepository) GetEmergencyValidators(ctx sdk.Context) ([]types.Validator, error) {
	iter, err := s.validators.Indexes.Emergency.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	return indexes.CollectValues(ctx, s.validators, iter)
}

// RemoveValidator removes the validator from the repository.
func (s *RestakingRepository) RemoveValidator(ctx sdk.Context, operatorAddr string) error {
	return s.validators.Remove(ctx, operatorAddr)
}

package repository

import (
	"cosmossdk.io/collections/indexes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

func (s *RestakingRepository) GetValidatorByEvmAddr(ctx sdk.Context, addr string) (types.Validator, error) {
	valAddr, err := s.validators.Indexes.EvmAddress.MatchExact(ctx, addr)
	if err != nil {
		return types.Validator{}, err
	}

	return s.validators.Get(ctx, valAddr)
}

// SetValidator sets the validator in the repository.
func (s *RestakingRepository) SetValidator(ctx sdk.Context, addr sdk.ValAddress, validator types.Validator) error {
	return s.validators.Set(ctx, addr.String(), validator)
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

// RemoveValidatorByOperatorAddr removes the validator from the repository.
func (s *RestakingRepository) RemoveValidatorByOperatorAddr(ctx sdk.Context, addr string) error {
	valAddr, err := s.validators.Indexes.EvmAddress.MatchExact(ctx, addr)
	if err != nil {
		return err
	}

	return s.validators.Remove(ctx, valAddr)
}

// GetValidator returns validator by operator address.
func (s *RestakingRepository) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (types.Validator, error) {
	return s.validators.Get(ctx, addr.String())
}

package types

import (
	context "context"

	stakingtypes "cosmossdk.io/x/staking/types"
)

type EpochsKeeper interface {
	// TODO Add methods imported from epochs should be defined here
}

type StakingKeeper interface {
	// Add methods imported from staking should be defined here
	GetBondedValidators(ctx context.Context) (validators []stakingtypes.Validator, err error)
	UpdateValidatorsStakes(ctx context.Context, validatorsStakes map[string]uint64)
}

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// GetValidator returns the validator information for a given address.
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (types.Validator, error) {
	validator, err := k.staking.GetValidator(ctx, addr)
	if err != nil {
		return types.Validator{}, err
	}

	restakingValidator, err := k.ValidatorsMap.Get(ctx, addr.String())
	if err != nil {
		return types.Validator{}, err
	}

	return types.Validator{
		Address:       validator.GetOperator(),
		CosmosAddress: addr.String(),
		IsEmergency:   restakingValidator.GetIsEmergency(),
		VotingPower:   restakingValidator.GetVotingPower(),
		Status:        restakingValidator.GetStatus(),
	}, nil
}

package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// InitGenesis initializes the genesis state of the restaking module.
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) error {
	if err := k.repository.SetLastUpdate(ctx, state.LastUpdate); err != nil {
		return sdkerrors.Wrap(err, "last update")
	}

	for _, validator := range state.Validators {
		if err := k.repository.SetValidator(ctx, validator.OperatorAddress, validator); err != nil {
			return sdkerrors.Wrap(err, "validators")
		}
	}

	return nil
}

func (k Keeper) ExportGenesis(ctx sdk.Context) (*types.GenesisState, error) {
	genesis := types.DefaultGenesis()

	var err error
	genesis.LastUpdate, err = k.repository.GetLastUpdate(ctx)
	if err != nil {
		return genesis, sdkerrors.Wrap(err, "last update")
	}

	genesis.Validators, err = k.repository.GetAllValidators(ctx)
	if err != nil {
		return genesis, err
	}

	return genesis, nil
}

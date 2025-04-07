package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// InitGenesis initializes the genesis state of the restaking module.
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) error {
	if err := k.lastUpdate.Set(ctx, state.LastUpdate); err != nil {
		return sdkerrors.Wrap(err, "last update")
	}

	for _, validator := range state.Validators {
		if err := k.validators.Set(ctx, validator.OperatorAddress, validator); err != nil {
			return sdkerrors.Wrap(err, "validators")
		}
	}

	return nil
}

func (k Keeper) ExportGenesis(ctx sdk.Context) (*types.GenesisState, error) {
	genesis := types.DefaultGenesis()

	var err error
	genesis.LastUpdate, err = k.lastUpdate.Get(ctx)
	if err != nil {
		return genesis, sdkerrors.Wrap(err, "last update")
	}

	iter, err := k.validators.Iterate(ctx, nil)

	if err != nil {
		return genesis, sdkerrors.Wrap(err, "validators iterate")
	}

	defer iter.Close()

	genesis.Validators, err = iter.Values()
	if err != nil {
		return genesis, sdkerrors.Wrap(err, "validators values")
	}

	return genesis, nil
}

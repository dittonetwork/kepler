package restaking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dittonetwork/kepler/x/restaking/keeper"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	if err := k.LastUpdate.Set(ctx, genState.LastUpdate); err != nil {
		panic(err)
	}

	for _, validator := range genState.Validators {
		if err := k.ValidatorsMap.Set(ctx, validator.Address, validator); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	lastUpdate, err := k.LastUpdate.Get(ctx)
	if err != nil {
		panic(err)
	}

	genesis.LastUpdate = lastUpdate

	// Iterate through all validators and add them to genesis state
	validators := make([]types.Validator, 0)
	iter, err := k.ValidatorsMap.Iterate(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var validator types.Validator
		validator, err = iter.Value()
		if err != nil {
			panic(err)
		}
		validators = append(validators, validator)
	}

	genesis.Validators = validators

	return genesis
}

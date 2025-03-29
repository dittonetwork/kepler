package executors

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dittonetwork/kepler/x/executors/keeper"
	"github.com/dittonetwork/kepler/x/executors/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.Executors {
		err := k.Executors.Set(ctx, elem.Address, elem)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	var executors []types.Executor
	iter, err := k.Executors.Iterate(ctx, nil)
	if err != nil {
		panic(err)
	}

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var v types.Executor
		v, err = iter.Value()
		if err != nil {
			panic(err)
		}

		executors = append(executors, v)
	}

	genesis.Executors = executors

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

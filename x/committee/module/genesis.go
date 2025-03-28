package committee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dittonetwork/kepler/x/committee/keeper"
	"github.com/dittonetwork/kepler/x/committee/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// Set the last epoch
	if err := k.LastEpoch.Set(ctx, genState.LastEpoch); err != nil {
		panic(err)
	}

	for _, committee := range genState.Committees {
		if err := k.Committees.Set(ctx, committee.Epoch, committee); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	lastEpoch, err := k.LastEpoch.Get(ctx)
	if err != nil {
		panic(err)
	}

	genesis.LastEpoch = lastEpoch

	committees := make([]types.Committee, 0)
	iter, err := k.Committees.Iterate(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var committee types.Committee
		committee, err = iter.Value()
		if err != nil {
			panic(err)
		}

		committees = append(committees, committee)
	}

	genesis.Committees = committees

	return genesis
}

package job

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kepler/x/job/keeper"
	"kepler/x/job/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(_ sdk.Context, _ keeper.Keeper, _ types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(_ sdk.Context, _ keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

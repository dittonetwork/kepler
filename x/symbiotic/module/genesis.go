package symbiotic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kepler/x/symbiotic/keeper"
	"kepler/x/symbiotic/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the stakedAmountInfo
	for _, elem := range genState.StakedAmountInfoList {
		k.SetStakedAmountInfo(ctx, elem)
	}
	// Set if defined
	if genState.ContractAddress != nil {
		k.SetContractAddress(ctx, *genState.ContractAddress)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.StakedAmountInfoList = k.GetAllStakedAmountInfo(ctx)
	// Get all contractAddress
	contractAddress, found := k.GetContractAddress(ctx)
	if found {
		genesis.ContractAddress = &contractAddress
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

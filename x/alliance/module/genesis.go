package alliance

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kepler/x/alliance/keeper"
	"kepler/x/alliance/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.SharedEntropy != nil {
		k.SetSharedEntropy(ctx, *genState.SharedEntropy)
	}
	// Set if defined
	if genState.QuorumParams != nil {
		k.SetQuorumParams(ctx, *genState.QuorumParams)
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

	// Get all sharedEntropy
	sharedEntropy, found := k.GetSharedEntropy(ctx)
	if found {
		genesis.SharedEntropy = &sharedEntropy
	}
	// Get all quorumParams
	quorumParams, found := k.GetQuorumParams(ctx)
	if found {
		genesis.QuorumParams = &quorumParams
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

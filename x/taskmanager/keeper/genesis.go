package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dittonetwork/kepler/x/taskmanager/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(_ sdk.Context, _ types.GenesisState) error {
	return nil
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(_ sdk.Context) (*types.GenesisState, error) {
	genesis := types.DefaultGenesis()

	return genesis, nil
}

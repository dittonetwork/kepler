package keeper

import (
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dittonetwork/kepler/x/committee/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// Set the last epoch
	if err := k.repository.SetLastEpoch(ctx, genState.LastEpoch); err != nil {
		panic(err)
	}

	for _, committee := range genState.Committees {
		if err := k.repository.SetCommittee(ctx, committee.Epoch, committee); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	lastEpoch, err := k.repository.GetLastEpoch(ctx)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		panic(err)
	}

	genesis.LastEpoch = lastEpoch

	committees := make([]types.Committee, 0)
	err = k.repository.IterateCommittees(ctx, func(committee types.Committee) error {
		committees = append(committees, committee)
		return nil
	})
	if err != nil {
		panic(err)
	}

	genesis.Committees = committees

	return genesis
}

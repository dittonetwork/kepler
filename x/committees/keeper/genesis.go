package keeper

import (
	"context"

	"kepler/x/committees/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.CommitteesList {
		if err := k.Committees.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.CommitteesSeq.Set(ctx, genState.CommitteesCount); err != nil {
		return err

		// ExportGenesis returns the module's exported genesis.
	}
	return k.Params.Set(ctx, genState.Params)
}

func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = k.Committees.Walk(ctx, nil, func(key uint64, elem types.Committees) (bool, error) {
		genesis.CommitteesList = append(genesis.CommitteesList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.CommitteesCount, err = k.CommitteesSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}

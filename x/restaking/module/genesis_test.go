package restaking_test

import (
	"testing"

	keepertest "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/testutil/nullify"
	restaking "github.com/dittonetwork/kepler/x/restaking/module"
	"github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RestakingKeeper(t)
	restaking.InitGenesis(ctx, k, genesisState)
	got := restaking.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

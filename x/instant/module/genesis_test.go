package instant_test

import (
	"testing"

	keepertest "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/testutil/nullify"
	instant "github.com/dittonetwork/kepler/x/instant/module"
	"github.com/dittonetwork/kepler/x/instant/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InstantKeeper(t)
	instant.InitGenesis(ctx, k, genesisState)
	got := instant.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

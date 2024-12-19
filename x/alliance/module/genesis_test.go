package alliance_test

import (
	"testing"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	alliance "kepler/x/alliance/module"
	"kepler/x/alliance/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SharedEntropy: &types.SharedEntropy{
			Entropy: 77,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AllianceKeeper(t)
	alliance.InitGenesis(ctx, k, genesisState)
	got := alliance.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SharedEntropy, got.SharedEntropy)
	// this line is used by starport scaffolding # genesis/test/assert
}

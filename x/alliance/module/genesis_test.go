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
		QuorumParams: &types.QuorumParams{
			MaxParticipants:  60,
			ThresholdPercent: 90,
			LifetimeInBlocks: 87,
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
	require.Equal(t, genesisState.QuorumParams, got.QuorumParams)
	// this line is used by starport scaffolding # genesis/test/assert
}

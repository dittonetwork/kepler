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
		AlliancesTimelineList: []types.AlliancesTimeline{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AlliancesTimelineCount: 2,
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
	require.ElementsMatch(t, genesisState.AlliancesTimelineList, got.AlliancesTimelineList)
	require.Equal(t, genesisState.AlliancesTimelineCount, got.AlliancesTimelineCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

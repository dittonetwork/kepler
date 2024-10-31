package beacon_test

import (
	"testing"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	beacon "kepler/x/beacon/module"
	"kepler/x/beacon/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FinalizedBlockInfo: &types.FinalizedBlockInfo{
			SlotNum:        42,
			BlockTimestamp: 81,
			BlockNum:       54,
			BlockHash:      "84",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BeaconKeeper(t)
	beacon.InitGenesis(ctx, k, genesisState)
	got := beacon.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.FinalizedBlockInfo, got.FinalizedBlockInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}

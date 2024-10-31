package symbiotic_test

import (
	"testing"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	symbiotic "kepler/x/symbiotic/module"
	"kepler/x/symbiotic/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		StakedAmountInfoList: []types.StakedAmountInfo{
			{
				EthereumAddress: "0",
			},
			{
				EthereumAddress: "1",
			},
		},
		ContractAddress: &types.ContractAddress{
			Address: "73",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SymbioticKeeper(t)
	symbiotic.InitGenesis(ctx, k, genesisState)
	got := symbiotic.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StakedAmountInfoList, got.StakedAmountInfoList)
	require.Equal(t, genesisState.ContractAddress, got.ContractAddress)
	// this line is used by starport scaffolding # genesis/test/assert
}

package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	"kepler/x/symbiotic/keeper"
	"kepler/x/symbiotic/types"
)

func createTestContractAddress(keeper keeper.Keeper, ctx context.Context) types.ContractAddress {
	item := types.ContractAddress{}
	keeper.SetContractAddress(ctx, item)
	return item
}

func TestContractAddressGet(t *testing.T) {
	keeper, ctx := keepertest.SymbioticKeeper(t)
	item := createTestContractAddress(keeper, ctx)
	rst, found := keeper.GetContractAddress(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestContractAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.SymbioticKeeper(t)
	createTestContractAddress(keeper, ctx)
	keeper.RemoveContractAddress(ctx)
	_, found := keeper.GetContractAddress(ctx)
	require.False(t, found)
}

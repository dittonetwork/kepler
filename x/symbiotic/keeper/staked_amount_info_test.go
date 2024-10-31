package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	"kepler/x/symbiotic/keeper"
	"kepler/x/symbiotic/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStakedAmountInfo(keeper keeper.Keeper, ctx context.Context, n int) []types.StakedAmountInfo {
	items := make([]types.StakedAmountInfo, n)
	for i := range items {
		items[i].EthereumAddress = strconv.Itoa(i)

		keeper.SetStakedAmountInfo(ctx, items[i])
	}
	return items
}

func TestStakedAmountInfoGet(t *testing.T) {
	keeper, ctx := keepertest.SymbioticKeeper(t)
	items := createNStakedAmountInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStakedAmountInfo(ctx,
			item.EthereumAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStakedAmountInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.SymbioticKeeper(t)
	items := createNStakedAmountInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStakedAmountInfo(ctx,
			item.EthereumAddress,
		)
		_, found := keeper.GetStakedAmountInfo(ctx,
			item.EthereumAddress,
		)
		require.False(t, found)
	}
}

func TestStakedAmountInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.SymbioticKeeper(t)
	items := createNStakedAmountInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStakedAmountInfo(ctx)),
	)
}

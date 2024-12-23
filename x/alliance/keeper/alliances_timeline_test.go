package keeper_test

import (
	"context"
	"testing"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	"kepler/x/alliance/keeper"
	"kepler/x/alliance/types"

	"github.com/stretchr/testify/require"
)

func createNAlliancesTimeline(keeper keeper.Keeper, ctx context.Context, n int) []types.AlliancesTimeline {
	items := make([]types.AlliancesTimeline, n)
	for i := range items {
		items[i].Id = keeper.AppendAlliancesTimeline(ctx, items[i])
	}
	return items
}

func TestAlliancesTimelineGet(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	items := createNAlliancesTimeline(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAlliancesTimeline(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAlliancesTimelineRemove(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	items := createNAlliancesTimeline(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAlliancesTimeline(ctx, item.Id)
		_, found := keeper.GetAlliancesTimeline(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAlliancesTimelineGetAll(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	items := createNAlliancesTimeline(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAlliancesTimeline(ctx)),
	)
}

func TestAlliancesTimelineCount(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	items := createNAlliancesTimeline(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAlliancesTimelineCount(ctx))
}

package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	"kepler/x/beacon/keeper"
	"kepler/x/beacon/types"
)

func createTestFinalizedBlockInfo(keeper keeper.Keeper, ctx context.Context) types.FinalizedBlockInfo {
	item := types.FinalizedBlockInfo{}
	keeper.SetFinalizedBlockInfo(ctx, item)
	return item
}

func TestFinalizedBlockInfoGet(t *testing.T) {
	keeper, ctx := keepertest.BeaconKeeper(t)
	item := createTestFinalizedBlockInfo(keeper, ctx)
	rst, found := keeper.GetFinalizedBlockInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestFinalizedBlockInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.BeaconKeeper(t)
	createTestFinalizedBlockInfo(keeper, ctx)
	keeper.RemoveFinalizedBlockInfo(ctx)
	_, found := keeper.GetFinalizedBlockInfo(ctx)
	require.False(t, found)
}

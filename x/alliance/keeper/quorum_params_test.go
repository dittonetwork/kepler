package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	"kepler/x/alliance/keeper"
	"kepler/x/alliance/types"
)

func createTestQuorumParams(keeper keeper.Keeper, ctx context.Context) types.QuorumParams {
	item := types.QuorumParams{}
	keeper.SetQuorumParams(ctx, item)
	return item
}

func TestQuorumParamsGet(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	item := createTestQuorumParams(keeper, ctx)
	rst, found := keeper.GetQuorumParams(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestQuorumParamsRemove(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	createTestQuorumParams(keeper, ctx)
	keeper.RemoveQuorumParams(ctx)
	_, found := keeper.GetQuorumParams(ctx)
	require.False(t, found)
}

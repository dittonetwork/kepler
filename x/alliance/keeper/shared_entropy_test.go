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

func createTestSharedEntropy(keeper keeper.Keeper, ctx context.Context) types.SharedEntropy {
	item := types.SharedEntropy{}
	keeper.SetSharedEntropy(ctx, item)
	return item
}

func TestSharedEntropyGet(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	item := createTestSharedEntropy(keeper, ctx)
	rst, found := keeper.GetSharedEntropy(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestSharedEntropyRemove(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	createTestSharedEntropy(keeper, ctx)
	keeper.RemoveSharedEntropy(ctx)
	_, found := keeper.GetSharedEntropy(ctx)
	require.False(t, found)
}

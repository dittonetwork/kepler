package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/x/epochs/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.EpochsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/x/committee/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CommitteeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

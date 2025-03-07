package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/x/instant/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.InstantKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

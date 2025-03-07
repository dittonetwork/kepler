package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/x/job/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.JobKeeper(t)

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{}, response)
}

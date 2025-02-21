package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/x/workflow/keeper"
	"kepler/x/workflow/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.WorkflowKeeper(t)
	cmtKeeper, _ := keepertest.CommitteeKeeper(t)
	jobKeeper, _ := keepertest.JobKeeper(t)
	return k, keeper.NewMsgServerImpl(k, cmtKeeper, jobKeeper), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

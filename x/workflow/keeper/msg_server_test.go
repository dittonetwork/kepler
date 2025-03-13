package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/x/workflow/keeper"
	"github.com/dittonetwork/kepler/x/workflow/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	jobKeeper, _ := keepertest.JobKeeper(t)
	k, ctx := keepertest.WorkflowKeeper(t, jobKeeper)
	cmtKeeper, _ := keepertest.CommitteeKeeper(t)

	return k, keeper.NewMsgServerImpl(k, cmtKeeper, jobKeeper), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

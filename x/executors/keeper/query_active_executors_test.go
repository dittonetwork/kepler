package keeper_test

import (
	"testing"

	"github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/require"
)

func TestGetActiveExecutors_Success(t *testing.T) {
	k, ctx := keeper.ExecutorsKeeper(t)

	// Prepare two executors: one active and one inactive.
	activeExec := types.Executor{
		Address:  "cosmos1active",
		IsActive: true,
	}
	inactiveExec := types.Executor{
		Address:  "cosmos1inactive",
		IsActive: false,
	}

	err := k.Executors.Set(ctx, activeExec.Address, activeExec)
	require.NoError(t, err)
	err = k.Executors.Set(ctx, inactiveExec.Address, inactiveExec)
	require.NoError(t, err)

	req := &types.QueryActiveExecutorsRequest{}
	resp, err := k.GetActiveExecutors(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	// Only the active executor should be returned.
	require.Len(t, resp.Executors, 1)
	require.Equal(t, "cosmos1active", resp.Executors[0].Address)
}

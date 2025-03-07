package keeper_test

import (
	"github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/x/workflow/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBaseKeeper_GetActiveAutomations(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t)

	automations, err := k.GetActiveAutomations(ctx, &types.QueryGetActiveAutomationsRequest{})
	require.NoError(t, err)
	require.NotNil(t, automations)

	// Add active automation
	automation := newValidAutomation()
	err = k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	automations, err = k.GetActiveAutomations(ctx, &types.QueryGetActiveAutomationsRequest{})
	require.NoError(t, err)
	require.NotNil(t, automations)
	require.Len(t, automations.ActiveAutomations, 1)
}

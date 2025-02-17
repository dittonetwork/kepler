package keeper_test

import (
	"kepler/testutil/keeper"
	"kepler/x/workflow/types"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestInsertAutomation tests the InsertAutomation function
func TestInsertAutomation(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t)

	// Add an automation
	automation := newTestAutomation(5, types.AutomationStatus_AUTOMATION_STATUS_ACTIVE)
	err := k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	// Retrieve and verify
	retrieved, err := k.GetAutomation(ctx, 5)
	require.NoError(t, err)
	require.Equal(t, automation, retrieved)
}

// TestSetAutomationStatus tests the SetAutomationStatus function
func TestSetAutomationStatus(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t)

	// Add an automation
	automation := newTestAutomation(5, types.AutomationStatus_AUTOMATION_STATUS_ACTIVE)
	err := k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	// Set the status
	err = k.SetAutomationStatus(ctx, 5, types.AutomationStatus_AUTOMATION_STATUS_PAUSED)
	require.NoError(t, err)

	// Retrieve and verify
	retrieved, err := k.GetAutomation(ctx, 5)
	require.NoError(t, err)
	require.Equal(t, types.AutomationStatus_AUTOMATION_STATUS_PAUSED, retrieved.Status)
}

// TestSetActiveAutomation tests the SetActiveAutomation function
func TestSetActiveAutomation(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t)

	// Add an active automation
	err := k.SetActiveAutomation(ctx, 5)
	require.NoError(t, err)

	// Retrieve and verify
	ids, err := k.GetActiveAutomationIDs(ctx)
	require.NoError(t, err)
	require.ElementsMatch(t, []uint64{5}, ids)
}

// TestGetActiveAutomations tests the GetActiveAutomations function
func TestGetActiveAutomations(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t)

	// Add multiple active automations
	err := k.SetActiveAutomation(ctx, 5)
	require.NoError(t, err)
	err = k.SetActiveAutomation(ctx, 31)
	require.NoError(t, err)
	err = k.SetActiveAutomation(ctx, 25)
	require.NoError(t, err)
	// Retrieve and verify
	ids, err := k.GetActiveAutomationIDs(ctx)
	require.NoError(t, err)
	require.ElementsMatch(t, []uint64{5, 31, 25}, ids)
}

// TestRemoveActiveAutomation tests the RemoveActiveAutomation function
func TestRemoveActiveAutomation(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t)

	// Add an active automation
	err := k.SetActiveAutomation(ctx, 5)
	require.NoError(t, err)
	// Add an active automation that we will remove
	err = k.SetActiveAutomation(ctx, 7)
	require.NoError(t, err)

	// Retrieve and verify
	ids, err := k.GetActiveAutomationIDs(ctx)
	require.NoError(t, err)
	require.ElementsMatch(t, []uint64{5, 7}, ids)

	// Remove the active automation
	err = k.RemoveActiveAutomation(ctx, 7)
	require.NoError(t, err)

	// Retrieve and verify
	ids, err = k.GetActiveAutomationIDs(ctx)
	require.NoError(t, err)
	require.ElementsMatch(t, []uint64{5}, ids)
}

// Helper function to create a test automation
func newTestAutomation(id uint64, status types.AutomationStatus) types.Automation {
	return types.Automation{
		Id:     id,
		Status: status,
	}
}

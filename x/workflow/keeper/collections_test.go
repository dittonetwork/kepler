package keeper_test

import (
	"testing"
	"time"

	"github.com/dittonetwork/kepler/testutil/keeper"
	jobTypes "github.com/dittonetwork/kepler/x/job/types"
	"github.com/dittonetwork/kepler/x/workflow/types"
	"github.com/dittonetwork/kepler/x/workflow/types/mock"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/require"
)

// TestInsertAutomation tests the InsertAutomation function
func TestInsertAutomation(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t, nil)

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
	k, ctx := keeper.WorkflowKeeper(t, nil)

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

// TestFindActiveAutomations tests the GetActiveAutomations function
func TestFindActiveAutomations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	j := mock.NewMockJobKeeper(ctrl)
	k, ctx := keeper.WorkflowKeeper(t, j)

	// Add multiple active automations
	automation := newTestAutomation(5, types.AutomationStatus_AUTOMATION_STATUS_ACTIVE)
	err := k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	automation = newTestAutomation(7, types.AutomationStatus_AUTOMATION_STATUS_ACTIVE)
	err = k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	automation = newTestAutomation(9, types.AutomationStatus_AUTOMATION_STATUS_ACTIVE)
	err = k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	// one of em will be paused
	automation = newTestAutomation(11, types.AutomationStatus_AUTOMATION_STATUS_PAUSED)
	err = k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	j.EXPECT().GetLastSuccessfulJobByAutomation(gomock.Any(), uint64(5)).Return(jobTypes.Job{}, nil)
	j.EXPECT().GetLastSuccessfulJobByAutomation(gomock.Any(), uint64(7)).Return(jobTypes.Job{}, nil)
	j.EXPECT().GetLastSuccessfulJobByAutomation(gomock.Any(), uint64(9)).Return(jobTypes.Job{}, nil)
	// Retrieve and verify
	activeAutomations, err := k.FindActiveAutomations(ctx)
	require.NoError(t, err)
	require.Len(t, activeAutomations, 3)
}

// Helper function to create a test automation
func newTestAutomation(id uint64, status types.AutomationStatus) types.Automation {
	return types.Automation{
		Id:     id,
		Status: status,
	}
}

func newValidAutomation() types.Automation {
	triggers := []*types.Trigger{
		{
			Trigger: &types.Trigger_Count{Count: &types.CountTrigger{
				RepeatCount: 1,
			}},
		},
	}
	actions := []*types.Action{
		{
			&types.Action_OnChain{OnChain: &types.OnChainAction{
				ContractAddress: "0x1234",
				ChainId:         "1",
				TxCallData:      []byte("tx_call_data"),
			}},
		},
	}

	expireAt := time.Now().Add(time.Hour).Unix()

	return types.Automation{
		Id:       1,
		Triggers: triggers,
		Actions:  actions,
		Status:   types.AutomationStatus_AUTOMATION_STATUS_ACTIVE,
		ExpireAt: expireAt,
	}
}

func TestCancelAutomation(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t, nil)

	// Add an automation
	automation := newTestAutomation(5, types.AutomationStatus_AUTOMATION_STATUS_ACTIVE)
	err := k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	// Cancel Active automation
	err = k.CancelAutomation(ctx, 5)
	require.NoError(t, err)

	// Retrieve and verify
	automation, err = k.GetAutomation(ctx, 5)
	require.NoError(t, err)
	require.Equal(t, types.AutomationStatus_AUTOMATION_STATUS_CANCELED, automation.Status)
}

func TestCancelAutomationFail(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t, nil)

	// Add an automation
	automation := newTestAutomation(5, types.AutomationStatus_AUTOMATION_STATUS_DONE)
	err := k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	// Cancel Active automation
	err = k.CancelAutomation(ctx, 5)
	require.Error(t, err)
}

func TestActivateAutomation(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t, nil)

	// Add an automation
	automation := newTestAutomation(5, types.AutomationStatus_AUTOMATION_STATUS_PAUSED)
	err := k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	// Activate Active automation
	err = k.ActivateAutomation(ctx, 5)
	require.NoError(t, err)

	// Retrieve and verify
	automation, err = k.GetAutomation(ctx, 5)
	require.NoError(t, err)
	require.Equal(t, types.AutomationStatus_AUTOMATION_STATUS_ACTIVE, automation.Status)
}

func TestActivateAutomationFail(t *testing.T) {
	k, ctx := keeper.WorkflowKeeper(t, nil)

	// Add an automation
	automation := newTestAutomation(5, types.AutomationStatus_AUTOMATION_STATUS_DONE)
	err := k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	// Activate Active automation
	err = k.ActivateAutomation(ctx, 5)
	require.Error(t, err)
}

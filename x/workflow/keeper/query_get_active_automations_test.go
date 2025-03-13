package keeper_test

import (
	"testing"

	"github.com/dittonetwork/kepler/testutil/keeper"
	jobTypes "github.com/dittonetwork/kepler/x/job/types"
	"github.com/dittonetwork/kepler/x/workflow/types"
	"github.com/dittonetwork/kepler/x/workflow/types/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestBaseKeeper_GetActiveAutomations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	j := mock.NewMockJobKeeper(ctrl)
	k, ctx := keeper.WorkflowKeeper(t, j)

	automations, err := k.GetActiveAutomations(ctx, &types.QueryGetActiveAutomationsRequest{})
	require.NoError(t, err)
	require.NotNil(t, automations)

	// Add active automation
	automation := newValidAutomation()
	err = k.InsertAutomation(ctx, automation)
	require.NoError(t, err)

	j.EXPECT().GetLastSuccessfulJobByAutomation(gomock.Any(), uint64(1)).Return(jobTypes.Job{Id: 256, AutomationId: 1}, nil)
	automations, err = k.GetActiveAutomations(ctx, &types.QueryGetActiveAutomationsRequest{})
	require.NoError(t, err)
	require.NotNil(t, automations)
	require.Len(t, automations.ActiveAutomations, 1)
	require.Equal(t, uint64(256), automations.ActiveAutomations[0].LastSuccessfulJob.Id)
}

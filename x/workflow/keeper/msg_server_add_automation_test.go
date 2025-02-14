package keeper_test

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"kepler/x/workflow/keeper"
	"kepler/x/workflow/types"
	"kepler/x/workflow/types/mock"
	"testing"
	"time"
)

// TestAddAutomation tests the InsertAutomation function
func TestAddAutomationSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := testutil.DefaultContext(
		storetypes.NewKVStoreKey(types.StoreKey),
		storetypes.NewTransientStoreKey("transient_"+types.StoreKey),
	)

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
				ContractAddress: []byte("0x1234"),
				ChainId:         "1",
				TxCallData:      []byte("tx_call_data"),
			}},
		},
	}

	expireAt := time.Now().Add(time.Hour).Unix()
	automation := types.Automation{
		Id:       1,
		Triggers: triggers,
		Actions:  actions,
		Status:   types.AutomationStatus_AUTOMATION_STATUS_ACTIVE,
		ExpireAt: expireAt,
	}

	// Create a mock keeper
	mockKeeper := mock.NewMockKeeper(ctrl)
	mockKeeper.EXPECT().GetNextAutomationID(ctx).Return(uint64(1), nil)
	mockKeeper.EXPECT().InsertAutomation(ctx, automation).Return(nil)

	// Create a new message server
	msgServer := keeper.NewMsgServerImpl(mockKeeper)

	// Create a new message
	msg := &types.MsgAddAutomation{
		Triggers: triggers,
		Actions:  actions,
		ExpireAt: expireAt,
	}

	// Call the AddAutomation function
	resp, err := msgServer.AddAutomation(ctx, msg)
	if err != nil {
		t.Fatalf("AddAutomation failed: %v", err)
	}

	// Check the response
	if resp == nil {
		t.Fatalf("AddAutomation response is nil")
	}

	assert.Equal(t, uint64(1), resp.Id)
}

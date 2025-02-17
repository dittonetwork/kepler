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
)

// TestAddAutomation tests the InsertAutomation function
func TestAddAutomationSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := testutil.DefaultContext(
		storetypes.NewKVStoreKey(types.StoreKey),
		storetypes.NewTransientStoreKey("transient_"+types.StoreKey),
	)

	automation := newValidAutomation()

	// Create a mock keeper
	mockKeeper := mock.NewMockKeeper(ctrl)
	mockKeeper.EXPECT().GetNextAutomationID(ctx).Return(uint64(1), nil)
	mockKeeper.EXPECT().InsertAutomation(ctx, automation).Return(nil)

	// Create a new message server
	msgServer := keeper.NewMsgServerImpl(mockKeeper)

	// Create a new message
	msg := &types.MsgAddAutomation{
		Triggers: automation.Triggers,
		Actions:  automation.Actions,
		ExpireAt: automation.ExpireAt,
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

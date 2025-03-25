package keeper_test

import (
	"errors"
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	jobTypes "github.com/dittonetwork/kepler/x/job/types"
	"github.com/dittonetwork/kepler/x/workflow/keeper"
	"github.com/dittonetwork/kepler/x/workflow/types"
	"github.com/dittonetwork/kepler/x/workflow/types/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestMsgServer_SubmitJobResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := testutil.DefaultContext(
		storetypes.NewKVStoreKey(types.StoreKey),
		storetypes.NewTransientStoreKey("transient_"+types.StoreKey),
	)

	mockKeeper := mock.NewMockKeeper(ctrl)

	cmtKeeper := mock.NewMockCommitteeKeeper(ctrl)
	jobKeeper := mock.NewMockJobKeeper(ctrl)
	msgServer := keeper.NewMsgServerImpl(mockKeeper, cmtKeeper, jobKeeper)

	cases := map[string]struct {
		prepare     func()
		expectedErr error
	}{
		"automation not found": {
			prepare: func() {
				mockKeeper.EXPECT().GetAutomation(
					ctx,
					uint64(1),
				).Return(types.Automation{}, types.ErrAutomationNotFound)
			},
			expectedErr: types.ErrAutomationNotFound,
		},
		"automation is not active": {
			prepare: func() {
				mockKeeper.EXPECT().GetAutomation(
					ctx,
					uint64(1),
				).Return(types.Automation{
					Status: types.AutomationStatus_AUTOMATION_STATUS_CANCELED,
				}, nil)
			},
			expectedErr: errors.New("automation is not active"),
		},
		"job created": {
			prepare: func() {
				mockKeeper.EXPECT().GetAutomation(
					ctx,
					uint64(1),
				).Return(types.Automation{
					Status: types.AutomationStatus_AUTOMATION_STATUS_ACTIVE,
				}, nil)

				jobKeeper.EXPECT().CreateJob(
					ctx,
					jobTypes.Job_STATUS_EXECUTED,
					"committee123",
					"chain1",
					uint64(1),
					"0xabcdef",
					"cosmos1validaddress",
					uint64(1625150701),
					uint64(1625150801),
					uint64(1625150901),
					[][]byte{[]byte("signature1"), []byte("signature2")},
					[]byte("payload data"),
				).Return(nil)
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tc.prepare()

			msg := &types.MsgSubmitJobResult{
				Creator:      "cosmos1validaddress",
				Status:       jobTypes.Job_STATUS_EXECUTED.String(),
				CommitteeId:  "committee123",
				ChainId:      "chain1",
				AutomationId: 1,
				TxHash:       "0xabcdef",
				CreatedAt:    1625150701,
				ExecutedAt:   1625150801,
				SignedAt:     1625150901,
				Signs:        [][]byte{[]byte("signature1"), []byte("signature2")},
				Payload:      []byte("payload data"),
			}

			_, err := msgServer.SubmitJobResult(ctx, msg)
			if tc.expectedErr != nil {
				require.Contains(t, err.Error(), tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

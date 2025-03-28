package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/x/committee/types"
	"github.com/dittonetwork/kepler/x/committee/types/mock"
)

func TestCreateCommittee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	executorsMock := mock.NewMockExecutors(ctrl)

	testCases := []struct {
		name            string
		epoch           uint32
		latestEpoch     uint32
		committeeExists bool
		mockSetup       func()
		expectedError   error
	}{
		{
			name:            "Committee already exists",
			epoch:           10,
			committeeExists: true,
			mockSetup:       func() {},
			expectedError:   types.ErrCommitteeAlreadyExists,
		},
		{
			name:            "Latest epoch less than or equal to given epoch",
			epoch:           10,
			latestEpoch:     5,
			committeeExists: false,
			mockSetup:       func() {},
			expectedError:   types.ErrInvalidEpoch,
		},
		{
			name:            "Failed to get emergency executors",
			epoch:           5,
			latestEpoch:     10,
			committeeExists: false,
			mockSetup: func() {
				executorsMock.EXPECT().
					GetEmergencyExecutors(gomock.Any()).
					Return(nil, types.ErrInvalidSigner)
			},
			expectedError: types.ErrInvalidSigner,
		},
		{
			name:            "Success case",
			epoch:           5,
			latestEpoch:     10,
			committeeExists: false,
			mockSetup: func() {
				// Create sample executors to return
				addr1, _ := sdk.AccAddressFromBech32("cosmos1v9jxgun9wdenzc33zgq")
				addr2, _ := sdk.AccAddressFromBech32("cosmos1s4ycalgh3gjemd5z4qj4w9n8vz3cplecafqa97")

				executors := []types.ExecutorI{
					&testExecutor{
						address:     addr1,
						votingPower: 10,
					},
					&testExecutor{
						address:     addr2,
						votingPower: 20,
					},
				}

				executorsMock.EXPECT().
					GetEmergencyExecutors(gomock.Any()).
					Return(executors, nil)
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create keeper with mocks
			k, ctx := keepertest.CommitteeKeeperWithMocks(t, executorsMock)

			// Setup the test case
			tc.mockSetup()

			// Set latest epoch
			if tc.latestEpoch > 0 {
				require.NoError(t, k.LatestEpoch.Set(ctx, tc.latestEpoch))
			}

			// Set committee existence
			if !tc.committeeExists {
				require.NoError(t, k.Committees.Set(ctx, tc.epoch, types.Committee{}))
			}

			// Execute the method
			committee, err := k.CreateCommittee(ctx, tc.epoch)

			// Check the results
			if tc.expectedError != nil {
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.epoch, committee.Epoch)
				require.True(t, committee.IsEmergency)
				require.Len(t, committee.Executors, 2)

				// Check that addresses and voting powers match what we expect
				// Don't compare the actual address strings as they will be formatted
				require.Equal(t, uint32(10), committee.Executors[0].VotingPower)
				require.Equal(t, uint32(20), committee.Executors[1].VotingPower)
			}
		})
	}
}

// testExecutor is a mock implementation of ExecutorI for testing
type testExecutor struct {
	address     sdk.AccAddress
	votingPower uint32
}

func (e *testExecutor) GetAddress() sdk.AccAddress {
	return e.address
}

func (e *testExecutor) GetVotingPower() uint32 {
	return e.votingPower
}

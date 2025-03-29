package keeper_test

import (
	"github.com/dittonetwork/kepler/x/committee/types"
	executortypes "github.com/dittonetwork/kepler/x/executors/types"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"
	"go.uber.org/mock/gomock"
)

func (s *TestSuite) TestCreateCommittee() {
	testCases := []struct {
		name            string
		epoch           uint32
		lastEpoch       uint32
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
			name:            "Epoch equal to last epoch",
			epoch:           10,
			lastEpoch:       10,
			committeeExists: false,
			mockSetup:       func() {},
			expectedError:   types.ErrInvalidEpoch,
		},
		{
			name:            "Epoch less than last epoch",
			epoch:           5,
			lastEpoch:       10,
			committeeExists: false,
			mockSetup:       func() {},
			expectedError:   types.ErrInvalidEpoch,
		},
		{
			name:            "Failed to get emergency executors",
			epoch:           15,
			lastEpoch:       10,
			committeeExists: false,
			mockSetup: func() {
				s.executorKeeper.EXPECT().
					GetEmergencyExecutors(gomock.Any()).
					Return(nil, types.ErrInvalidSigner)
			},
			expectedError: types.ErrInvalidSigner,
		},
		{
			name:            "Success case - epoch greater than last epoch",
			epoch:           11,
			lastEpoch:       10,
			committeeExists: false,
			mockSetup: func() {
				executors := []executortypes.Executor{
					{
						Address:      "cosmos1v9jxgun9wdenzc33zgq23q8r9hfv2z4x0762r",
						OwnerAddress: "cosmosvaloper1w7f3xx7e75p4l7qdym5msqem9rd4dyc4mq79dm",
						IsActive:     true,
					},
					{
						Address:      "cosmos1s4ycalgh3gjemd5z4qj4w9n8vz3cplecafqa97",
						OwnerAddress: "cosmosvaloper1w7f3xx7e75p4l7qdym5msqem9rd4dyc4mq79dm",
						IsActive:     true,
					},
				}
				s.executorKeeper.EXPECT().
					GetEmergencyExecutors(gomock.Any()).
					Return(executors, nil)

				// Mock GetValidator calls for each executor
				s.restakingKeeper.EXPECT().
					GetValidator(gomock.Any(), gomock.Any()).
					Return(restakingtypes.Validator{VotingPower: 10}, nil).
					AnyTimes()
			},
			expectedError: nil,
		},
		{
			name:            "Boundary case - epoch exactly one more than last epoch",
			epoch:           11,
			lastEpoch:       10,
			committeeExists: false,
			mockSetup: func() {
				executors := []executortypes.Executor{
					{
						Address:      "cosmos1v9jxgun9wdenzc33zgq23q8r9hfv2z4x0762r",
						OwnerAddress: "cosmosvaloper1w7f3xx7e75p4l7qdym5msqem9rd4dyc4mq79dm",
						IsActive:     true,
					},
					{
						Address:      "cosmos1s4ycalgh3gjemd5z4qj4w9n8vz3cplecafqa97",
						OwnerAddress: "cosmosvaloper1w7f3xx7e75p4l7qdym5msqem9rd4dyc4mq79dm",
						IsActive:     true,
					},
				}
				s.executorKeeper.EXPECT().
					GetEmergencyExecutors(gomock.Any()).
					Return(executors, nil)

				// Mock GetValidator calls for each executor
				s.restakingKeeper.EXPECT().
					GetValidator(gomock.Any(), gomock.Any()).
					Return(restakingtypes.Validator{VotingPower: 20}, nil).
					AnyTimes()
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Reset the keeper state for each test case
			s.SetupTest()

			// Get keeper from suite
			k := s.keeper
			ctx := s.ctx

			// Setup the test case
			tc.mockSetup()

			// Set last epoch
			if tc.lastEpoch > 0 {
				s.Require().NoError(k.LastEpoch.Set(ctx, tc.lastEpoch))
			}

			// Set committee existence
			if tc.committeeExists {
				s.Require().NoError(k.Committees.Set(ctx, tc.epoch, types.Committee{}))
			}

			// Call the actual CreateCommittee method
			committee, err := k.CreateCommittee(ctx, tc.epoch)

			// Check the results
			if tc.expectedError != nil {
				s.Require().ErrorIs(err, tc.expectedError)
				return
			}

			// For successful cases
			s.Require().NoError(err)
			s.Require().Equal(tc.epoch, committee.Epoch)
			s.Require().True(committee.IsEmergency)
			s.Require().Equal(ctx.HeaderInfo().Hash, committee.Seed)
			s.Require().Len(committee.Executors, 2)

			// Verify that LastEpoch was updated to the new epoch value
			lastEpoch, err := k.LastEpoch.Get(ctx)
			s.Require().NoError(err)
			s.Require().Equal(tc.epoch, lastEpoch, "LastEpoch should be updated to the new epoch value")
		})
	}
}

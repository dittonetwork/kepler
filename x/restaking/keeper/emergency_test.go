package keeper_test

import (
	"errors"

	"github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func (s *TestSuite) TestGetActiveEmergencyValidators() {
	testCases := []struct {
		name            string
		setupMock       func()
		expectedResults []types.Validator
		expectedError   error
	}{
		{
			name: "successfully retrieves only bonded validators",
			setupMock: func() {
				validators := []types.Validator{
					{
						OperatorAddress: "validator1",
						Status:          types.Bonded,
					},
					{
						OperatorAddress: "validator2",
						Status:          types.Unbonded,
					},
					{
						OperatorAddress: "validator3",
						Status:          types.Bonded,
					},
				}

				s.repository.EXPECT().
					GetEmergencyValidators(gomock.Any()).
					Return(validators, nil)
			},
			expectedResults: []types.Validator{
				{
					OperatorAddress: "validator1",
					Status:          types.Bonded,
				},
				{
					OperatorAddress: "validator3",
					Status:          types.Bonded,
				},
			},
			expectedError: nil,
		},
		{
			name: "returns empty list when no validators are bonded",
			setupMock: func() {
				validators := []types.Validator{
					{
						OperatorAddress: "validator1",
						Status:          types.Unbonded,
					},
					{
						OperatorAddress: "validator2",
						Status:          types.Unbonding,
					},
				}

				s.repository.EXPECT().
					GetEmergencyValidators(gomock.Any()).
					Return(validators, nil)
			},
			expectedResults: []types.Validator{},
			expectedError:   nil,
		},
		{
			name: "handles error from repository",
			setupMock: func() {
				s.repository.EXPECT().
					GetEmergencyValidators(gomock.Any()).
					Return(nil, errors.New("repository error"))
			},
			expectedResults: nil,
			expectedError:   errors.New("repository error"),
		},
		{
			name: "returns empty list when repository returns empty list",
			setupMock: func() {
				s.repository.EXPECT().
					GetEmergencyValidators(gomock.Any()).
					Return([]types.Validator{}, nil)
			},
			expectedResults: []types.Validator{},
			expectedError:   nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup mock expectations
			tc.setupMock()

			// Call the function being tested
			validators, err := s.keeper.GetActiveEmergencyValidators(s.ctx)

			// Check results
			if tc.expectedError != nil {
				require.Error(s.T(), err)
				require.Equal(s.T(), tc.expectedError.Error(), err.Error())
			} else {
				require.NoError(s.T(), err)
				require.Equal(s.T(), len(tc.expectedResults), len(validators))
				for i, validator := range validators {
					require.Equal(s.T(), tc.expectedResults[i].OperatorAddress, validator.OperatorAddress)
					require.Equal(s.T(), tc.expectedResults[i].Status, validator.Status)
				}
			}
		})
	}
}

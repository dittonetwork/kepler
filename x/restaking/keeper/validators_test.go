package keeper_test

import (
	"errors"
	"time"

	"github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const validBlockHash = "0x1234567890123456789012345678901234567890123456789012345678901234"

func (s *TestSuite) TestNeedValidatorsUpdate() {
	testCases := []struct {
		name          string
		epoch         int64
		setupMock     func()
		expectedNeed  bool
		expectedError error
	}{
		{
			name:  "needs update when epoch > last epoch",
			epoch: 5,
			setupMock: func() {
				lastUpdate := types.UpdateInfo{
					EpochNum:    3,
					BlockHeight: 100,
					BlockHash:   validBlockHash,
					Timestamp:   time.Now(),
				}

				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(lastUpdate, nil)
			},
			expectedNeed:  true,
			expectedError: nil,
		},
		{
			name:  "does not need update when epoch == last epoch",
			epoch: 3,
			setupMock: func() {
				lastUpdate := types.UpdateInfo{
					EpochNum:    3,
					BlockHeight: 100,
					BlockHash:   validBlockHash,
					Timestamp:   time.Now(),
				}

				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(lastUpdate, nil)
			},
			expectedNeed:  false,
			expectedError: nil,
		},
		{
			name:  "does not need update when epoch < last epoch",
			epoch: 2,
			setupMock: func() {
				lastUpdate := types.UpdateInfo{
					EpochNum:    3,
					BlockHeight: 100,
					BlockHash:   validBlockHash,
					Timestamp:   time.Now(),
				}

				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(lastUpdate, nil)
			},
			expectedNeed:  false,
			expectedError: nil,
		},
		{
			name:  "returns error when GetLastUpdate fails",
			epoch: 5,
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{}, errors.New("failed to get last update"))
			},
			expectedNeed:  false,
			expectedError: errors.New("failed to get last update"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup mock expectations
			tc.setupMock()

			// Call the function being tested
			needUpdate, err := s.keeper.NeedValidatorsUpdate(s.ctx, tc.epoch)

			// Check results
			if tc.expectedError != nil {
				require.Error(s.T(), err)
				require.Equal(s.T(), tc.expectedError.Error(), err.Error())
			} else {
				require.NoError(s.T(), err)
				require.Equal(s.T(), tc.expectedNeed, needUpdate)
			}
		})
	}
}

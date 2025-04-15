package keeper_test

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func (s *TestSuite) TestMsgUpdateValidatorsSet() {

	validAuthority := s.keeper.GetAuthority()
	invalidAuthority := "invalid_authority"
	validOperatorAddress := "0x1234567890123456789012345678901234567890"

	now := time.Now()
	validPubkey := &codectypes.Any{
		TypeUrl: "/cosmos.crypto.ed25519.PubKey",
		Value:   []byte("pubkey1"),
	}

	testCases := []struct {
		name           string
		setupMock      func()
		msg            *types.MsgUpdateValidatorsSet
		expectedError  error
		expectedResult *types.MsgUpdateValidatorsSetResponse
	}{
		{
			name: "successful validators update",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    1,
						BlockHeight: 100,
						BlockHash:   validBlockHash,
						Timestamp:   now,
					}, nil)

				// Prepare for makeDeltaUpdates
				s.repository.EXPECT().
					GetAllValidators(gomock.Any()).
					Return([]types.Validator{}, nil)

				// Prepare for processUpdatedValidators, processCreatedOperators, processDeletedOperators
				s.repository.EXPECT().
					SetPendingOperator(gomock.Any(), gomock.Eq(validOperatorAddress), gomock.Any()).
					Return(nil).
					AnyTimes()

				s.repository.EXPECT().
					RemovePendingOperator(gomock.Any(), gomock.Any()).
					Return(nil).
					AnyTimes()

				s.repository.EXPECT().
					GetValidatorByEvmAddr(gomock.Any(), gomock.Any()).
					Return(types.Validator{}, nil).
					AnyTimes()

				s.repository.EXPECT().
					RemoveValidatorByOperatorAddr(gomock.Any(), gomock.Any()).
					Return(nil).
					AnyTimes()

				s.repository.EXPECT().
					SetValidator(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil).
					AnyTimes()

				s.repository.EXPECT().
					SetLastUpdate(gomock.Any(), gomock.Any()).
					Return(nil).
					AnyTimes()
			},
			msg: &types.MsgUpdateValidatorsSet{
				Authority: validAuthority,
				Operators: []types.Operator{
					{
						Address:         validOperatorAddress,
						Status:          types.Bonded,
						ConsensusPubkey: validPubkey,
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2,
					BlockHeight: 200,
					BlockHash:   validBlockHash,
					Timestamp:   now.Add(time.Hour),
				},
			},
			expectedError:  nil,
			expectedResult: &types.MsgUpdateValidatorsSetResponse{},
		},
		{
			name: "invalid authority",
			setupMock: func() {
				// No mocks needed for this test as the error occurs before mocked functions are called
			},
			msg: &types.MsgUpdateValidatorsSet{
				Authority: invalidAuthority,
				Operators: []types.Operator{
					{
						Address:         validOperatorAddress,
						Status:          types.Bonded,
						ConsensusPubkey: validPubkey,
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2,
					BlockHeight: 200,
					BlockHash:   validBlockHash,
					Timestamp:   now,
				},
			},
			expectedError:  govtypes.ErrInvalidSigner,
			expectedResult: nil,
		},
		{
			name: "block height lower than last update",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    1,
						BlockHeight: 300, // Higher than in message
						BlockHash:   validBlockHash,
						Timestamp:   now,
					}, nil)
			},
			msg: &types.MsgUpdateValidatorsSet{
				Authority: validAuthority,
				Operators: []types.Operator{
					{
						Address:         validOperatorAddress,
						Status:          types.Bonded,
						ConsensusPubkey: validPubkey,
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2,
					BlockHeight: 200, // Lower than last update
					BlockHash:   validBlockHash,
					Timestamp:   now,
				},
			},
			expectedError:  types.ErrUpdateValidator,
			expectedResult: nil,
		},
		{
			name: "epoch number lower than last update",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    3, // Higher than in message
						BlockHeight: 100,
						BlockHash:   validBlockHash,
						Timestamp:   now,
					}, nil)
			},
			msg: &types.MsgUpdateValidatorsSet{
				Authority: validAuthority,
				Operators: []types.Operator{
					{
						Address:         validOperatorAddress,
						Status:          types.Bonded,
						ConsensusPubkey: validPubkey,
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2, // Lower than last update
					BlockHeight: 200,
					BlockHash:   validBlockHash,
					Timestamp:   now,
				},
			},
			expectedError:  types.ErrUpdateValidator,
			expectedResult: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup mocks
			tc.setupMock()

			// Call function
			res, err := s.msgServer.UpdateValidatorsSet(s.ctx, tc.msg)

			// Check results
			if tc.expectedError != nil {
				require.Error(s.T(), err)
				require.ErrorIs(s.T(), err, tc.expectedError)
				require.Nil(s.T(), res)
			} else {
				require.NoError(s.T(), err)
				require.NotNil(s.T(), res)
				require.Equal(s.T(), tc.expectedResult, res)
			}
		})
	}
}

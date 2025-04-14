package keeper_test

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func (s *TestSuite) TestValidateUpdateValidatorSet() {
	testCases := []struct {
		name          string
		setupMock     func()
		update        types.ValidatorsUpdate
		expectedError error
	}{
		{
			name: "fails when block height is lower than last update",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    1,
						BlockHeight: 200, // Higher than the update
						BlockHash:   validBlockHash,
						Timestamp:   time.Now(),
					}, nil)
			},
			update: types.ValidatorsUpdate{
				Operators: []types.Operator{
					{
						Address: "0x111222333",
						Status:  types.Unbonding,
						ConsensusPubkey: &codectypes.Any{
							TypeUrl: "/cosmos.crypto.ed25519.PubKey",
							Value:   []byte("pubkey1"),
						},
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2,
					BlockHeight: 100, // Lower than the last update
					BlockHash:   validBlockHash,
					Timestamp:   time.Now(),
				},
			},
			expectedError: types.ErrUpdateValidator,
		},
		{
			name: "fails when epoch is lower than last update",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    3, // Higher than the update
						BlockHeight: 100,
						BlockHash:   validBlockHash,
						Timestamp:   time.Now(),
					}, nil)
			},
			update: types.ValidatorsUpdate{
				Operators: []types.Operator{
					{
						Address: "0x111222333",
						Status:  types.Unbonding,
						ConsensusPubkey: &codectypes.Any{
							TypeUrl: "/cosmos.crypto.ed25519.PubKey",
							Value:   []byte("pubkey1"),
						},
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2, // Lower than the last update
					BlockHeight: 200,
					BlockHash:   validBlockHash,
					Timestamp:   time.Now(),
				},
			},
			expectedError: types.ErrUpdateValidator,
		},
		{
			name: "fails when block hash is invalid",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    1,
						BlockHeight: 100,
						BlockHash:   validBlockHash,
						Timestamp:   time.Now(),
					}, nil)
			},
			update: types.ValidatorsUpdate{
				Operators: []types.Operator{
					{
						Address: "0x111222333",
						Status:  types.Unbonding,
						ConsensusPubkey: &codectypes.Any{
							TypeUrl: "/cosmos.crypto.ed25519.PubKey",
							Value:   []byte("pubkey1"),
						},
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2,
					BlockHeight: 200,
					BlockHash:   "invalid", // Invalid block hash
					Timestamp:   time.Now(),
				},
			},
			expectedError: types.ErrUpdateValidator,
		},
		{
			name: "fails when operator address is empty",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    1,
						BlockHeight: 100,
						BlockHash:   validBlockHash,
						Timestamp:   time.Now(),
					}, nil)
			},
			update: types.ValidatorsUpdate{
				Operators: []types.Operator{
					{
						Address: "", // Empty operator address
						Status:  types.Unbonding,
						ConsensusPubkey: &codectypes.Any{
							TypeUrl: "/cosmos.crypto.ed25519.PubKey",
							Value:   []byte("pubkey1"),
						},
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2,
					BlockHeight: 200,
					BlockHash:   validBlockHash,
					Timestamp:   time.Now(),
				},
			},
			expectedError: types.ErrUpdateValidator,
		},
		{
			name: "fails when operator address is not a valid ethereum address",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    1,
						BlockHeight: 100,
						BlockHash:   validBlockHash,
						Timestamp:   time.Now(),
					}, nil)
			},
			update: types.ValidatorsUpdate{
				Operators: []types.Operator{
					{
						Address: "not-hex-address", // Invalid Ethereum address
						Status:  types.Unbonding,
						ConsensusPubkey: &codectypes.Any{
							TypeUrl: "/cosmos.crypto.ed25519.PubKey",
							Value:   []byte("pubkey1"),
						},
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2,
					BlockHeight: 200,
					BlockHash:   validBlockHash,
					Timestamp:   time.Now(),
				},
			},
			expectedError: types.ErrUpdateValidator,
		},
		{
			name: "fails when consensus public key is empty",
			setupMock: func() {
				s.repository.EXPECT().
					GetLastUpdate(gomock.Any()).
					Return(types.UpdateInfo{
						EpochNum:    1,
						BlockHeight: 100,
						BlockHash:   validBlockHash,
						Timestamp:   time.Now(),
					}, nil)
			},
			update: types.ValidatorsUpdate{
				Operators: []types.Operator{
					{
						Address:         "0x111222333",
						Status:          types.Unbonding,
						ConsensusPubkey: nil, // Empty consensus public key
					},
				},
				Info: types.UpdateInfo{
					EpochNum:    2,
					BlockHeight: 200,
					BlockHash:   validBlockHash,
					Timestamp:   time.Now(),
				},
			},
			expectedError: types.ErrUpdateValidator,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup mock expectations
			tc.setupMock()

			// Call validateUpdateValidatorSet method indirectly through UpdateValidatorSet
			err := s.keeper.UpdateValidatorSet(s.ctx, tc.update)

			// Check results - all test cases should be failing with expected error
			require.Error(s.T(), err)
			require.ErrorIs(s.T(), err, tc.expectedError)
		})
	}
}

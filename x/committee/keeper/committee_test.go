package keeper_test

import (
	"testing"

	cmtcrypto "github.com/cometbft/cometbft/crypto"
	prototypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdksecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
			mockSetup: func() {
				s.repo.EXPECT().
					HasCommittee(gomock.Any(), gomock.Any()).
					Return(true, nil)
			},
			expectedError: types.ErrCommitteeAlreadyExists,
		},
		{
			name:            "Epoch equal to last epoch",
			epoch:           10,
			lastEpoch:       10,
			committeeExists: false,
			mockSetup: func() {
				s.repo.EXPECT().
					HasCommittee(gomock.Any(), gomock.Any()).
					Return(false, nil)
				s.repo.EXPECT().
					GetLastEpoch(gomock.Any()).
					Return(uint32(10), nil)
			},
			expectedError: types.ErrInvalidEpoch,
		},
		{
			name:            "Epoch less than last epoch",
			epoch:           5,
			lastEpoch:       10,
			committeeExists: false,
			mockSetup: func() {
				s.repo.EXPECT().
					HasCommittee(gomock.Any(), gomock.Any()).
					Return(false, nil)
				s.repo.EXPECT().
					GetLastEpoch(gomock.Any()).
					Return(uint32(9), nil)
			},
			expectedError: types.ErrInvalidEpoch,
		},
		{
			name:            "Failed to get emergency executors",
			epoch:           15,
			lastEpoch:       10,
			committeeExists: false,
			mockSetup: func() {
				s.repo.EXPECT().
					HasCommittee(gomock.Any(), gomock.Any()).
					Return(false, nil)
				s.repo.EXPECT().
					GetLastEpoch(gomock.Any()).
					Return(uint32(10), nil)
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
				s.repo.EXPECT().
					HasCommittee(gomock.Any(), gomock.Any()).
					Return(false, nil)
				s.repo.EXPECT().
					GetLastEpoch(gomock.Any()).
					Return(uint32(10), nil)

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

				committeeExecutors := make([]types.Executor, len(executors))
				for i, executor := range executors {
					committeeExecutors[i] = types.Executor{
						Address:     executor.Address,
						VotingPower: 10,
					}
				}

				address := s.keeper.GetMultisigAddress(committeeExecutors)
				s.repo.EXPECT().
					SetCommittee(gomock.Any(), uint32(11), types.Committee{
						Epoch:       11,
						Executors:   committeeExecutors,
						IsEmergency: true,
						Seed:        nil,
						Address:     address,
					}).
					Return(nil)

				s.repo.EXPECT().
					SetLastEpoch(gomock.Any(), uint32(11)).
					Return(nil)

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
				s.repo.EXPECT().
					HasCommittee(gomock.Any(), gomock.Any()).
					Return(false, nil)
				s.repo.EXPECT().
					GetLastEpoch(gomock.Any()).
					Return(uint32(10), nil)

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

				committeeExecutors := make([]types.Executor, len(executors))
				for i, executor := range executors {
					committeeExecutors[i] = types.Executor{
						Address:     executor.Address,
						VotingPower: 20,
					}
				}

				address := s.keeper.GetMultisigAddress(committeeExecutors)
				s.repo.EXPECT().
					SetCommittee(gomock.Any(), uint32(11), types.Committee{
						Epoch:       11,
						Executors:   committeeExecutors,
						IsEmergency: true,
						Seed:        nil,
						Address:     address,
					}).
					Return(nil)

				s.repo.EXPECT().
					SetLastEpoch(gomock.Any(), uint32(11)).
					Return(nil)

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
		})
	}
}

func (s *TestSuite) TestGetMultisigAddress() {
	testCases := []struct {
		name string
		init func() ([]types.Executor, string)
	}{
		{
			name: "success case",
			init: func() ([]types.Executor, string) {
				executors := []types.Executor{
					{PublicKey: generatePubKey(s.T())},
					{PublicKey: generatePubKey(s.T())},
					{PublicKey: generatePubKey(s.T())},
					{PublicKey: generatePubKey(s.T())},
				}

				anyPubKeys := make([]*prototypes.Any, 0, len(executors))

				for _, executor := range executors {
					var pk sdksecp.PubKey
					copy(pk.Key, executor.PublicKey)

					any, err := prototypes.NewAnyWithValue(&pk)
					s.Require().NoError(err)

					anyPubKeys = append(anyPubKeys, any)
				}
				m := multisig.LegacyAminoPubKey{
					Threshold: 3,
					PubKeys:   anyPubKeys,
				}
				address := cmtcrypto.AddressHash(m.Bytes())

				return executors, sdk.AccAddress(address).String()
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			executors, tcAddress := tc.init()

			address := s.keeper.GetMultisigAddress(executors)

			_, err := sdk.AccAddressFromBech32(address)
			s.Require().NoError(err)
			s.Require().NotEmpty(address)
			s.Require().Equal(tcAddress, address)
		})
	}
}

func generatePubKey(t *testing.T) []byte {
	t.Helper()

	privKey := secp256k1.GenPrivKey()
	return privKey.PubKey().Bytes()
}

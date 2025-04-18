package keeper_test

import (
	"testing"

	cmtcrypto "github.com/cometbft/cometbft/crypto"
	prototypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdksecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/dittonetwork/kepler/x/committee/types"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"
	"go.uber.org/mock/gomock"
)

func (s *TestSuite) TestCreateCommittee() {
	testCases := []struct {
		name            string
		epoch           int64
		lastEpoch       int64
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
					Return(int64(10), nil)
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
					Return(int64(9), nil)
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
					Return(int64(10), nil)
				s.restakingKeeper.EXPECT().
					GetActiveEmergencyValidators(gomock.Any()).
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
					Return(int64(10), nil)
				s.restakingKeeper.EXPECT().
					GetActiveEmergencyValidators(gomock.Any()).
					Return([]restakingtypes.Validator{
						{
							OperatorAddress: s.alice.ValAddress.String(),
							VotingPower:     10,
						},
						{
							OperatorAddress: s.bob.ValAddress.String(),
							VotingPower:     10,
						},
					}, nil)

				executors := []types.Executor{
					{
						Address:     s.alice.Address.String(),
						VotingPower: 10,
					},
					{
						Address:     s.bob.Address.String(),
						VotingPower: 10,
					},
				}

				pubKeyAlice, err := prototypes.NewAnyWithValue(s.alice.PubKey)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), sdk.MustAccAddressFromBech32(executors[0].Address)).
					Return(&authtypes.BaseAccount{
						Address:       executors[0].Address,
						PubKey:        pubKeyAlice,
						AccountNumber: uint64(0),
					}).Times(2)

				pubKeyBob, err := prototypes.NewAnyWithValue(s.bob.PubKey)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), sdk.MustAccAddressFromBech32(executors[1].Address)).
					Return(&authtypes.BaseAccount{
						Address:       executors[1].Address,
						PubKey:        pubKeyBob,
						AccountNumber: uint64(1),
					}).Times(2)

				address, err := s.keeper.GetMultisigAddress(s.ctx, executors)
				s.Require().NoError(err)

				s.repo.EXPECT().
					SetCommittee(gomock.Any(), int64(11), types.Committee{
						Epoch:       11,
						Executors:   executors,
						IsEmergency: true,
						Seed:        nil,
						Address:     address,
					}).
					Return(nil)

				s.repo.EXPECT().
					SetLastEpoch(gomock.Any(), int64(11)).
					Return(nil)
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
					Return(int64(10), nil)

				s.restakingKeeper.EXPECT().
					GetActiveEmergencyValidators(gomock.Any()).
					Return([]restakingtypes.Validator{
						{
							OperatorAddress: s.alice.ValAddress.String(),
							VotingPower:     20,
						},
						{
							OperatorAddress: s.bob.ValAddress.String(),
							VotingPower:     20,
						},
					}, nil)

				executors := []types.Executor{
					{
						Address:     s.alice.Address.String(),
						VotingPower: 20,
					},
					{
						Address:     s.bob.Address.String(),
						VotingPower: 20,
					},
				}

				pubKeyAlice, err := prototypes.NewAnyWithValue(s.alice.PubKey)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), sdk.MustAccAddressFromBech32(executors[0].Address)).
					Return(&authtypes.BaseAccount{
						Address:       executors[0].Address,
						PubKey:        pubKeyAlice,
						AccountNumber: uint64(0),
					}).Times(2)

				pubKeyBob, err := prototypes.NewAnyWithValue(s.bob.PubKey)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), sdk.MustAccAddressFromBech32(executors[1].Address)).
					Return(&authtypes.BaseAccount{
						Address:       executors[1].Address,
						PubKey:        pubKeyBob,
						AccountNumber: uint64(1),
					}).Times(2)

				address, err := s.keeper.GetMultisigAddress(s.ctx, executors)
				s.Require().NoError(err)
				s.repo.EXPECT().
					SetCommittee(gomock.Any(), int64(11), types.Committee{
						Epoch:       11,
						Executors:   executors,
						IsEmergency: true,
						Seed:        nil,
						Address:     address,
					}).
					Return(nil)

				s.repo.EXPECT().
					SetLastEpoch(gomock.Any(), int64(11)).
					Return(nil)
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
				publicKeys := [][]byte{
					generatePubKey(s.T()),
					generatePubKey(s.T()),
					generatePubKey(s.T()),
					generatePubKey(s.T()),
				}

				anyPubKeys := make([]*prototypes.Any, 0, len(publicKeys))
				addresses := make([]sdk.AccAddress, 0, len(publicKeys))

				for _, p := range publicKeys {
					var pk sdksecp.PubKey
					pk.Key = make([]byte, len(p))
					copy(pk.Key[:], p)
					addresses = append(addresses, pk.Bytes())

					anyPk, err := prototypes.NewAnyWithValue(&pk)
					s.Require().NoError(err)

					anyPubKeys = append(anyPubKeys, anyPk)
				}
				m := multisig.LegacyAminoPubKey{
					Threshold: 3,
					PubKeys:   anyPubKeys,
				}
				address := cmtcrypto.AddressHash(m.Bytes())

				executors := make([]types.Executor, 0, len(addresses))
				for i, address := range addresses {
					executors = append(executors, types.Executor{
						Address:     address.String(),
						VotingPower: 10,
					})
					s.accountKeeper.EXPECT().
						GetAccount(gomock.Any(), address).
						Return(&authtypes.BaseAccount{
							Address:       address.String(),
							PubKey:        anyPubKeys[i],
							AccountNumber: uint64(i),
						})
				}

				return executors, sdk.AccAddress(address).String()
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			executors, tcAddress := tc.init()

			address, err := s.keeper.GetMultisigAddress(s.ctx, executors)
			s.Require().NoError(err)

			_, err = sdk.AccAddressFromBech32(address)
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

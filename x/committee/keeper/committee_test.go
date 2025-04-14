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
					Return(uint32(10), nil)
				s.restakingKeeper.EXPECT().
					GetActiveEmergencyValidators(gomock.Any()).
					Return([]restakingtypes.Validator{
						{
							OperatorAddress: "cosmos1kkyr80lkuku58h7e2v84egemscmem304mdra4f",
							VotingPower:     10,
						},
						{
							OperatorAddress: "cosmos1w3k88z6h0q5nsylh68xxdjh909qch82dlp3g5e",
							VotingPower:     10,
						},
					}, nil)

				executors := []types.Executor{
					{
						Address:     "cosmos1kkyr80lkuku58h7e2v84egemscmem304mdra4f",
						VotingPower: 10,
					},
					{
						Address:     "cosmos1w3k88z6h0q5nsylh68xxdjh909qch82dlp3g5e",
						VotingPower: 10,
					},
				}

				var pk1 secp256k1.PubKey
				pk1.Key = []byte{3, 35, 137, 5, 237, 152, 146, 214, 9, 136, 235, 5, 3, 177, 22, 250, 12, 125, 38, 196, 108, 145, 211, 192, 1, 55, 245, 134, 223, 206, 52, 19, 14}
				pubKey1, err := prototypes.NewAnyWithValue(&pk1)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), sdk.MustAccAddressFromBech32(executors[0].Address)).
					Return(&authtypes.BaseAccount{
						Address:       executors[0].Address,
						PubKey:        pubKey1,
						AccountNumber: uint64(0),
					}).Times(2)

				var pk2 secp256k1.PubKey
				pk2.Key = []byte{3, 251, 110, 231, 205, 75, 69, 84, 2, 243, 138, 31, 184, 101, 131, 102, 110, 80, 73, 82, 234, 132, 253, 101, 117, 228, 30, 163, 41, 134, 159, 196, 72}
				pubKey2, err := prototypes.NewAnyWithValue(&pk2)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), sdk.MustAccAddressFromBech32(executors[1].Address)).
					Return(&authtypes.BaseAccount{
						Address:       executors[1].Address,
						PubKey:        pubKey2,
						AccountNumber: uint64(1),
					}).Times(2)

				address, err := s.keeper.GetMultisigAddress(executors)
				s.Require().NoError(err)

				s.repo.EXPECT().
					SetCommittee(gomock.Any(), uint32(11), types.Committee{
						Epoch:       11,
						Executors:   executors,
						IsEmergency: true,
						Seed:        nil,
						Address:     address,
					}).
					Return(nil)

				s.repo.EXPECT().
					SetLastEpoch(gomock.Any(), uint32(11)).
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
					Return(uint32(10), nil)

				s.restakingKeeper.EXPECT().
					GetActiveEmergencyValidators(gomock.Any()).
					Return([]restakingtypes.Validator{
						{
							OperatorAddress: "cosmos1kkyr80lkuku58h7e2v84egemscmem304mdra4f",
							VotingPower:     20,
						},
						{
							OperatorAddress: "cosmos1w3k88z6h0q5nsylh68xxdjh909qch82dlp3g5e",
							VotingPower:     20,
						},
					}, nil)

				executors := []types.Executor{
					{
						Address:     "cosmos1kkyr80lkuku58h7e2v84egemscmem304mdra4f",
						VotingPower: 20,
					},
					{
						Address:     "cosmos1w3k88z6h0q5nsylh68xxdjh909qch82dlp3g5e",
						VotingPower: 20,
					},
				}

				var pk1 secp256k1.PubKey
				pk1.Key = []byte{3, 35, 137, 5, 237, 152, 146, 214, 9, 136, 235, 5, 3, 177, 22, 250, 12, 125, 38, 196, 108, 145, 211, 192, 1, 55, 245, 134, 223, 206, 52, 19, 14}
				pubKey1, err := prototypes.NewAnyWithValue(&pk1)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), sdk.MustAccAddressFromBech32(executors[0].Address)).
					Return(&authtypes.BaseAccount{
						Address:       executors[0].Address,
						PubKey:        pubKey1,
						AccountNumber: uint64(0),
					}).Times(2)

				var pk2 secp256k1.PubKey
				pk2.Key = []byte{3, 251, 110, 231, 205, 75, 69, 84, 2, 243, 138, 31, 184, 101, 131, 102, 110, 80, 73, 82, 234, 132, 253, 101, 117, 228, 30, 163, 41, 134, 159, 196, 72}
				pubKey2, err := prototypes.NewAnyWithValue(&pk2)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), sdk.MustAccAddressFromBech32(executors[1].Address)).
					Return(&authtypes.BaseAccount{
						Address:       executors[1].Address,
						PubKey:        pubKey2,
						AccountNumber: uint64(1),
					}).Times(2)

				address, err := s.keeper.GetMultisigAddress(executors)
				s.Require().NoError(err)
				s.repo.EXPECT().
					SetCommittee(gomock.Any(), uint32(11), types.Committee{
						Epoch:       11,
						Executors:   executors,
						IsEmergency: true,
						Seed:        nil,
						Address:     address,
					}).
					Return(nil)

				s.repo.EXPECT().
					SetLastEpoch(gomock.Any(), uint32(11)).
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
					addresses = append(addresses, sdk.AccAddress(pk.Bytes()))

					any, err := prototypes.NewAnyWithValue(&pk)
					s.Require().NoError(err)

					anyPubKeys = append(anyPubKeys, any)
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

			address, err := s.keeper.GetMultisigAddress(executors)
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

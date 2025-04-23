package keeper_test

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/dittonetwork/kepler/api/kepler/committee"
	"github.com/dittonetwork/kepler/x/committee/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func (s *TestSuite) TestQueryCommittee() {
	testCases := []struct {
		name          string
		setup         func()
		expectedError error
	}{
		{
			name: "alice executor",
			setup: func() {
				s.repo.EXPECT().
					GetCommittee(gomock.Any(), gomock.Any()).
					Return(types.Committee{
						Executors: []types.Executor{
							{
								Address:     s.alice.Address.String(),
								VotingPower: 1,
							},
						},
					}, nil)

				pubKeyAlice, err := codectypes.NewAnyWithValue(s.alice.PubKey)
				s.Require().NoError(err)
				s.accountKeeper.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Return(&authtypes.BaseAccount{
						Address:       s.alice.Address.String(),
						PubKey:        pubKeyAlice,
						AccountNumber: uint64(0),
					})
			},
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()
			tc.setup()
			committee, err := s.queryClient.Committee(s.ctx, &committee.QueryCommitteeRequest{Epoch: 1})
			require.NoError(s.T(), err)

			pubKey, err := codectypes.NewAnyWithValue(s.alice.PubKey)
			require.NoError(s.T(), err)

			require.Equal(s.T(), committee.Committee.Executors[0].Pubkey.Value, pubKey.Value)
		})
	}
}

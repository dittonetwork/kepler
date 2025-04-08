package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	api "github.com/dittonetwork/kepler/api/kepler/executors"
	"github.com/dittonetwork/kepler/x/executors/types"
	restaking "github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/stretchr/testify/require"
)

func (s *TestSuite) TestQueryGetEmergencyExecutors_Success() {
	valAddr, err := sdk.ValAddressFromBech32("cosmosvaloper1w7f3xx7e75p4l7qdym5msqem9rd4dyc4mq79dm")
	require.NoError(s.T(), err)

	// Prepare two executors: one active and one inactive.
	one := types.Executor{
		Address:      "cosmos1address1",
		OwnerAddress: valAddr.String(),
		IsActive:     true,
	}
	two := types.Executor{
		Address:      "cosmos1address2",
		OwnerAddress: "cosmos1address2",
		IsActive:     true,
	}

	err = s.keeper.Executors.Set(s.ctx, one.Address, one)
	require.NoError(s.T(), err)
	err = s.keeper.Executors.Set(s.ctx, two.Address, two)
	require.NoError(s.T(), err)

	s.restakingKeeper.EXPECT().GetActiveEmergencyValidators(s.ctx).Return([]restaking.Validator{
		{
			OperatorAddress: valAddr.String(),
		},
	}, nil)

	resp, err := s.queryClient.GetEmergencyExecutors(s.ctx, &api.QueryEmergencyExecutorsRequest{})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), resp)

	// Only the emergency executor should be returned.
	require.Len(s.T(), resp.Executors, 1)
	require.Equal(s.T(), "cosmosvaloper1w7f3xx7e75p4l7qdym5msqem9rd4dyc4mq79dm", resp.Executors[0].OwnerAddress)
}

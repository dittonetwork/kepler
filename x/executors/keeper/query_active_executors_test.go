package keeper_test

import (
	"github.com/dittonetwork/kepler/api/kepler/executors"
	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/require"
)

func (s *TestSuite) TestQueryGetActiveExecutors_Success() {
	// Prepare two executors: one active and one inactive.
	activeExec := types.Executor{
		Address:  "cosmos1active",
		IsActive: true,
	}
	inactiveExec := types.Executor{
		Address:  "cosmos1inactive",
		IsActive: false,
	}

	err := s.keeper.Executors.Set(s.ctx, activeExec.Address, activeExec)
	require.NoError(s.T(), err)
	err = s.keeper.Executors.Set(s.ctx, inactiveExec.Address, inactiveExec)
	require.NoError(s.T(), err)

	resp, err := s.queryClient.GetActiveExecutors(s.ctx, &executors.QueryActiveExecutorsRequest{})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), resp)
	// Only the active executor should be returned.
	require.Len(s.T(), resp.Executors, 1)
	require.Equal(s.T(), "cosmos1active", resp.Executors[0].Address)
}

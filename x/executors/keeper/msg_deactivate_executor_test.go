package keeper_test

import (
	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/require"
)

func (s *TestSuite) TestMsgDeactivateExecutor_Success() {
	resp, err := s.msgServer.AddExecutor(s.ctx, &types.MsgAddExecutor{
		Creator:      "cosmos1address1",
		OwnerAddress: "cosmos1owner",
		PublicKey:    "cosmos1publickey",
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), resp)

	deactivateResp, err := s.msgServer.DeactivateExecutor(s.ctx, &types.MsgDeactivateExecutor{
		Creator: "cosmos1address1",
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), deactivateResp)

	executor, err := s.keeper.Executors.Get(s.ctx, "cosmos1address1")
	require.NoError(s.T(), err)
	require.NotNil(s.T(), executor)
	require.Equal(s.T(), false, executor.IsActive)
}

func (s *TestSuite) TestMsgDeactivateExecutor_NotFound() {
	deactivateResp, err := s.msgServer.DeactivateExecutor(s.ctx, &types.MsgDeactivateExecutor{
		Creator: "cosmos1address2",
	})

	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "executor not found")
	require.Nil(s.T(), deactivateResp)
}

package keeper_test

import (
	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/require"
)

func (s *TestSuite) TestMsgActivateExecutor_Success() {
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

	activateResp, err := s.msgServer.ActivateExecutor(s.ctx, &types.MsgActivateExecutor{
		Creator: "cosmos1address1",
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), activateResp)

	executor, err := s.keeper.Executors.Get(s.ctx, "cosmos1address1")
	require.NoError(s.T(), err)
	require.NotNil(s.T(), executor)
	require.Equal(s.T(), true, executor.IsActive)
}

func (s *TestSuite) TestMsgActivateExecutor_OwnerHasActive() {
	resp, err := s.msgServer.AddExecutor(s.ctx, &types.MsgAddExecutor{
		Creator:      "cosmos1address1",
		OwnerAddress: "cosmos1owner",
		PublicKey:    "cosmos1publickey",
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), resp)

	resp, err = s.msgServer.AddExecutor(s.ctx, &types.MsgAddExecutor{
		Creator:      "cosmos1address2",
		OwnerAddress: "cosmos1owner",
		PublicKey:    "cosmos1publickey",
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), resp)

	// additional check if 2nd inserted as not active
	executor, err := s.keeper.Executors.Get(s.ctx, "cosmos1address2")
	require.NoError(s.T(), err)
	require.NotNil(s.T(), executor)
	require.Equal(s.T(), false, executor.IsActive)

	// try to activate him
	activateResp, err := s.msgServer.ActivateExecutor(s.ctx, &types.MsgActivateExecutor{
		Creator: "cosmos1address2",
	})

	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), types.ErrAlreadyHasActiveExecutor.Error())
	require.Nil(s.T(), activateResp)
}

func (s *TestSuite) TestMsgActivateExecutor_NotFound() {
	activateResp, err := s.msgServer.ActivateExecutor(s.ctx, &types.MsgActivateExecutor{
		Creator: "cosmos1address2",
	})

	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "executor not found")
	require.Nil(s.T(), activateResp)
}

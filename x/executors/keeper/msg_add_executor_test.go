package keeper_test

import (
	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/require"
)

func (s *TestSuite) TestMsgAddExecutor_Success() {
	resp, err := s.msgServer.AddExecutor(s.ctx, &types.MsgAddExecutor{
		Creator:      "cosmos1address1",
		OwnerAddress: "cosmos1owner",
		PublicKey:    "cosmos1publickey",
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), resp)
	require.Equal(s.T(), "cosmos1address1", resp.GetExecutor().GetAddress())
	require.Equal(s.T(), "cosmos1owner", resp.GetExecutor().GetOwnerAddress())
	require.Equal(s.T(), false, resp.GetExecutor().GetIsActive())
	require.Equal(s.T(), "cosmos1publickey", resp.GetExecutor().GetPublicKey())
	require.NotZero(s.T(), resp.GetExecutor().GetCreatedAt())
}

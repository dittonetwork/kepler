package keeper_test

import (
	"testing"

	testKeeper "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/x/executors/keeper"
	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/require"
)

func TestMsgAddExecutor_Success(t *testing.T) {
	k, ctx := testKeeper.ExecutorsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	// first executor for owner must be inserted as active
	req := &types.MsgAddExecutor{
		Creator:      "cosmos1address1",
		OwnerAddress: "cosmos1owner",
		PublicKey:    "cosmos1publickey",
	}

	resp, err := msgServer.AddExecutor(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "cosmos1address1", resp.GetExecutor().GetAddress())
	require.Equal(t, "cosmos1owner", resp.GetExecutor().GetOwnerAddress())
	require.Equal(t, true, resp.GetExecutor().GetIsActive())
	require.Equal(t, "cosmos1publickey", resp.GetExecutor().GetPublicKey())
	require.NotZero(t, resp.GetExecutor().GetCreatedAt())

	// second executor for owner must be inserted as inactive, because there is already an active executor
	req = &types.MsgAddExecutor{
		Creator:      "cosmos1address2",
		OwnerAddress: "cosmos1owner",
		PublicKey:    "cosmos1publickey",
	}

	resp, err = msgServer.AddExecutor(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "cosmos1address2", resp.GetExecutor().GetAddress())
	require.Equal(t, "cosmos1owner", resp.GetExecutor().GetOwnerAddress())
	require.Equal(t, false, resp.GetExecutor().GetIsActive())
	require.Equal(t, "cosmos1publickey", resp.GetExecutor().GetPublicKey())
	require.NotZero(t, resp.GetExecutor().GetCreatedAt())
}

package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/x/beacon/keeper"
	"kepler/x/beacon/types"
)

func TestFinalizedBlockInfoMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BeaconKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	expected := &types.MsgCreateFinalizedBlockInfo{Creator: creator}
	_, err := srv.CreateFinalizedBlockInfo(ctx, expected)
	require.NoError(t, err)
	rst, found := k.GetFinalizedBlockInfo(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestFinalizedBlockInfoMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateFinalizedBlockInfo
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateFinalizedBlockInfo{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateFinalizedBlockInfo{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BeaconKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateFinalizedBlockInfo{Creator: creator}
			_, err := srv.CreateFinalizedBlockInfo(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateFinalizedBlockInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetFinalizedBlockInfo(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestFinalizedBlockInfoMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteFinalizedBlockInfo
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteFinalizedBlockInfo{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteFinalizedBlockInfo{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BeaconKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateFinalizedBlockInfo(ctx, &types.MsgCreateFinalizedBlockInfo{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteFinalizedBlockInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetFinalizedBlockInfo(ctx)
				require.False(t, found)
			}
		})
	}
}

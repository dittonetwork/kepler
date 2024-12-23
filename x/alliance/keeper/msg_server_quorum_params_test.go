package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/x/alliance/keeper"
	"kepler/x/alliance/types"
)

func TestQuorumParamsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AllianceKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	expected := &types.MsgCreateQuorumParams{Creator: creator}
	_, err := srv.CreateQuorumParams(ctx, expected)
	require.NoError(t, err)
	rst, found := k.GetQuorumParams(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestQuorumParamsMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateQuorumParams
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateQuorumParams{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateQuorumParams{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AllianceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateQuorumParams{Creator: creator}
			_, err := srv.CreateQuorumParams(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateQuorumParams(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetQuorumParams(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestQuorumParamsMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteQuorumParams
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteQuorumParams{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteQuorumParams{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AllianceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateQuorumParams(ctx, &types.MsgCreateQuorumParams{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteQuorumParams(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetQuorumParams(ctx)
				require.False(t, found)
			}
		})
	}
}

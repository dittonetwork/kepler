package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/x/alliance/keeper"
	"kepler/x/alliance/types"
)

func TestSharedEntropyMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AllianceKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	expected := &types.MsgCreateSharedEntropy{Creator: creator}
	_, err := srv.CreateSharedEntropy(ctx, expected)
	require.NoError(t, err)
	rst, found := k.GetSharedEntropy(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestSharedEntropyMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateSharedEntropy
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateSharedEntropy{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateSharedEntropy{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AllianceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateSharedEntropy{Creator: creator}
			_, err := srv.CreateSharedEntropy(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateSharedEntropy(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetSharedEntropy(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestSharedEntropyMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteSharedEntropy
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteSharedEntropy{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteSharedEntropy{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AllianceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateSharedEntropy(ctx, &types.MsgCreateSharedEntropy{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteSharedEntropy(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetSharedEntropy(ctx)
				require.False(t, found)
			}
		})
	}
}

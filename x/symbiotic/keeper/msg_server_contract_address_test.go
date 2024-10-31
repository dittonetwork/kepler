package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/x/symbiotic/keeper"
	"kepler/x/symbiotic/types"
)

func TestContractAddressMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.SymbioticKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	expected := &types.MsgCreateContractAddress{Creator: creator}
	_, err := srv.CreateContractAddress(ctx, expected)
	require.NoError(t, err)
	rst, found := k.GetContractAddress(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestContractAddressMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateContractAddress
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateContractAddress{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateContractAddress{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SymbioticKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateContractAddress{Creator: creator}
			_, err := srv.CreateContractAddress(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateContractAddress(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetContractAddress(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestContractAddressMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteContractAddress
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteContractAddress{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteContractAddress{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SymbioticKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateContractAddress(ctx, &types.MsgCreateContractAddress{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteContractAddress(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetContractAddress(ctx)
				require.False(t, found)
			}
		})
	}
}

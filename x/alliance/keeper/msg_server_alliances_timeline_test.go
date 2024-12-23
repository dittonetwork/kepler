package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"kepler/x/alliance/types"
)

func TestAlliancesTimelineMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAlliancesTimeline(wctx, &types.MsgCreateAlliancesTimeline{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestAlliancesTimelineMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateAlliancesTimeline
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAlliancesTimeline{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAlliancesTimeline{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAlliancesTimeline{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateAlliancesTimeline(wctx, &types.MsgCreateAlliancesTimeline{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAlliancesTimeline(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAlliancesTimelineMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteAlliancesTimeline
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAlliancesTimeline{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAlliancesTimeline{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAlliancesTimeline{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateAlliancesTimeline(wctx, &types.MsgCreateAlliancesTimeline{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAlliancesTimeline(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

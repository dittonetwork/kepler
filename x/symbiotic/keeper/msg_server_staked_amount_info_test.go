package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "kepler/testutil/keeper"
	"kepler/x/symbiotic/keeper"
	"kepler/x/symbiotic/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStakedAmountInfoMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.SymbioticKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateStakedAmountInfo{Creator: creator,
			EthereumAddress: strconv.Itoa(i),
		}
		_, err := srv.CreateStakedAmountInfo(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetStakedAmountInfo(ctx,
			expected.EthereumAddress,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestStakedAmountInfoMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateStakedAmountInfo
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateStakedAmountInfo{Creator: creator,
				EthereumAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateStakedAmountInfo{Creator: "B",
				EthereumAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateStakedAmountInfo{Creator: creator,
				EthereumAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SymbioticKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateStakedAmountInfo{Creator: creator,
				EthereumAddress: strconv.Itoa(0),
			}
			_, err := srv.CreateStakedAmountInfo(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateStakedAmountInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetStakedAmountInfo(ctx,
					expected.EthereumAddress,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestStakedAmountInfoMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteStakedAmountInfo
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteStakedAmountInfo{Creator: creator,
				EthereumAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteStakedAmountInfo{Creator: "B",
				EthereumAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteStakedAmountInfo{Creator: creator,
				EthereumAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SymbioticKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateStakedAmountInfo(ctx, &types.MsgCreateStakedAmountInfo{Creator: creator,
				EthereumAddress: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteStakedAmountInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetStakedAmountInfo(ctx,
					tc.request.EthereumAddress,
				)
				require.False(t, found)
			}
		})
	}
}

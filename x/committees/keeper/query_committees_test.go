package keeper_test

import (
	"context"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"kepler/testutil/nullify"
	"kepler/x/committees/keeper"
	"kepler/x/committees/types"
)

func createNCommittees(keeper keeper.Keeper, ctx context.Context, n int) []types.Committees {
	items := make([]types.Committees, n)
	for i := range items {
		iu := uint64(i)
		items[i].Id = iu
		_ = keeper.Committees.Set(ctx, iu, items[i])
		_ = keeper.CommitteesSeq.Set(ctx, iu)
	}
	return items
}

func TestCommitteesQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNCommittees(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetCommitteesRequest
		response *types.QueryGetCommitteesResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCommitteesRequest{Id: msgs[0].Id},
			response: &types.QueryGetCommitteesResponse{Committees: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCommitteesRequest{Id: msgs[1].Id},
			response: &types.QueryGetCommitteesResponse{Committees: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCommitteesRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetCommittees(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestCommitteesQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNCommittees(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCommitteesRequest {
		return &types.QueryAllCommitteesRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListCommittees(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Committees), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Committees),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListCommittees(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Committees), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Committees),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListCommittees(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Committees),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListCommittees(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

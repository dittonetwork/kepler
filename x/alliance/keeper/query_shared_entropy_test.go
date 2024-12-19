package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "kepler/testutil/keeper"
	"kepler/testutil/nullify"
	"kepler/x/alliance/types"
)

func TestSharedEntropyQuery(t *testing.T) {
	keeper, ctx := keepertest.AllianceKeeper(t)
	item := createTestSharedEntropy(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetSharedEntropyRequest
		response *types.QueryGetSharedEntropyResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetSharedEntropyRequest{},
			response: &types.QueryGetSharedEntropyResponse{SharedEntropy: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SharedEntropy(ctx, tc.request)
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

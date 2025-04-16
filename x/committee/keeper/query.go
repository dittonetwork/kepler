package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{Keeper: k}
}

func (k queryServer) Committee(
	ctx context.Context,
	req *types.QueryCommitteeRequest,
) (*types.QueryCommitteeResponse, error) {
	committee, err := k.repository.GetCommittee(sdk.UnwrapSDKContext(ctx), req.Epoch)
	if err != nil {
		return nil, err
	}

	return &types.QueryCommitteeResponse{
		Committee: &committee,
	}, nil
}

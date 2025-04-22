package keeper

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
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

	for i, executor := range committee.Executors {
		accAddress, err := sdk.AccAddressFromBech32(executor.GetAddress())
		if err != nil {
			return nil, err
		}
		account := k.account.GetAccount(sdk.UnwrapSDKContext(ctx), accAddress)
		pubkey := account.GetPubKey()
		anyPubkey, err := codectypes.NewAnyWithValue(pubkey)
		if err != nil {
			return nil, err
		}
		committee.Executors[i].Pubkey = anyPubkey
	}
	return &types.QueryCommitteeResponse{
		Committee: &committee,
	}, nil
}

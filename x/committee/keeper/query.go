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
		accAddress, addrErr := sdk.AccAddressFromBech32(executor.GetAddress())
		if addrErr != nil {
			return nil, addrErr
		}
		account := k.account.GetAccount(sdk.UnwrapSDKContext(ctx), accAddress)
		pubkey := account.GetPubKey()
		anyPubkey, anyErr := codectypes.NewAnyWithValue(pubkey)
		if anyErr != nil {
			return nil, anyErr
		}
		committee.Executors[i].Pubkey = anyPubkey
	}
	return &types.QueryCommitteeResponse{
		Committee: &committee,
	}, nil
}

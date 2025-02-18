package keeper

import (
	"context"
	"fmt"
	"math"

	"kepler/x/committee/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CanBeSigned(
	goCtx context.Context,
	req *types.QueryCanBeSignedRequest,
) (*types.QueryCanBeSignedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	addresses := make([]string, 0, len(req.Signatures))
	for _, signature := range req.Signatures {
		address, err := getAddressFromSignature(req.JobPayload, signature)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		addresses = append(addresses, address)
	}

	commID, err := k.Committees.Indexes.Active.MatchExact(ctx, req.ChainId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	committee, err := k.Committees.Get(ctx, commID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	membersAddresses := make(map[string]bool)
	for _, member := range committee.Members {
		membersAddresses[member.Address] = true
	}

	for _, address := range addresses {
		if _, ok := membersAddresses[address]; !ok {
			return nil, status.Error(codes.Internal, "address is not a member of the committee")
		}
	}

	if !isSuperMajority(len(committee.Members), len(addresses)) {
		return nil, status.Error(codes.Internal, "not enough votes")
	}

	return &types.QueryCanBeSignedResponse{
		CanBeSigned: true,
	}, nil
}

func getAddressFromSignature(jobMsg []byte, signature []byte) (string, error) {
	hsh := crypto.Keccak256(jobMsg)
	pubKey, err := crypto.SigToPub(hsh, signature)
	if err != nil {
		return "", fmt.Errorf("failed to recover public key from signature: %w", err)
	}

	return crypto.PubkeyToAddress(*pubKey).Hex(), nil
}

func isSuperMajority(totalVotes, votesFor int) bool {
	if votesFor > totalVotes || totalVotes == 0 {
		return false
	}

	//nolint:mnd // just a formula
	requiredVotes := math.Ceil(float64(totalVotes) * 2 / 3)
	return votesFor >= int(requiredVotes)
}

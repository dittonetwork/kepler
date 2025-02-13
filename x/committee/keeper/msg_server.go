package keeper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"kepler/x/committee/types"

	"github.com/cometbft/cometbft/crypto"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

var ErrInvalidCommitmentHash = errors.New("randao: invalid commitment hash")
var ErrCommitmentAlreadyExists = errors.New("randao: commitment on this epoch already written")

// CommitRandao implements types.MsgServer.
func (k msgServer) CommitRandao(
	ctx context.Context,
	msg *types.MsgCommitRandao,
) (*types.MsgCommitRandaoResponse, error) {
	// TODO: For CommitRandao, the steps are: validate the validator's status,
	// check if the commit phase is active, verify the commitment isn't a duplicate, and store the commitment.
	valAddr, err := k.validatorAddressCodec.StringToBytes(msg.Validator)
	if err != nil {
		return nil, err
	}

	commitment := types.CommitRandao{
		CommitmentHash: msg.CommitmentHash,
	}

	err = k.Keeper.setCommitment(ctx, msg.ExecutionChainId, msg.EpochId, valAddr, commitment)
	if err != nil {
		return nil, err
	}
	return &types.MsgCommitRandaoResponse{}, nil
}

// RevealRandao implements types.MsgServer.
func (k msgServer) RevealRandao(
	ctx context.Context,
	msg *types.MsgRevealRandao,
) (*types.MsgRevealRandaoResponse, error) {
	valAddr, err := k.validatorAddressCodec.StringToBytes(msg.Validator)
	if err != nil {
		return nil, err
	}

	randaoCommitmentEpoch := msg.EpochId - 1
	commitment, err := k.Keeper.GetCommitment(
		ctx,
		msg.ExecutionChainId,
		randaoCommitmentEpoch,
		valAddr,
	)
	if err != nil {
		return nil, fmt.Errorf("commitment not found: %w", err)
	}

	// Verify the hash matches the revealed seed
	hashedReveal := crypto.Sha256(msg.RandomSeed)
	if !bytes.Equal(hashedReveal, commitment.CommitmentHash) {
		return nil, ErrInvalidCommitmentHash
	}

	reveal := types.RevealRandao{
		RandomSeed: msg.RandomSeed,
	}
	err = k.Keeper.setReveal(ctx, msg.ExecutionChainId, msg.EpochId, valAddr, reveal)
	if err != nil {
		return nil, err
	}
	return &types.MsgRevealRandaoResponse{}, nil
}

package keeper

import (
	"context"
	"kepler/x/committee/types"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetCommitment(
	ctx context.Context,
	chainID string,
	epochID uint64,
	valAddr sdk.ValAddress,
) (types.CommitRandao, error) {
	return k.RandaoCommitments.Get(ctx, collections.Join3(chainID, epochID, valAddr))
}

func (k Keeper) HasCommitment(
	ctx context.Context,
	chainID string,
	epochID uint64,
	valAddr sdk.ValAddress,
) (bool, error) {
	return k.RandaoCommitments.Has(ctx, collections.Join3(chainID, epochID, valAddr))
}

func (k Keeper) setCommitment(
	ctx context.Context,
	chainID string,
	epochID uint64,
	valAddr sdk.ValAddress,
	data types.CommitRandao,
) error {
	hasCommitment, err := k.HasCommitment(ctx, chainID, epochID, valAddr)
	if err != nil {
		return err
	}
	if hasCommitment {
		return ErrCommitmentAlreadyExists
	}

	return k.RandaoCommitments.Set(ctx, collections.Join3(chainID, epochID, valAddr), data)
}

func (k Keeper) GetReveal(
	ctx context.Context,
	chainID string,
	epochID uint64,
	valAddr sdk.ValAddress,
) (types.RevealRandao, error) {
	return k.RandaoReveals.Get(ctx, collections.Join3(chainID, epochID, valAddr))
}

func (k Keeper) setReveal(
	ctx context.Context,
	chainID string,
	epochID uint64,
	valAddr sdk.ValAddress,
	data types.RevealRandao,
) error {
	return k.RandaoReveals.Set(ctx, collections.Join3(chainID, epochID, valAddr), data)
}

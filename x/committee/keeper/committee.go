package keeper

import (
	"context"
	"math"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	sdksecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

// CreateCommittee creates a new committee by the given epoch.
func (k Keeper) CreateCommittee(ctx sdk.Context, epoch uint32) (types.Committee, error) {
	var committee types.Committee

	ok, err := k.repository.HasCommittee(ctx, epoch)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to check if committee exists")
	}

	// check if the committee already exists
	if ok {
		return types.Committee{}, types.ErrCommitteeAlreadyExists
	}

	var lastSavedEpoch uint32
	lastSavedEpoch, err = k.repository.GetLastEpoch(ctx)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to get last saved epoch")
	}

	// check if the new epoch is greater than the last saved epoch
	if epoch <= lastSavedEpoch {
		return types.Committee{}, types.ErrInvalidEpoch
	}

	committee, err = k.createEmergencyCommittee(ctx, epoch)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to create emergency committee")
	}
	committee.Address, err = k.GetMultisigAddress(committee.Executors)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to get multisig address")
	}

	err = k.repository.SetCommittee(ctx, epoch, committee)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to set committee")
	}
	err = k.repository.SetLastEpoch(ctx, epoch)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to set last epoch")
	}

	k.Logger(ctx).With("committee", committee).Info("committee created")

	return committee, nil
}

func (k Keeper) GetMultisigAddress(executors []types.Executor) (string, error) {
	pubKeys := make([]cryptotypes.PubKey, 0, len(executors))
	for _, each := range executors {
		addr, err := sdk.AccAddressFromBech32(each.GetAddress())
		if err != nil {
			return "", err
		}
		acc := k.account.GetAccount(context.Background(), addr)

		var pk sdksecp.PubKey
		pk.Key = make([]byte, len(acc.GetPubKey().Bytes()))
		copy(pk.Key, acc.GetPubKey().Bytes())

		pubKeys = append(pubKeys, &pk)
	}

	// minimum votes for BFT consensus
	// https://pmg.csail.mit.edu/papers/osdi99.pdf
	threshold := int(math.Floor(2*(float64(len(executors))-1)/3) + 1) //nolint:mnd // correct formula

	multiPubKey := multisig.NewLegacyAminoPubKey(threshold, pubKeys)

	return sdk.AccAddress(multiPubKey.Address()).String(), nil
}

// createEmergencyCommittee creates an emergency committee by the given epoch.
func (k Keeper) createEmergencyCommittee(ctx sdk.Context, epoch uint32) (types.Committee, error) {
	emergencyValidators, err := k.restaking.GetActiveEmergencyValidators(ctx)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to get emergency executors")
	}

	committeeExecutors := make([]types.Executor, len(emergencyValidators))
	for i, validator := range emergencyValidators {
		committeeExecutors[i] = types.Executor{
			Address:     validator.OperatorAddress,
			VotingPower: validator.VotingPower,
		}
	}

	return types.Committee{
		IsEmergency: true,
		Epoch:       epoch,
		Seed:        ctx.HeaderInfo().Hash,
		Executors:   committeeExecutors,
	}, nil
}

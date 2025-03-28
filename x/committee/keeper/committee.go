package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

// CreateCommittee creates a new committee by the given epoch.
func (k Keeper) CreateCommittee(ctx sdk.Context, epoch uint32) (types.Committee, error) {
	var committee types.Committee

	ok, err := k.Committees.Has(ctx, epoch)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to check if committee exists")
	}

	// check if the committee already exists
	if ok {
		return types.Committee{}, types.ErrCommitteeAlreadyExists
	}

	var latestEpoch uint32
	latestEpoch, err = k.LatestEpoch.Get(ctx)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to get latest epoch")
	}

	// check if the latest epoch is less than the given epoch
	if latestEpoch <= epoch {
		return types.Committee{}, types.ErrInvalidEpoch
	}

	committee, err = k.createEmergencyCommittee(ctx, epoch)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to create emergency committee")
	}

	err = k.Committees.Set(ctx, epoch, committee)
	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to set committee")
	}

	k.Logger().With("committee", committee).Info("committee created")

	return committee, nil
}

// createEmergencyCommittee creates an emergency committee by the given epoch.
func (k Keeper) createEmergencyCommittee(ctx sdk.Context, epoch uint32) (types.Committee, error) {
	executors, err := k.executors.GetEmergencyExecutors(ctx)

	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to get emergency executors")
	}

	committeeExecutors := make([]types.Executor, len(executors))

	for i, executor := range executors {
		committeeExecutors[i] = types.Executor{
			Address:     executor.GetAddress().String(),
			VotingPower: executor.GetVotingPower(),
		}
	}

	return types.Committee{
		IsEmergency: true,
		Epoch:       epoch,
		Seed:        ctx.HeaderInfo().Hash,
		Executors:   committeeExecutors,
	}, nil
}

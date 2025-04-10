package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
	restaking "github.com/dittonetwork/kepler/x/restaking/types"
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

// createEmergencyCommittee creates an emergency committee by the given epoch.
func (k Keeper) createEmergencyCommittee(ctx sdk.Context, epoch uint32) (types.Committee, error) {
	executors, err := k.executors.GetEmergencyExecutors(ctx)

	if err != nil {
		return types.Committee{}, sdkerrors.Wrap(err, "failed to get emergency executors")
	}

	committeeExecutors := make([]types.Executor, len(executors))

	for i, executor := range executors {
		var validator restaking.Validator
		validator, err = k.restaking.GetValidator(ctx, sdk.ValAddress(executor.GetOwnerAddress()))

		if err != nil {
			return types.Committee{}, sdkerrors.Wrap(err, "failed to get validator")
		}

		committeeExecutors[i] = types.Executor{
			Address:     executor.GetAddress(),
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

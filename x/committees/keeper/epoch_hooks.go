package keeper

import (
	"context"
	"encoding/binary"
	"kepler/x/committees/types"
	"math/rand"

	epochstypes "cosmossdk.io/x/epochs/types"
	stakingtypes "cosmossdk.io/x/staking/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

var _ epochstypes.EpochHooks = (*Keeper)(nil)

// AfterEpochEnd is called when epoch is going to be ended, epochNumber is the number of epoch that is ending.
func (k Keeper) AfterEpochEnd(ctx context.Context, epochIdentifier string, epochNumber int64) error {
	return nil
}

// BeforeEpochStart is called when epoch is going to be started, epochNumber is the number of epoch that is starting.
func (k Keeper) BeforeEpochStart(ctx context.Context, epochIdentifier string, epochNumber int64) error {
	return k.FillCommitteesTimeline(ctx)
}

func (k Keeper) getAllCommittees(goCtx context.Context) ([]types.Committees, error) {
	iter, err := k.Committees.Iterate(goCtx, nil)
	if err != nil {
		return nil, err
	}

	return iter.Values()
}

// FillCommitteesTimeline feels an array of future aliances useing validators list
func (k Keeper) FillCommitteesTimeline(goCtx context.Context) error {
	committees, err := k.getAllCommittees(goCtx)
	if err != nil {
		return err
	}

	if len(committees) != 0 {
		// epoch is almost gone, remove the first commitee
		k.Logger.Info("removing the oldest commitee")
		committees = committees[1:]
	}

	sdkCtx := sdktypes.UnwrapSDKContext(goCtx)

	const COMMITTEES_IN_FUTURE = 2
	const MAX_PARTICIPANTS = 10
	seed := binary.LittleEndian.Uint64(sdkCtx.BlockHeader().AppHash)
	for i := 0; i < COMMITTEES_IN_FUTURE-len(committees); i++ {
		validators, err := k.selectValidatorsForCommittee(goCtx, uint(MAX_PARTICIPANTS), uint(seed)+uint(i))
		if err != nil {
			return err
		}
		validatorsAddresses := make([]string, len(validators))
		for j, validator := range validators {
			validatorsAddresses[j] = validator.OperatorAddress
		}
		newCommittee := types.Committees{
			Participants: validatorsAddresses,
		}
		committees = append(committees, newCommittee)
	}

	for i, committee := range committees {
		k.Committees.Set(goCtx, uint64(i), committee)
	}

	// if there were more commitees than we need, remove extra ones
	for i := COMMITTEES_IN_FUTURE; i < len(committees); i++ {
		k.Committees.Remove(goCtx, uint64(i))
	}

	return nil
}

// generateRandomIndices generates k unique random indices from 0 to min(n-1, k).
func generateRandomIndices(n uint, k uint, seed uint) []int {
	rand := rand.New(rand.NewSource(int64(seed)))

	if k > n {
		k = n
	}
	indices := rand.Perm(int(n))[:k]
	return indices
}

// selectValidatorsForCommittee returns validators from pool to form new committee
func (k Keeper) selectValidatorsForCommittee(goCtx context.Context, maxParticipants uint, seed uint) ([]stakingtypes.Validator, error) {
	allValidators, err := k.stakingKeeper.GetBondedValidators(goCtx)
	if err != nil {
		return nil, err
	}

	randomIndices := generateRandomIndices(uint(len(allValidators)), maxParticipants, seed)

	var generatedAliance []stakingtypes.Validator
	for i, randomIndex := range randomIndices {
		if uint(i) == maxParticipants {
			break
		}
		generatedAliance = append(generatedAliance, allValidators[randomIndex])
	}

	return generatedAliance, nil
}

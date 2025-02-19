package keeper

import (
	"fmt"
	"kepler/x/job/types"

	"github.com/pkg/errors"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	collectionNameJob = "job"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		committeeKeeper types.CommitteeKeeper

		// Jobs key: jobID | value: job
		// This is used to store jobs
		Jobs collections.Map[uint64, types.Job]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	committeeKeeper types.CommitteeKeeper,
) Keeper {
	sb := collections.NewSchemaBuilder(storeService)

	return Keeper{
		cdc:             cdc,
		storeService:    storeService,
		logger:          logger,
		committeeKeeper: committeeKeeper,
		Jobs: collections.NewMap(sb, types.JobsPrefix, collectionNameJob, collections.Uint64Key,
			codec.CollValue[types.Job](cdc)),
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateJob(ctx sdk.Context, job types.Job) error {
	if len(job.Signs) == 0 {
		return types.ErrJobSignsIsNil
	}
	has, err := k.Jobs.Has(ctx, job.Id)
	if err != nil {
		return fmt.Errorf("check job exists: %w", err)
	}
	if has {
		return fmt.Errorf("job with id %d already exists: %w", job.Id, types.ErrJobAlreadyExists)
	}

	committeeExists, err := k.committeeKeeper.IsCommitteeExists(ctx, job.CommitteeId)
	if err != nil {
		return fmt.Errorf("check job committee exists: %w", err)
	}
	if !committeeExists {
		return fmt.Errorf("committee_id: %s, %w ", job.CommitteeId, types.ErrCommitteeDoesntExists)
	}

	// TODO: need to check validity of signs and passed payload

	jobBytes, err := job.Marshal()
	if err != nil {
		return fmt.Errorf("marshal job: %w", err)
	}
	signsValid, err := k.committeeKeeper.CanBeSigned(ctx, job.ChainId, job.Signs, jobBytes)
	if err != nil {
		return fmt.Errorf("check job signs: %w", err)
	}

	if !signsValid {
		job.Status = types.Job_STATUS_INVALID
	}

	err = k.Jobs.Set(ctx, job.Id, job)
	if err != nil {
		return fmt.Errorf("set job: %w", err)
	}

	return nil
}

func (k Keeper) GetJobByID(ctx sdk.Context, jobID uint64) (types.Job, bool, error) {
	job, err := k.Jobs.Get(ctx, jobID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return job, false, nil
		}
		return job, false, err
	}
	return job, true, nil
}

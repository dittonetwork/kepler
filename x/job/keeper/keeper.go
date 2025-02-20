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
	collectionNameJob   = "job"
	collectionNameJobID = "job_id"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		committeeKeeper types.CommitteeKeeper

		// Jobs key: jobID | value: job
		// This is used to store jobs
		Jobs  collections.Map[uint64, types.Job]
		JobID collections.Sequence
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
		JobID: collections.NewSequence(sb, types.JobsPrefix, collectionNameJobID),
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateJob(
	ctx sdk.Context,
	status types.Job_Status,
	committeeID string,
	chainID string,
	automationID uint64,
	txHash string,
	executorAddress string,
	createdAt uint64,
	executedAt uint64,
	signedAt uint64,
	signs [][]byte,
	payload []byte,
) error {
	if len(signs) == 0 {
		return types.ErrJobSignsIsNil
	}

	jobID, err := k.JobID.Next(ctx)
	if err != nil {
		return fmt.Errorf("get next job id: %w", err)
	}

	committeeExists, err := k.committeeKeeper.IsCommitteeExists(ctx, committeeID)
	if err != nil {
		return fmt.Errorf("check job committee exists: %w", err)
	}
	if !committeeExists {
		return fmt.Errorf("committee_id: %s, %w ", committeeID, types.ErrCommitteeDoesntExists)
	}

	signsValid, err := k.committeeKeeper.CanBeSigned(ctx, committeeID, chainID, signs, payload)
	if err != nil {
		return fmt.Errorf("check job signs: %w", err)
	}

	newJobStatus := status
	if !signsValid {
		newJobStatus = types.Job_STATUS_INVALID
	}

	err = k.Jobs.Set(ctx, jobID, types.Job{
		Id:              jobID,
		Status:          newJobStatus,
		CommitteeId:     committeeID,
		ChainId:         chainID,
		AutomationId:    automationID,
		TxHash:          txHash,
		ExecutorAddress: executorAddress,
		CreatedAt:       createdAt,
		ExecutedAt:      executedAt,
		SignedAt:        signedAt,
		Signs:           signs,
	})
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

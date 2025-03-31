package keeper

import (
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/executors/types"
)

type Idx struct {
	// Emergency index by emergency status
	ExecutorsByOwnerAddress *indexes.Multi[string, string, types.Executor]
}

func NewIndexes(sb *collections.SchemaBuilder) Idx {
	return Idx{
		ExecutorsByOwnerAddress: indexes.NewMulti(
			sb,
			types.ExecutorsPrefix,
			types.CollectionIndexExecutorsByOwnerAddress,
			collections.StringKey,
			collections.StringKey,
			func(_ string, val types.Executor) (string, error) {
				return val.OwnerAddress, nil
			}),
	}
}

func (a Idx) IndexesList() []collections.Index[string, types.Executor] {
	return []collections.Index[string, types.Executor]{
		a.ExecutorsByOwnerAddress,
	}
}

// addExecutor generates a new executor ID, sets the creation timestamp,
// and stores the executor. It returns the stored executor.
func (k Keeper) addExecutor(ctx sdk.Context, executor types.Executor) (*types.Executor, error) {
	executor.CreatedAt = time.Now().UTC().Unix()
	if err := k.Executors.Set(ctx, executor.Address, executor); err != nil {
		return nil, fmt.Errorf("failed to store executor: %w", err)
	}

	return &executor, nil
}

// getAllExecutors returns a slice of all executors stored.
func (k Keeper) getAllExecutors(ctx sdk.Context) ([]types.Executor, error) {
	var executors []types.Executor
	iter, err := k.Executors.Iterate(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to iterate executors: %w", err)
	}

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var v types.Executor
		v, err = iter.Value()
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal executor: %w", err)
		}

		executors = append(executors, v)
	}

	return executors, nil
}

// GetExecutorsByOwnerAddress returns a slice of executors with the given owner address.
func (k Keeper) GetExecutorsByOwnerAddress(ctx sdk.Context, ownerAddress string) ([]types.Executor, error) {
	iter, err := k.Executors.Indexes.ExecutorsByOwnerAddress.MatchExact(
		ctx,
		ownerAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get executors by owner address: %w", err)
	}

	pks, err := iter.PrimaryKeys()
	if err != nil {
		return nil, fmt.Errorf("failed to get primary keys: %w", err)
	}

	executors := make([]types.Executor, len(pks))
	for i, pk := range pks {
		executor, inErr := k.Executors.Get(ctx, pk)
		if inErr != nil {
			return nil, fmt.Errorf("failed to get executor: %w", inErr)
		}

		executors[i] = executor
	}

	return executors, nil
}

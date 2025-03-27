package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/executors/types"
)

// AddExecutor generates a new executor ID, sets the creation timestamp,
// and stores the executor. It returns the stored executor.
func (k Keeper) AddExecutor(ctx sdk.Context, executor types.Executor) (*types.Executor, error) {
	executor.CreatedAt = time.Now().UTC().Unix()
	if err := k.Executors.Set(ctx, executor.Address, executor); err != nil {
		return nil, fmt.Errorf("failed to store executor: %w", err)
	}

	return &executor, nil
}

// GetAllExecutors returns a slice of all executors stored.
func (k Keeper) GetAllExecutors(ctx sdk.Context) ([]types.Executor, error) {
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

// SetIsActive sets isActive for executor.
func (k Keeper) SetIsActive(ctx sdk.Context, addr string, isActive bool) error {
	v, err := k.Executors.Get(ctx, addr)
	if err != nil {
		return fmt.Errorf("failed to get executor: %w", err)
	}

	v.IsActive = isActive

	if err = k.Executors.Set(ctx, addr, v); err != nil {
		return fmt.Errorf("failed to update executor activity: %w", err)
	}

	return nil
}

package keeper_test

import (
	"errors"
	"testing"

	"cosmossdk.io/collections"
	"github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/require"
)

func TestAddExecutor(t *testing.T) {
	k, ctx := keeper.ExecutorsKeeper(t)

	exec := types.Executor{
		Address:      "cosmos1testaddress",
		OwnerAddress: "0x1234567890abcdef1234567890abcdef12345678",
		PublicKey:    "cosmospub1testkey",
		IsActive:     true,
	}
	added, err := k.AddExecutor(ctx, exec)
	require.NoError(t, err)
	require.NotNil(t, added)
	require.NotZero(t, added.CreatedAt)

	// Verify the executor is stored.
	stored, err := k.Executors.Get(ctx, exec.Address)
	require.NoError(t, err)
	require.Equal(t, added, &stored)
}

func TestGetAllExecutors(t *testing.T) {
	k, ctx := keeper.ExecutorsKeeper(t)
	// Add two executors.
	exec1 := types.Executor{
		Address:      "cosmos1addr1",
		OwnerAddress: "0xaaa",
		PublicKey:    "pubkey1",
		IsActive:     true,
	}
	exec2 := types.Executor{
		Address:      "cosmos1addr2",
		OwnerAddress: "0xbbb",
		PublicKey:    "pubkey2",
		IsActive:     false,
	}
	_, err := k.AddExecutor(ctx, exec1)
	require.NoError(t, err)
	_, err = k.AddExecutor(ctx, exec2)
	require.NoError(t, err)

	all, err := k.GetAllExecutors(ctx)
	require.NoError(t, err)
	require.Len(t, all, 2)
}

func TestSetIsActive(t *testing.T) {
	k, ctx := keeper.ExecutorsKeeper(t)

	exec := types.Executor{
		Address:      "cosmos1toggle",
		OwnerAddress: "0xtoggle",
		PublicKey:    "pubkeytoggle",
		IsActive:     false,
	}
	_, err := k.AddExecutor(ctx, exec)
	require.NoError(t, err)

	// Toggle from false to true.
	err = k.SetIsActive(ctx, exec.Address, true)
	require.NoError(t, err)
	updated, err := k.Executors.Get(ctx, exec.Address)
	require.NoError(t, err)
	require.True(t, updated.IsActive)

	// Toggle back to false.
	err = k.SetIsActive(ctx, exec.Address, false)
	require.NoError(t, err)
	updated, err = k.Executors.Get(ctx, exec.Address)
	require.NoError(t, err)
	require.False(t, updated.IsActive)

	// Not found case
	err = k.SetIsActive(ctx, "cosmos1notfound", true)
	require.Error(t, err)
	require.True(t, errors.Is(err, collections.ErrNotFound))
}

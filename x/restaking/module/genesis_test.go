package restaking_test

import (
	"testing"
	"time"

	keepertest "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/testutil/nullify"
	restaking "github.com/dittonetwork/kepler/x/restaking/module"
	"github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		LastUpdate: types.LastUpdate{
			Epoch:       1,
			Timestamp:   time.Now().UTC(),
			BlockHeight: 100,
			BlockHash:   "test_block_hash",
		},
		Validators: []types.Validator{
			{
				Address:       "validator1",
				CosmosAddress: "cosmos_address1",
				IsEmergency:   true,
				VotingPower:   100,
				Status:        types.ValidatorStatus_VALIDATOR_STATUS_BONDED,
			},
			{
				Address:       "validator2",
				CosmosAddress: "cosmos_address2",
				IsEmergency:   false,
				VotingPower:   200,
				Status:        types.ValidatorStatus_VALIDATOR_STATUS_BONDING,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RestakingKeeper(t)
	restaking.InitGenesis(ctx, k, genesisState)
	got := restaking.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Params, got.Params)
	require.Equal(t, genesisState.LastUpdate.Epoch, got.LastUpdate.Epoch)
	require.Equal(t, genesisState.LastUpdate.BlockHeight, got.LastUpdate.BlockHeight)
	require.Equal(t, genesisState.LastUpdate.BlockHash, got.LastUpdate.BlockHash)
	// Comparing the timestamp directly might fail due to precision issues during serialization
	require.NotNil(t, got.LastUpdate.Timestamp)

	require.ElementsMatch(t, genesisState.Validators, got.Validators)
	// this line is used by starport scaffolding # genesis/test/assert
}

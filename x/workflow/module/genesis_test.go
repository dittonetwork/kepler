package workflow_test

import (
	"testing"

	keepertest "github.com/dittonetwork/kepler/testutil/keeper"
	"github.com/dittonetwork/kepler/testutil/nullify"
	workflow "github.com/dittonetwork/kepler/x/workflow/module"
	"github.com/dittonetwork/kepler/x/workflow/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WorkflowKeeper(t, nil)
	workflow.InitGenesis(ctx, k, genesisState)
	got := workflow.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

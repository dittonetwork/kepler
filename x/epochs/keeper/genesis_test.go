package keeper_test

import (
	"testing"

	"github.com/dittonetwork/kepler/x/epochs/types"

	"github.com/stretchr/testify/require"
)

func TestEpochsExportGenesis(t *testing.T) {
	ctx, keeper := Setup(t)

	chainStartTime := ctx.HeaderInfo().Time
	chainStartHeight := ctx.HeaderInfo().Height

	genesis, err := keeper.ExportGenesis(ctx)
	require.NoError(t, err)
	require.Len(t, genesis.Epochs, 4)

	expectedEpochs := types.DefaultGenesis().Epochs
	for i := range expectedEpochs {
		expectedEpochs[i].CurrentEpochStartHeight = chainStartHeight
		expectedEpochs[i].StartTime = chainStartTime
	}

	require.Equal(t, expectedEpochs, genesis.Epochs)
}

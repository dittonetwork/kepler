package keeper_test

import (
	"kepler/x/epochs/types"
)

func (s *KeeperTestSuite) TestQueryEpochInfos() {
	s.SetupTest()

	// Check that querying epoch infos on default genesis returns the default genesis epoch infos
	epochInfosResponse, err := s.queryClient.EpochInfos(s.Ctx, &types.QueryEpochsInfoRequest{})
	s.Require().NoError(err)
	s.Require().Len(epochInfosResponse.Epochs, 4)
	expectedEpochs := types.DefaultGenesis().Epochs
	for id := range expectedEpochs {
		expectedEpochs[id].StartTime = s.Ctx.HeaderInfo().Time
		expectedEpochs[id].CurrentEpochStartHeight = s.Ctx.HeaderInfo().Height
	}

	s.Require().Equal(expectedEpochs, epochInfosResponse.Epochs)
}

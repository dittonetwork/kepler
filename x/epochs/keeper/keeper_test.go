package keeper_test

import (
	epochskeeper "kepler/x/epochs/keeper"
	"kepler/x/epochs/types"
	"testing"
	"time"

	"cosmossdk.io/core/header"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
)

type KeeperTestSuite struct {
	suite.Suite
	Ctx          sdk.Context
	EpochsKeeper *epochskeeper.Keeper

	queryClient types.QueryClient
}

func (s *KeeperTestSuite) SetupTest() {
	ctx, keeper := Setup(s.T())

	s.Ctx = ctx
	s.EpochsKeeper = keeper

	queryRouter := baseapp.NewGRPCQueryRouter()
	cfg := module.NewConfigurator(nil, nil, queryRouter)
	types.RegisterQueryServer(cfg.QueryServer(), epochskeeper.NewQuerier(*s.EpochsKeeper))
	grpcQueryService := &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: queryRouter,
		Ctx:             s.Ctx,
	}
	encCfg := moduletestutil.MakeTestEncodingConfig()
	grpcQueryService.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	s.queryClient = types.NewQueryClient(grpcQueryService)
}

func Setup(t *testing.T) (sdk.Context, *epochskeeper.Keeper) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(t, key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithHeaderInfo(header.Info{Time: time.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()

	keeper := epochskeeper.NewKeeper(encCfg.Codec, storeService, ctx.Logger())
	keeper = keeper.SetHooks(types.NewMultiEpochHooks())

	ctx.WithHeaderInfo(header.Info{Height: 1, Time: time.Now().UTC(), ChainID: "epochs"})
	err := keeper.InitGenesis(ctx, *types.DefaultGenesis())

	require.NoError(t, err)

	SetEpochStartTime(ctx, *keeper)

	return ctx, keeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func SetEpochStartTime(ctx sdk.Context, epochsKeeper epochskeeper.Keeper) {
	epochs, err := epochsKeeper.AllEpochInfos(ctx)
	if err != nil {
		panic(err)
	}
	for _, epoch := range epochs {
		epoch.StartTime = ctx.HeaderInfo().Time
		err := epochsKeeper.EpochInfo.Remove(ctx, epoch.Identifier)
		if err != nil {
			panic(err)
		}
		err = epochsKeeper.AddEpochInfo(ctx, epoch)
		if err != nil {
			panic(err)
		}
	}
}

package keeper_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/dittonetwork/kepler/api/kepler/executors"
	"github.com/dittonetwork/kepler/x/executors/keeper"
	executorsmodule "github.com/dittonetwork/kepler/x/executors/module"
	exectestutil "github.com/dittonetwork/kepler/x/executors/testutil"
	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient executors.QueryClient
	msgServer   types.MsgServer
	keeper      keeper.Keeper

	restakingKeeper *exectestutil.MockRestakingKeeper
}

func TestKeeper(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig(executorsmodule.AppModuleBasic{})

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	restakingKeeper := exectestutil.NewMockRestakingKeeper(ctrl)

	executorsKeeper := keeper.NewKeeper(encCfg.Codec, storeService, restakingKeeper)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, keeper.NewQueryServerImpl(executorsKeeper))

	s.restakingKeeper = restakingKeeper
	s.keeper = executorsKeeper
	s.queryClient = executors.NewQueryClient(queryHelper)
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(executorsKeeper)

	s.Require().Equal(ctx.Logger().With("module", "x/"+types.ModuleName), s.keeper.Logger(ctx))
}

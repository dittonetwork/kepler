package keeper_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/dittonetwork/kepler/api/kepler/restaking"
	committeetypes "github.com/dittonetwork/kepler/x/committee/types"
	"github.com/dittonetwork/kepler/x/restaking/keeper"
	restakingmodule "github.com/dittonetwork/kepler/x/restaking/module"
	restakingtestutils "github.com/dittonetwork/kepler/x/restaking/testutil"
	"github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient restaking.QueryClient
	msgServer   types.MsgServer
	keeper      keeper.Keeper
	repository  *restakingtestutils.MockRepository
}

func TestKeeper(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig(restakingmodule.AppModuleBasic{})

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	s.repository = restakingtestutils.NewMockRepository(ctrl)
	accountKeeper := restakingtestutils.NewMockAccountKeeper(ctrl)

	accountKeeper.EXPECT().AddressCodec().Return(address.NewBech32Codec("cosmos")).Times(1)

	restakingKeeper := keeper.NewKeeper(
		encCfg.Codec,
		ctx.Logger(),
		s.repository,
		accountKeeper,
		authtypes.NewModuleAddress(committeetypes.ModuleName).String(),
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, keeper.NewQueryServerImpl(*restakingKeeper))

	s.keeper = *restakingKeeper
	s.queryClient = restaking.NewQueryClient(queryHelper)
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(*restakingKeeper)

	s.Require().Equal(s.keeper.Logger().With("module", "x/"+types.ModuleName), s.keeper.Logger())
}

package keeper_test

import (
	"fmt"
	"testing"

	addresscodec "cosmossdk.io/core/address"
	storetypes "cosmossdk.io/store/types"
	"github.com/cometbft/cometbft/crypto/sr25519"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/dittonetwork/kepler/api/kepler/committee"
	"github.com/dittonetwork/kepler/testutil/sample"
	"github.com/dittonetwork/kepler/x/committee/keeper"
	committeemodule "github.com/dittonetwork/kepler/x/committee/module"
	committeetestutil "github.com/dittonetwork/kepler/x/committee/testutil"
	"github.com/dittonetwork/kepler/x/committee/types"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestSuite struct {
	suite.Suite

	ctx             sdk.Context
	queryClient     committee.QueryClient
	msgServer       types.MsgServer
	keeper          keeper.Keeper
	authority       string
	valAddressCodec addresscodec.Codec

	bankKeeper      *committeetestutil.MockBankKeeper
	accountKeeper   *committeetestutil.MockAccountKeeper
	restakingKeeper *committeetestutil.MockRestakingKeeper
	repo            *committeetestutil.MockRepository

	alice sample.Account
	bob   sample.Account
}

func TestNewKeeper(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(
		s.T(), key, storetypes.NewTransientStoreKey("transient_test"),
	)
	ctx := testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig(committeemodule.AppModuleBasic{})

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	accountKeeper := committeetestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := committeetestutil.NewMockBankKeeper(ctrl)
	restakingKeeper := committeetestutil.NewMockRestakingKeeper(ctrl)

	pubKey := sr25519.GenPrivKey().PubKey()

	repo := committeetestutil.NewMockRepository(ctrl)

	// Set sample accounts
	s.alice = sample.GetAccount("alice")
	s.bob = sample.GetAccount("bob")

	// Generate Bech32 encoded address
	s.authority = sdk.MustBech32ifyAddressBytes(
		sdk.GetConfig().GetBech32AccountAddrPrefix(), pubKey.Address(),
	)
	s.valAddressCodec = address.NewBech32Codec("cosmosvaloper")

	committeeKeeper := keeper.NewKeeper(
		s.authority,
		accountKeeper,
		bankKeeper,
		restakingKeeper,
		repo,
		ctx.Logger(),
		nil,
		encCfg.Amino,
		encCfg.Codec,
		"hour",
		s.valAddressCodec,
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, keeper.NewQueryServerImpl(committeeKeeper))

	s.accountKeeper = accountKeeper
	s.bankKeeper = bankKeeper
	s.restakingKeeper = restakingKeeper
	s.keeper = committeeKeeper
	s.queryClient = committee.NewQueryClient(queryHelper)
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(committeeKeeper)
	s.repo = repo

	s.Require().Equal(
		ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName)),
		s.keeper.Logger(),
	)

	// workaround
	// @TODO link or create issue
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
	).Return(nil).AnyTimes()
}

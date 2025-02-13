package keeper_test

import (
	"kepler/x/committee/keeper"
	"kepler/x/committee/types"
	"testing"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	codectestutils "github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	testkeeper "kepler/testutil/keeper"

	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx     sdk.Context
	keeper  keeper.Keeper
	msgSrvr types.MsgServer
}

func (s *KeeperTestSuite) SetupTest() {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	s.ctx = testutil.DefaultContext(
		storeKey,
		storetypes.NewTransientStoreKey("transient_"+types.StoreKey),
	)
	interfaceRegistry := codectestutils.CodecOptions{}.NewInterfaceRegistry()
	c := codec.NewProtoCodec(interfaceRegistry)

	authtype := authtypes.NewModuleAddress(types.ModuleName)
	ek, _ := testkeeper.EpochsKeeper(s.T())
	s.keeper = keeper.NewKeeper(
		c,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtype.String(),
		address.NewBech32Codec("cosmos"),
		ek,
	)
	s.msgSrvr = keeper.NewMsgServerImpl(s.keeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

package keeper

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	committeetypes "github.com/dittonetwork/kepler/x/committee/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/dittonetwork/kepler/x/restaking/keeper"
	restakingtestutil "github.com/dittonetwork/kepler/x/restaking/testutil"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

func RestakingKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	// Mock staking keeper
	ctrl := gomock.NewController(t)
	repository := restakingtestutil.NewMockRepository(ctrl)

	k := keeper.NewKeeper(
		cdc,
		log.NewNopLogger(),
		repository,
		restakingtestutil.NewMockAccountKeeper(ctrl),
		authtypes.NewModuleAddress(committeetypes.ModuleName).String(),
		restakingtestutil.NewMockEpochsKeeper(ctrl),
		"hour",
		address.NewBech32Codec("dittovaloper"),
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	return *k, ctx
}

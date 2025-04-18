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
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	committeemock "github.com/dittonetwork/kepler/x/committee/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/dittonetwork/kepler/x/committee/keeper"
	"github.com/dittonetwork/kepler/x/committee/types"
)

func CommitteeKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	// Mock staking keeper
	ctrl := gomock.NewController(t)
	restakingKeeper := committeemock.NewMockRestakingKeeper(ctrl)
	accountKeeper := committeemock.NewMockAccountKeeper(ctrl)
	repo := committeemock.NewMockRepository(ctrl)
	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	k := keeper.NewKeeper(
		authority.String(),
		accountKeeper,
		restakingKeeper,
		repo,
		ctx.Logger(),
		nil,
		codec.NewLegacyAmino(),
		codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
		"hour",
		address.NewBech32Codec("dittovaloper"),
	)

	// Initialize params
	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx
}

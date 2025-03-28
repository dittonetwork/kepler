package keeper_test

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/dittonetwork/kepler/x/executors/keeper"
	restakingTypes "github.com/dittonetwork/kepler/x/restaking/types"

	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/dittonetwork/kepler/x/executors/types/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetEmergencyExecutors_Success(t *testing.T) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	// Mock staking keeper
	ctrl := gomock.NewController(t)
	restaking := mock.NewMockRestakingKeeper(ctrl)

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		restaking,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	valAddr, err := sdk.ValAddressFromBech32("cosmosvaloper1w7f3xx7e75p4l7qdym5msqem9rd4dyc4mq79dm")
	require.NoError(t, err)

	// Prepare two executors: one active and one inactive.
	one := types.Executor{
		Address:  valAddr.String(),
		IsActive: true,
	}
	two := types.Executor{
		Address:  "cosmos1address2",
		IsActive: true,
	}

	err = k.Executors.Set(ctx, one.Address, one)
	require.NoError(t, err)
	err = k.Executors.Set(ctx, two.Address, two)
	require.NoError(t, err)

	restaking.EXPECT().GetActiveEmergencyValidators(gomock.Any()).Return([]restakingTypes.EmergencyValidator{
		{
			Address: valAddr,
		},
	})

	req := &types.QueryEmergencyExecutorsRequest{}
	resp, err := k.GetEmergencyExecutors(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Only the emergency executor should be returned.
	require.Len(t, resp.Executors, 1)
	require.Equal(t, "cosmosvaloper1w7f3xx7e75p4l7qdym5msqem9rd4dyc4mq79dm", resp.Executors[0].Address)
}

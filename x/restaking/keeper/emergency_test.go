package keeper_test

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/dittonetwork/kepler/x/restaking/keeper"
	"github.com/dittonetwork/kepler/x/restaking/types"
	restakingmock "github.com/dittonetwork/kepler/x/restaking/types/mock"
)

func TestGetActiveEmergencyValidators(t *testing.T) {
	// Initialize SDK configuration for tests
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("cosmos", "cosmospub")
	config.SetBech32PrefixForValidator("cosmosvaloper", "cosmosvaloperpub")
	config.SetBech32PrefixForConsensusNode("cosmosvalcons", "cosmosvalconspub")

	// Function to create a fresh test setup for each test
	setupTest := func(t *testing.T) (sdk.Context, *restakingmock.MockStakingKeeper, keeper.Keeper, *gomock.Controller) {
		// Create a new controller for each test
		ctrl := gomock.NewController(t)

		// Create a mock staking keeper
		stakingKeeper := restakingmock.NewMockStakingKeeper(ctrl)

		// Create a test Keeper with the mock staking keeper
		storeKey := storetypes.NewKVStoreKey(types.StoreKey)
		db := dbm.NewMemDB()
		stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
		stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
		require.NoError(t, stateStore.LoadLatestVersion())

		registry := codectypes.NewInterfaceRegistry()
		cdc := codec.NewProtoCodec(registry)
		authority := authtypes.NewModuleAddress(govtypes.ModuleName)

		k := keeper.NewKeeper(
			cdc,
			runtime.NewKVStoreService(storeKey),
			log.NewNopLogger(),
			authority.String(),
			stakingKeeper,
		)

		ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

		// Initialize params
		err := k.SetParams(ctx, types.DefaultParams())
		require.NoError(t, err)

		return ctx, stakingKeeper, k, ctrl
	}

	t.Run("empty list when no emergency validators", func(t *testing.T) {
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// No validators are set up, so we expect an empty list
		validators := k.GetActiveEmergencyValidators(ctx)
		require.Empty(t, validators)
	})

	t.Run("only returns bonded emergency validators", func(t *testing.T) {
		ctx, stakingKeeper, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Create three validators with different statuses
		// One bonded emergency validator
		bondedVal := createValidator(t, 1, stakingtypes.Bonded)
		bondedValAddr := sdk.ValAddress(bondedVal.GetOperator())

		// One unbonded emergency validator
		unbondedVal := createValidator(t, 2, stakingtypes.Unbonded)
		unbondedValAddr := sdk.ValAddress(unbondedVal.GetOperator())

		// One unbonding emergency validator
		unbondingVal := createValidator(t, 3, stakingtypes.Unbonding)
		unbondingValAddr := sdk.ValAddress(unbondingVal.GetOperator())

		// Set up emergency validators in the store
		require.NoError(t, k.SetValidator(ctx, types.Validator{
			CosmosAddress: bondedValAddr.String(),
			IsEmergency:   true,
		}))
		require.NoError(t, k.SetValidator(ctx, types.Validator{
			CosmosAddress: unbondedValAddr.String(),
			IsEmergency:   true,
		}))
		require.NoError(t, k.SetValidator(ctx, types.Validator{
			CosmosAddress: unbondingValAddr.String(),
			IsEmergency:   true,
		}))

		// Set up mock expectations for staking keeper
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), bondedValAddr).
			Return(bondedVal, nil)
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), unbondedValAddr).
			Return(unbondedVal, nil)
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), unbondingValAddr).
			Return(unbondingVal, nil)

		// Call the function under test
		validators := k.GetActiveEmergencyValidators(ctx)

		// Only the bonded validator should be returned
		require.Len(t, validators, 1)
		require.Equal(t, bondedValAddr, validators[0].Address)
		require.Equal(t, bondedVal.GetConsensusPower(sdk.DefaultPowerReduction), validators[0].VotingPower)
	})

	t.Run("handles invalid validator addresses", func(t *testing.T) {
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Set an emergency validator with an invalid address
		require.NoError(t, k.SetValidator(ctx, types.Validator{
			CosmosAddress: "invalid-address",
			IsEmergency:   true,
		}))

		// Call the function under test - should not panic and return empty list
		validators := k.GetActiveEmergencyValidators(ctx)
		require.Empty(t, validators)
	})

	t.Run("handles getValidator errors", func(t *testing.T) {
		ctx, stakingKeeper, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Create a valid validator address
		valPubKey := ed25519.GenPrivKey().PubKey()
		valAddr := sdk.ValAddress(valPubKey.Address())

		// Set up emergency validator in the store
		require.NoError(t, k.SetValidator(ctx, types.Validator{
			CosmosAddress: valAddr.String(),
			IsEmergency:   true,
		}))

		// Mock staking keeper to return an error
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), valAddr).
			Return(stakingtypes.Validator{}, stakingtypes.ErrNoValidatorFound)

		// Call the function under test - should handle the error and return empty list
		validators := k.GetActiveEmergencyValidators(ctx)
		require.Empty(t, validators)
	})
}

// Helper function to create a validator with a specific status
func createValidator(t *testing.T, index int, status stakingtypes.BondStatus) stakingtypes.Validator {
	pubKey := ed25519.GenPrivKey().PubKey()
	anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
	require.NoError(t, err)

	power := int64(100 + index) // Different voting power for each validator
	tokens := sdk.TokensFromConsensusPower(power, sdk.DefaultPowerReduction)

	return stakingtypes.Validator{
		OperatorAddress: sdk.ValAddress(pubKey.Address()).String(),
		ConsensusPubkey: anyPubKey,
		Status:          status,
		Tokens:          tokens,
	}
}

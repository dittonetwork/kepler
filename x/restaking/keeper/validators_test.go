package keeper_test

import (
	"errors"
	"testing"
	"time"

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

func TestUpdateValidatorSet(t *testing.T) {
	// Initialize SDK configuration for tests
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("cosmos", "cosmospub")
	config.SetBech32PrefixForValidator("cosmosvaloper", "cosmosvaloperpub")
	config.SetBech32PrefixForConsensusNode("cosmosvalcons", "cosmosvalconspub")

	pubKey := ed25519.GenPrivKey().PubKey()

	// Generate Bech32 encoded public keys
	pubKeyBech32Acc := sdk.MustBech32ifyAddressBytes("cosmospub", pubKey.Address())
	pubKeyBech32Val := sdk.MustBech32ifyAddressBytes("cosmosvaloperpub", pubKey.Address())

	operatorAddress := sdk.AccAddress(pubKey.Address())
	valAddress := sdk.ValAddress(operatorAddress)

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

		// Initialize LastUpdate with a default value to support the new validation in UpdateValidatorSet
		err = k.LastUpdate.Set(ctx, types.LastUpdate{
			Epoch:       0,
			Timestamp:   time.Now(),
			BlockHeight: 0,
			BlockHash:   "initial-hash",
		})
		require.NoError(t, err)

		return ctx, stakingKeeper, k, ctrl
	}

	// Function to create a test setup without initializing LastUpdate
	setupWithoutLastUpdate := func(t *testing.T) (sdk.Context, *restakingmock.MockStakingKeeper, keeper.Keeper, *gomock.Controller) {
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

		// Don't initialize LastUpdate for this setup

		return ctx, stakingKeeper, k, ctrl
	}

	t.Run("creating new validator", func(t *testing.T) {
		ctx, stakingKeeper, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Setup mock expectations
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), valAddress).
			Return(stakingtypes.Validator{}, stakingtypes.ErrNoValidatorFound).AnyTimes()

		// Allow any number of calls to SetValidator with any parameters
		stakingKeeper.EXPECT().
			SetValidator(gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()

		params := types.UpdateValidatorSetParams{
			Operators: []types.Operator{
				{
					Address:   operatorAddress.String(),
					PublicKey: pubKeyBech32Acc,
					Status:    types.OperatorStatusBonded,
					Tokens:    100,
				},
			},
			EpochNumber: 1, // Ensure EpochNumber is greater than the default (0)
			BlockHeight: 1, // Ensure BlockHeight is greater than the default (0)
			BlockHash:   "test-hash",
		}

		err := k.UpdateValidatorSet(ctx, params)
		require.NoError(t, err)
	})

	t.Run("updating validator", func(t *testing.T) {
		ctx, stakingKeeper, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Create an existing validator
		anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
		require.NoError(t, err)

		existingValidator := stakingtypes.Validator{
			OperatorAddress: valAddress.String(),
			ConsensusPubkey: anyPubKey,
			Status:          stakingtypes.Bonded,
			Tokens:          sdk.TokensFromConsensusPower(50, sdk.DefaultPowerReduction),
		}

		// Setup mock expectations for getting the validator
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), valAddress).
			Return(existingValidator, nil)

		// First call to SetValidator for updating pubkey
		stakingKeeper.EXPECT().
			SetValidator(gomock.Any(), gomock.Any()).
			Return(nil)

		// Second call to SetValidator for updating status and tokens
		stakingKeeper.EXPECT().
			SetValidator(gomock.Any(), gomock.Any()).
			Return(nil)

		params := types.UpdateValidatorSetParams{
			Operators: []types.Operator{
				{
					Address:   operatorAddress.String(),
					PublicKey: pubKeyBech32Val,
					Status:    types.OperatorStatusBonded,
					Tokens:    200,
				},
			},
			EpochNumber: 1, // Ensure EpochNumber is greater than the default (0)
			BlockHeight: 1, // Ensure BlockHeight is greater than the default (0)
			BlockHash:   "test-hash",
		}

		err = k.UpdateValidatorSet(ctx, params)
		require.NoError(t, err)
	})

	t.Run("should handle invalid operator address", func(t *testing.T) {
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		params := types.UpdateValidatorSetParams{
			Operators: []types.Operator{
				{
					Address:   "invalid-address",
					PublicKey: pubKeyBech32Val,
					Status:    types.OperatorStatusBonded,
					Tokens:    100,
				},
			},
			EpochNumber: 1, // Ensure EpochNumber is greater than the default (0)
			BlockHeight: 1, // Ensure BlockHeight is greater than the default (0)
			BlockHash:   "test-hash",
		}

		err := k.UpdateValidatorSet(ctx, params)
		require.NoError(t, err) // Should not return error, just skip invalid addresses
	})

	t.Run("should handle invalid public key", func(t *testing.T) {
		ctx, stakingKeeper, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Setup mock expectations
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), valAddress).
			Return(stakingtypes.Validator{}, stakingtypes.ErrNoValidatorFound)

		params := types.UpdateValidatorSetParams{
			Operators: []types.Operator{
				{
					Address:   operatorAddress.String(),
					PublicKey: "invalid-pubkey",
					Status:    types.OperatorStatusBonded,
					Tokens:    100,
				},
			},
			EpochNumber: 1, // Ensure EpochNumber is greater than the default (0)
			BlockHeight: 1, // Ensure BlockHeight is greater than the default (0)
			BlockHash:   "test-hash",
		}

		err := k.UpdateValidatorSet(ctx, params)
		require.NoError(t, err) // Should not return error, just skip invalid public keys
	})

	t.Run("should return error on validator fetch failure", func(t *testing.T) {
		ctx, stakingKeeper, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Setup mock expectations for a different error than Not Found
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), valAddress).
			Return(stakingtypes.Validator{}, errors.New("database error"))

		params := types.UpdateValidatorSetParams{
			Operators: []types.Operator{
				{
					Address:   operatorAddress.String(),
					PublicKey: pubKeyBech32Val,
					Status:    types.OperatorStatusBonded,
					Tokens:    100,
				},
			},
			EpochNumber: 1, // Ensure EpochNumber is greater than the default (0)
			BlockHeight: 1, // Ensure BlockHeight is greater than the default (0)
			BlockHash:   "test-hash",
		}

		err := k.UpdateValidatorSet(ctx, params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get validator")
	})

	t.Run("should return error on validator set failure", func(t *testing.T) {
		ctx, stakingKeeper, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Create an existing validator
		anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
		require.NoError(t, err)

		existingValidator := stakingtypes.Validator{
			OperatorAddress: valAddress.String(),
			ConsensusPubkey: anyPubKey,
			Status:          stakingtypes.Bonded,
			Tokens:          sdk.TokensFromConsensusPower(50, sdk.DefaultPowerReduction),
		}

		// Setup mock expectations
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), valAddress).
			Return(existingValidator, nil)

		// Set a failure when trying to set the validator
		stakingKeeper.EXPECT().
			SetValidator(gomock.Any(), gomock.Any()).
			Return(errors.New("failed to set validator"))

		params := types.UpdateValidatorSetParams{
			Operators: []types.Operator{
				{
					Address:   operatorAddress.String(),
					PublicKey: pubKeyBech32Val,
					Status:    types.OperatorStatusBonded,
					Tokens:    100,
				},
			},
			EpochNumber: 1, // Ensure EpochNumber is greater than the default (0)
			BlockHeight: 1, // Ensure BlockHeight is greater than the default (0)
			BlockHash:   "test-hash",
		}

		err = k.UpdateValidatorSet(ctx, params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to set validator")
	})

	t.Run("handling different validator statuses", func(t *testing.T) {
		t.Run("unbonded status", func(t *testing.T) {
			ctx, stakingKeeper, k, ctrl := setupTest(t)
			defer ctrl.Finish()

			// Create an existing validator
			anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
			require.NoError(t, err)

			existingValidator := stakingtypes.Validator{
				OperatorAddress: valAddress.String(),
				ConsensusPubkey: anyPubKey,
				Status:          stakingtypes.Bonded,
				Tokens:          sdk.TokensFromConsensusPower(50, sdk.DefaultPowerReduction),
			}

			// Setup mock expectations
			stakingKeeper.EXPECT().
				GetValidator(gomock.Any(), valAddress).
				Return(existingValidator, nil)

			// Set Validator calls
			stakingKeeper.EXPECT().
				SetValidator(gomock.Any(), gomock.Any()).
				Return(nil).
				AnyTimes()

			params := types.UpdateValidatorSetParams{
				Operators: []types.Operator{
					{
						Address:   operatorAddress.String(),
						PublicKey: pubKeyBech32Val,
						Status:    types.OperatorStatusUnbonded,
						Tokens:    100,
					},
				},
				EpochNumber: 1, // Ensure EpochNumber is greater than the default (0)
				BlockHeight: 1, // Ensure BlockHeight is greater than the default (0)
				BlockHash:   "test-hash",
			}

			err = k.UpdateValidatorSet(ctx, params)
			require.NoError(t, err)

			// Check validator status in local store
			val, err := k.ValidatorsMap.Get(ctx, operatorAddress.String())
			require.NoError(t, err)
			require.Equal(t, types.ValidatorStatus_VALIDATOR_STATUS_UNBONDED, val.Status)
		})

		t.Run("unbonding status", func(t *testing.T) {
			ctx, stakingKeeper, k, ctrl := setupTest(t)
			defer ctrl.Finish()

			// Create an existing validator
			anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
			require.NoError(t, err)

			existingValidator := stakingtypes.Validator{
				OperatorAddress: valAddress.String(),
				ConsensusPubkey: anyPubKey,
				Status:          stakingtypes.Bonded,
				Tokens:          sdk.TokensFromConsensusPower(50, sdk.DefaultPowerReduction),
			}

			// Setup mock expectations
			stakingKeeper.EXPECT().
				GetValidator(gomock.Any(), valAddress).
				Return(existingValidator, nil)

			// Set Validator calls
			stakingKeeper.EXPECT().
				SetValidator(gomock.Any(), gomock.Any()).
				Return(nil).
				AnyTimes()

			params := types.UpdateValidatorSetParams{
				Operators: []types.Operator{
					{
						Address:   operatorAddress.String(),
						PublicKey: pubKeyBech32Val,
						Status:    types.OperatorStatusUnbonding,
						Tokens:    100,
					},
				},
				EpochNumber: 1, // Ensure EpochNumber is greater than the default (0)
				BlockHeight: 1, // Ensure BlockHeight is greater than the default (0)
				BlockHash:   "test-hash",
			}

			err = k.UpdateValidatorSet(ctx, params)
			require.NoError(t, err)

			// Check validator status in local store
			val, err := k.ValidatorsMap.Get(ctx, operatorAddress.String())
			require.NoError(t, err)
			require.Equal(t, types.ValidatorStatus_VALIDATOR_STATUS_UNBONDING, val.Status)
		})
	})

	// New tests for updated code
	t.Run("should return error when update is not needed", func(t *testing.T) {
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Set last update epoch to 10
		err := k.LastUpdate.Set(ctx, types.LastUpdate{
			Epoch:       10,
			Timestamp:   time.Now(),
			BlockHeight: 100,
			BlockHash:   "test-hash",
		})
		require.NoError(t, err)

		// Test with current epoch 5 (less than last update)
		params := types.UpdateValidatorSetParams{
			EpochNumber: 5,
			BlockHeight: 200,
			BlockHash:   "new-hash",
			Operators:   []types.Operator{},
		}

		err = k.UpdateValidatorSet(ctx, params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "update not needed")
	})

	t.Run("should return error when block height is lower than last update", func(t *testing.T) {
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Set last update with epoch 5 and block height 100
		err := k.LastUpdate.Set(ctx, types.LastUpdate{
			Epoch:       5,
			Timestamp:   time.Now(),
			BlockHeight: 100,
			BlockHash:   "test-hash",
		})
		require.NoError(t, err)

		// Test with current epoch 10 (greater than last update) but block height 50 (lower than last update)
		params := types.UpdateValidatorSetParams{
			EpochNumber: 10,
			BlockHeight: 50, // Lower than last update
			BlockHash:   "new-hash",
			Operators:   []types.Operator{},
		}

		err = k.UpdateValidatorSet(ctx, params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "block height is lower than last update")
	})

	t.Run("should process update when conditions are met", func(t *testing.T) {
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Set last update with epoch 5 and block height 100
		err := k.LastUpdate.Set(ctx, types.LastUpdate{
			Epoch:       5,
			Timestamp:   time.Now(),
			BlockHeight: 100,
			BlockHash:   "test-hash",
		})
		require.NoError(t, err)

		// Test with current epoch 10 (greater than last update) and block height 200 (greater than last update)
		// But no operators to process, so it should just update the LastUpdate
		params := types.UpdateValidatorSetParams{
			EpochNumber: 10,
			BlockHeight: 200,
			BlockHash:   "new-hash",
			Operators:   []types.Operator{},
		}

		err = k.UpdateValidatorSet(ctx, params)
		require.NoError(t, err)

		// Verify LastUpdate was updated
		lastUpdate, err := k.LastUpdate.Get(ctx)
		require.NoError(t, err)
		require.Equal(t, params.EpochNumber, lastUpdate.Epoch)
		require.Equal(t, params.BlockHeight, lastUpdate.BlockHeight)
		require.Equal(t, params.BlockHash, lastUpdate.BlockHash)
	})

	t.Run("should return error when LastUpdate.Get fails", func(t *testing.T) {
		ctx, _, k, ctrl := setupWithoutLastUpdate(t)
		defer ctrl.Finish()

		// We don't set LastUpdate, so NeedValidatorsUpdate will return an error
		params := types.UpdateValidatorSetParams{
			EpochNumber: 10,
			BlockHeight: 200,
			BlockHash:   "new-hash",
			Operators:   []types.Operator{},
		}

		err := k.UpdateValidatorSet(ctx, params)
		require.Error(t, err)
	})
}

func TestNeedValidatorsUpdate(t *testing.T) {
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

	t.Run("should return true when epoch is greater than last update epoch", func(t *testing.T) {
		// Setup
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Set last update epoch to 5
		err := k.LastUpdate.Set(ctx, types.LastUpdate{
			Epoch:       5,
			Timestamp:   time.Now(),
			BlockHeight: 100,
			BlockHash:   "test-hash",
		})
		require.NoError(t, err)

		// Test with current epoch 10 (greater than last update)
		result, err := k.NeedValidatorsUpdate(ctx, 10)
		require.NoError(t, err)
		require.True(t, result, "should return true when current epoch is greater than last update epoch")
	})

	t.Run("should return false when epoch is equal to last update epoch", func(t *testing.T) {
		// Setup
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Set last update epoch to 5
		err := k.LastUpdate.Set(ctx, types.LastUpdate{
			Epoch:       5,
			Timestamp:   time.Now(),
			BlockHeight: 100,
			BlockHash:   "test-hash",
		})
		require.NoError(t, err)

		// Test with current epoch 5 (equal to last update)
		result, err := k.NeedValidatorsUpdate(ctx, 5)
		require.NoError(t, err)
		require.False(t, result, "should return false when current epoch is equal to last update epoch")
	})

	t.Run("should return false when epoch is less than last update epoch", func(t *testing.T) {
		// Setup
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Set last update epoch to 5
		err := k.LastUpdate.Set(ctx, types.LastUpdate{
			Epoch:       5,
			Timestamp:   time.Now(),
			BlockHeight: 100,
			BlockHash:   "test-hash",
		})
		require.NoError(t, err)

		// Test with current epoch 3 (less than last update)
		result, err := k.NeedValidatorsUpdate(ctx, 3)
		require.NoError(t, err)
		require.False(t, result, "should return false when current epoch is less than last update epoch")
	})

	t.Run("should handle error when LastUpdate.Get fails", func(t *testing.T) {
		ctx, _, k, ctrl := setupTest(t)
		defer ctrl.Finish()

		// We don't set LastUpdate, so Get will return an error

		result, err := k.NeedValidatorsUpdate(ctx, 10)
		require.False(t, result)
		require.Error(t, err)
	})
}

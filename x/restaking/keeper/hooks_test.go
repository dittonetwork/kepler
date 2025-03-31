package keeper_test

import (
	"context"
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

// MockRestakingHooks is a mock implementation of the RestakingHooks interface
type MockRestakingHooks struct {
	BeforeUnbondingCalled    bool
	BeforeUnbondingValidator stakingtypes.ValidatorI
	BeforeUnbondingError     error
}

// Ensure MockRestakingHooks implements the RestakingHooks interface
var _ types.RestakingHooks = &MockRestakingHooks{}

// BeforeValidatorBeginUnbonding implements RestakingHooks
func (m *MockRestakingHooks) BeforeValidatorBeginUnbonding(_ context.Context, validator stakingtypes.ValidatorI) error {
	m.BeforeUnbondingCalled = true
	m.BeforeUnbondingValidator = validator
	return m.BeforeUnbondingError
}

func TestUpdateValidatorSetWithHooks(t *testing.T) {
	// Initialize SDK configuration for tests
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("cosmos", "cosmospub")
	config.SetBech32PrefixForValidator("cosmosvaloper", "cosmosvaloperpub")
	config.SetBech32PrefixForConsensusNode("cosmosvalcons", "cosmosvalconspub")

	pubKey := ed25519.GenPrivKey().PubKey()

	// Generate Bech32 encoded public keys
	pubKeyBech32Val := sdk.MustBech32ifyAddressBytes("cosmosvaloperpub", pubKey.Address())

	operatorAddress := pubKey.Address().String()
	valAddress := sdk.ValAddress(pubKey.Address())

	// Function to create a fresh test setup for each test
	setupTest := func(t *testing.T) (sdk.Context, *restakingmock.MockStakingKeeper, keeper.Keeper, *MockRestakingHooks, *gomock.Controller) {
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

		// Initialize LastUpdate with default values
		err = k.LastUpdate.Set(ctx, types.LastUpdate{
			Epoch:       0,
			Timestamp:   time.Now(),
			BlockHeight: 0,
			BlockHash:   "",
		})
		require.NoError(t, err)

		// Create and set hooks
		mockHooks := &MockRestakingHooks{}
		k.SetHooks(types.NewMultiRestakingHooks(mockHooks))

		return ctx, stakingKeeper, *k, mockHooks, ctrl
	}

	t.Run("BeforeValidatorBeginUnbonding hook should be called for unbonding validators", func(t *testing.T) {
		ctx, stakingKeeper, k, mockHooks, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Create a mock validator for an existing validator
		anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
		require.NoError(t, err)

		existingValidator := stakingtypes.Validator{
			OperatorAddress: valAddress.String(),
			ConsensusPubkey: anyPubKey,
			Status:          stakingtypes.Bonded, // Starting as bonded
			Tokens:          sdk.TokensFromConsensusPower(50, sdk.DefaultPowerReduction),
		}

		// Setup mock expectations
		stakingKeeper.EXPECT().
			GetValidator(gomock.Any(), valAddress).
			Return(existingValidator, nil)

		// First call to SetValidator for updating pubkey
		stakingKeeper.EXPECT().
			SetValidator(gomock.Any(), gomock.Any()).
			Return(nil)

		// Second call to SetValidator for updating status and tokens
		// Before this, the BeforeValidatorBeginUnbonding hook should be called
		stakingKeeper.EXPECT().
			SetValidator(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx sdk.Context, validator stakingtypes.Validator) error {
				// For this test we're verifying the hook was called before this function
				// The validator status would be updated by the staking module normally
				validator.Status = stakingtypes.Unbonding
				return nil
			})

		params := types.UpdateValidatorSetParams{
			Operators: []types.Operator{
				{
					Address:   operatorAddress,
					PublicKey: pubKeyBech32Val,
					Status:    types.OperatorStatusUnbonding, // We want to unbond the validator
					Tokens:    100,
				},
			},
			EpochNumber: 1, // Set a higher epoch number
			BlockHeight: 100,
			BlockHash:   "test_hash",
		}

		err = k.UpdateValidatorSet(ctx, params)
		require.NoError(t, err)

		// Verify that the hook was called
		require.True(t, mockHooks.BeforeUnbondingCalled, "BeforeValidatorBeginUnbonding hook should have been called")
		require.NotNil(t, mockHooks.BeforeUnbondingValidator, "Hook should have received a validator")
		if v, ok := mockHooks.BeforeUnbondingValidator.(stakingtypes.Validator); ok {
			require.Equal(t, valAddress.String(), v.OperatorAddress, "Wrong validator passed to hook")
		} else {
			require.Equal(t, valAddress.String(), mockHooks.BeforeUnbondingValidator.GetOperator(), "Wrong validator passed to hook")
		}
	})

	t.Run("Should handle error from BeforeValidatorBeginUnbonding hook", func(t *testing.T) {
		ctx, stakingKeeper, k, mockHooks, ctrl := setupTest(t)
		defer ctrl.Finish()

		// Set hooks to return an error
		mockHooks.BeforeUnbondingError = errors.New("unauthorized")

		// Create a mock validator
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

		// First call to SetValidator for updating pubkey
		stakingKeeper.EXPECT().
			SetValidator(gomock.Any(), gomock.Any()).
			Return(nil)

		params := types.UpdateValidatorSetParams{
			Operators: []types.Operator{
				{
					Address:   operatorAddress,
					PublicKey: pubKeyBech32Val,
					Status:    types.OperatorStatusUnbonding,
					Tokens:    100,
				},
			},
			EpochNumber: 1, // Set a higher epoch number
			BlockHeight: 100,
			BlockHash:   "test_hash",
		}

		// Expect error from hook
		err = k.UpdateValidatorSet(ctx, params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "unauthorized")

		// Verify hook was called
		require.True(t, mockHooks.BeforeUnbondingCalled)
	})
}

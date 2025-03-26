package keeper_test

import (
	"errors"
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
					Tokens:    200,
				},
			},
		}

		err = k.UpdateValidatorSet(ctx, params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to set validator")
	})

	t.Run("handling different validator statuses", func(t *testing.T) {
		// Test for Unbonded status
		t.Run("unbonded status", func(t *testing.T) {
			ctx, stakingKeeper, k, ctrl := setupTest(t)
			defer ctrl.Finish()

			anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
			require.NoError(t, err)

			existingValidator := stakingtypes.Validator{
				OperatorAddress: valAddress.String(),
				ConsensusPubkey: anyPubKey,
				Status:          stakingtypes.Bonded,
				Tokens:          sdk.TokensFromConsensusPower(50, sdk.DefaultPowerReduction),
			}

			stakingKeeper.EXPECT().
				GetValidator(gomock.Any(), valAddress).
				Return(existingValidator, nil)

			// First call to update public key
			stakingKeeper.EXPECT().
				SetValidator(gomock.Any(), gomock.Any()).
				Return(nil)

			// Second call to update status and tokens
			stakingKeeper.EXPECT().
				SetValidator(gomock.Any(), gomock.Any()).
				Return(nil)

			params := types.UpdateValidatorSetParams{
				Operators: []types.Operator{
					{
						Address:   operatorAddress.String(),
						PublicKey: pubKeyBech32Val,
						Status:    types.OperatorStatusUnbonded,
						Tokens:    100,
					},
				},
			}

			err = k.UpdateValidatorSet(ctx, params)
			require.NoError(t, err)
		})

		// Test for Unbonding status
		t.Run("unbonding status", func(t *testing.T) {
			ctx, stakingKeeper, k, ctrl := setupTest(t)
			defer ctrl.Finish()

			anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
			require.NoError(t, err)

			existingValidator := stakingtypes.Validator{
				OperatorAddress: valAddress.String(),
				ConsensusPubkey: anyPubKey,
				Status:          stakingtypes.Bonded,
				Tokens:          sdk.TokensFromConsensusPower(50, sdk.DefaultPowerReduction),
			}

			stakingKeeper.EXPECT().
				GetValidator(gomock.Any(), valAddress).
				Return(existingValidator, nil)

			// First call to update public key
			stakingKeeper.EXPECT().
				SetValidator(gomock.Any(), gomock.Any()).
				Return(nil)

			// Second call to update status and tokens
			stakingKeeper.EXPECT().
				SetValidator(gomock.Any(), gomock.Any()).
				Return(nil)

			params := types.UpdateValidatorSetParams{
				Operators: []types.Operator{
					{
						Address:   operatorAddress.String(),
						PublicKey: pubKeyBech32Val,
						Status:    types.OperatorStatusUnbonding,
						Tokens:    100,
					},
				},
			}

			err = k.UpdateValidatorSet(ctx, params)
			require.NoError(t, err)
		})
	})
}

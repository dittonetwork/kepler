package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dittonetwork/kepler/x/restaking/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger
		staking      types.StakingKeeper

		hooks types.RestakingHooks

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		ValidatorsMap *collections.IndexedMap[string, types.Validator, Idx]
		LastUpdate    collections.Item[types.LastUpdate]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	stakeKeeper types.StakingKeeper,
) *Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)
	return &Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		staking:      stakeKeeper,
		logger:       logger,
		ValidatorsMap: collections.NewIndexedMap(
			sb,
			types.KeyPrefixValidator,
			types.CollectionNameValidators,
			collections.StringKey,
			codec.CollValue[types.Validator](cdc),
			NewIndexes(sb),
		),
		LastUpdate: collections.NewItem(
			sb,
			types.KeyPrefixLastUpdate,
			"last_update",
			codec.CollValue[types.LastUpdate](cdc),
		),
	}
}

// BondValidator complete the bonding process for a validator recognized in Bonding status.
func (k Keeper) BondValidator(_ context.Context, _ *types.BondValidatorRequest) (*types.BondValidatorResponse, error) {
	// TODO github.com/dittonetwork/kepler/issues/175
	panic("implement me")
}

// ValidatorStatus returns the status of a validator by its operator address.
func (k Keeper) ValidatorStatus(
	_ context.Context,
	_ *types.QueryValidatorStatusRequest,
) (*types.QueryValidatorStatusResponse, error) {
	// TODO github.com/dittonetwork/kepler/issues/176
	panic("implement me")
}

// Validators returns the list of all validators.
func (k Keeper) Validators(_ context.Context, _ *types.QueryValidatorsRequest) (*types.QueryValidatorsResponse, error) {
	// TODO github.com/dittonetwork/kepler/issues/177
	panic("implement me")
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetHooks set the gamm hooks.
func (k *Keeper) SetHooks(rh types.RestakingHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set hooks twice")
	}

	k.hooks = rh

	return k
}

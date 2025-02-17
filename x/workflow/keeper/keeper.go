package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kepler/x/workflow/types"
)

type (
	Keeper interface {
		GetAuthority() string
		Logger() log.Logger

		// collections methods
		InsertAutomation(ctx sdk.Context, automation types.Automation) error
		SetAutomationStatus(ctx sdk.Context, id uint64, status types.AutomationStatus) error
		GetAutomation(ctx sdk.Context, id uint64) (types.Automation, error)
		FindActiveAutomations(ctx sdk.Context) ([]*types.Automation, error)
		GetNextAutomationID(ctx sdk.Context) (uint64, error)
		CancelAutomation(ctx sdk.Context, id uint64) error
		ActivateAutomation(ctx sdk.Context, id uint64) error

		GetParams(ctx context.Context) types.Params
		SetParams(ctx context.Context, params types.Params) error
		GetActiveAutomations(
			goCtx context.Context,
			req *types.QueryGetActiveAutomationsRequest,
		) (*types.QueryGetActiveAutomationsResponse, error)
		Params(
			goCtx context.Context,
			req *types.QueryParamsRequest,
		) (*types.QueryParamsResponse, error)
	}

	BaseKeeper struct {
		cdc             codec.BinaryCodec
		storeService    store.KVStoreService
		logger          log.Logger
		commetteeKeeper types.CommitteeKeeper

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		// Automations key: automationID | value: automation
		// This is used to store automations
		Automations *collections.IndexedMap[uint64, types.Automation, Idx]
		// AutomationID: sequence for monotonically increasing automation IDs
		AutomationID collections.Sequence
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	committeeKeeper types.CommitteeKeeper,
) BaseKeeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)
	return BaseKeeper{
		cdc:             cdc,
		storeService:    storeService,
		authority:       authority,
		logger:          logger,
		commetteeKeeper: committeeKeeper,
		Automations: collections.NewIndexedMap(
			sb,
			types.KeyPrefixAutomation,
			types.CollectionNameAutomations,
			collections.Uint64Key,
			codec.CollValue[types.Automation](cdc),
			NewAutomationIndexes(sb),
		),
	}
}

// GetAuthority returns the module's authority.
func (k BaseKeeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k BaseKeeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

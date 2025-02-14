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

		InsertAutomation(ctx sdk.Context, automation types.Automation) error
		SetAutomationStatus(ctx sdk.Context, id uint64, status types.AutomationStatus) error
		GetAutomation(ctx sdk.Context, id uint64) (types.Automation, error)
		SetActiveAutomation(ctx sdk.Context, id uint64) error
		RemoveActiveAutomation(ctx sdk.Context, id uint64) error
		GetActiveAutomationIDs(ctx sdk.Context) ([]uint64, error)
		GetNextAutomationID(ctx sdk.Context) (uint64, error)

		GetParams(ctx context.Context) types.Params
		SetParams(ctx context.Context, params types.Params) error
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
	}
	BaseKeeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		// Automations key: automationID | value: automation
		// This is used to store automations
		Automations collections.Map[uint64, types.Automation]
		// AutomationQueue: key set of automation ids
		// This is used to store active automation ids
		AutomationQueue collections.KeySet[uint64]
		// AutomationID: sequence for monotonically increasing automation IDs
		AutomationID collections.Sequence
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) BaseKeeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)
	return BaseKeeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		Automations: collections.NewMap(
			sb,
			types.KeyPrefixAutomation,
			types.CollectionNameAutomations,
			collections.Uint64Key,
			codec.CollValue[types.Automation](cdc),
		),
		AutomationQueue: collections.NewKeySet(
			sb,
			types.KeyPrefixActiveAutomations,
			types.CollectionNameActiveAutomations,
			collections.Uint64Key,
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

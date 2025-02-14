package keeper

import (
	"fmt"

	"cosmossdk.io/collections"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kepler/x/workflow/types"
)

type (
	Keeper struct {
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
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)
	return Keeper{
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
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// AddAutomation stores an automation in KVStore.
func (k Keeper) AddAutomation(ctx sdk.Context, automation types.Automation) error {
	has, err := k.Automations.Has(ctx, automation.Id)
	if err != nil {
		return fmt.Errorf("failed to check if automation exists: %w", err)
	}
	if has {
		return fmt.Errorf("automation with id %d already exists", automation.Id)
	}

	err = k.Automations.Set(ctx, automation.Id, automation)
	if err != nil {
		return fmt.Errorf("failed to set automation: %w", err)
	}

	return nil
}

// SetAutomationStatus sets the status of an automation in KVStore.
func (k Keeper) SetAutomationStatus(ctx sdk.Context, id uint64, status types.AutomationStatus) error {
	automation, err := k.GetAutomation(ctx, id)
	if err != nil {
		return err
	}

	automation.Status = status
	err = k.Automations.Set(ctx, id, automation)
	if err != nil {
		return fmt.Errorf("failed to set automation status: %w", err)
	}

	return nil
}

// GetAutomation get an automation by ID.
func (k Keeper) GetAutomation(ctx sdk.Context, id uint64) (types.Automation, error) {
	automation, err := k.Automations.Get(ctx, id)
	if err != nil {
		return types.Automation{}, fmt.Errorf("failed to get automation: %w", err)
	}

	return automation, nil
}

// SetActiveAutomation stores an active automation ID in KVStore.
func (k Keeper) SetActiveAutomation(ctx sdk.Context, id uint64) error {
	err := k.AutomationQueue.Set(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to set active automation: %w", err)
	}

	return nil
}

// RemoveActiveAutomation removes an active automation ID from KVStore.
func (k Keeper) RemoveActiveAutomation(ctx sdk.Context, id uint64) error {
	err := k.AutomationQueue.Remove(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to remove active automation: %w", err)
	}

	return nil
}

// GetActiveAutomations returns all active automation IDs.
func (k Keeper) GetActiveAutomationIDs(ctx sdk.Context) ([]uint64, error) {
	idsIter, err := k.AutomationQueue.Iterate(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get active automations: %w", err)
	}

	ids, err := idsIter.Keys()
	if err != nil {
		return nil, fmt.Errorf("failed to get active automation keys: %w", err)
	}

	return ids, nil
}

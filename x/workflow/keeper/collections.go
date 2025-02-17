package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	"fmt"
	"kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Idx struct {
	// AutomationStatus index by status of automation
	AutomationStatus *indexes.Multi[string, uint64, types.Automation]
}

func NewAutomationIndexes(sb *collections.SchemaBuilder) Idx {
	return Idx{
		AutomationStatus: indexes.NewMulti(
			sb,
			types.KeyPrefixAutomation,
			types.CollectionIndexAutomationByStatus,
			collections.StringKey,
			collections.Uint64Key,
			func(pk uint64, val types.Automation) (string, error) {
				return val.GetStatus().String(), nil
			}),
	}
}

func (a Idx) IndexesList() []collections.Index[uint64, types.Automation] {
	return []collections.Index[uint64, types.Automation]{
		a.AutomationStatus,
	}
}

// InsertAutomation stores an automation in KVStore.
func (k BaseKeeper) InsertAutomation(ctx sdk.Context, automation types.Automation) error {
	has, err := k.Automations.Has(ctx, automation.Id)
	if err != nil {
		return fmt.Errorf("failed to check if automation exists: %w", err)
	}
	if has {
		return types.ErrAutomationAlreadyExists
	}

	err = k.Automations.Set(ctx, automation.Id, automation)
	if err != nil {
		return fmt.Errorf("failed to set automation: %w", err)
	}

	return nil
}

// SetAutomationStatus sets the status of an automation in KVStore.
func (k BaseKeeper) SetAutomationStatus(ctx sdk.Context, id uint64, status types.AutomationStatus) error {
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
func (k BaseKeeper) GetAutomation(ctx sdk.Context, id uint64) (types.Automation, error) {
	automation, err := k.Automations.Get(ctx, id)
	if err != nil {
		return types.Automation{}, fmt.Errorf("failed to get automation: %w", err)
	}

	return automation, nil
}

// FindActiveAutomations returns all active automation IDs.
func (k BaseKeeper) FindActiveAutomations(ctx sdk.Context) ([]*types.Automation, error) {
	iter, err := k.Automations.Indexes.AutomationStatus.MatchExact(
		ctx,
		types.AutomationStatus_AUTOMATION_STATUS_ACTIVE.String(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get active automations: %w", err)
	}

	pks, err := iter.PrimaryKeys()
	if err != nil {
		return nil, fmt.Errorf("failed to get primary keys: %w", err)
	}

	automations := make([]*types.Automation, len(pks))
	for i, pk := range pks {
		automation, err := k.GetAutomation(ctx, pk)
		if err != nil {
			return nil, fmt.Errorf("failed to get automation: %w", err)
		}
		automations[i] = &automation
	}

	return automations, nil
}

// GetNextAutomationID returns the next automation ID.
func (k BaseKeeper) GetNextAutomationID(ctx sdk.Context) (uint64, error) {
	id, err := k.AutomationID.Next(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get next automation ID: %w", err)
	}

	return id, nil
}

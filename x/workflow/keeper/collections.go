package keeper

import (
	"fmt"
	"kepler/x/workflow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

// SetActiveAutomation stores an active automation ID in KVStore.
func (k BaseKeeper) SetActiveAutomation(ctx sdk.Context, id uint64) error {
	err := k.AutomationQueue.Set(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to set active automation: %w", err)
	}

	return nil
}

// RemoveActiveAutomation removes an active automation ID from KVStore.
func (k BaseKeeper) RemoveActiveAutomation(ctx sdk.Context, id uint64) error {
	err := k.AutomationQueue.Remove(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to remove active automation: %w", err)
	}

	return nil
}

// GetActiveAutomationIDs returns all active automation IDs.
func (k BaseKeeper) GetActiveAutomationIDs(ctx sdk.Context) ([]uint64, error) {
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

// GetNextAutomationID returns the next automation ID.
func (k BaseKeeper) GetNextAutomationID(ctx sdk.Context) (uint64, error) {
	id, err := k.AutomationID.Next(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get next automation ID: %w", err)
	}

	return id, nil
}

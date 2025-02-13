package keeper

import (
	"fmt"

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

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
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

// AddAutomation stores an automation in KVStore
func (k Keeper) AddAutomation(ctx sdk.Context, automation types.Automation) error {
	s := k.storeService.OpenKVStore(ctx)
	key := []byte(types.KeyPrefixAutomation + fmt.Sprint(automation.Id)) // Key format: "automation:1"

	bz, err := k.cdc.Marshal(&automation)
	if err != nil {
		k.Logger().Error("failed to marshal automation", "error", err)
		return err
	}

	err = s.Set(key, bz)
	if err != nil {
		return err
	}

	return nil
}

// SetAutomationStatus sets the status of an automation in KVStore
func (k Keeper) SetAutomationStatus(ctx sdk.Context, id uint64, status types.AutomationStatus) error {
	s := k.storeService.OpenKVStore(ctx)
	key := []byte(types.KeyPrefixAutomation + fmt.Sprint(id))

	bz, err := s.Get(key)
	if err != nil {
		k.Logger().Error("failed to get automation", "error", err)
		return err
	}

	var automation types.Automation
	err = k.cdc.Unmarshal(bz, &automation)
	if err != nil {
		k.Logger().Error("failed to unmarshal automation", "error", err)
		return err
	}

	automation.Status = status

	bz, err = k.cdc.Marshal(&automation)
	if err != nil {
		k.Logger().Error("failed to marshal automation", "error", err)
		return err
	}

	err = s.Set(key, bz)
	if err != nil {
		return err
	}

	return nil
}

// GetAutomation get an automation by ID
func (k Keeper) GetAutomation(ctx sdk.Context, id uint64) (types.Automation, error) {
	s := k.storeService.OpenKVStore(ctx)
	key := []byte(types.KeyPrefixAutomation + fmt.Sprint(id))

	bz, err := s.Get(key)
	if err != nil {
		k.Logger().Error("failed to get automation", "error", err)
		return types.Automation{}, err
	}

	var automation types.Automation
	err = k.cdc.Unmarshal(bz, &automation)
	if err != nil {
		k.Logger().Error("failed to unmarshal automation", "error", err)
		return types.Automation{}, err
	}

	return automation, nil
}

// PutActiveAutomation stores an active automation ID in KVStore
func (k Keeper) PutActiveAutomation(ctx sdk.Context, id uint64) error {
	s := k.storeService.OpenKVStore(ctx)
	key := []byte(types.KeyActiveAutomations)

	bz, err := s.Get(key)
	if err != nil {
		k.Logger().Error("failed to get active automations", "error", err)
		return err
	}

	err = s.Set(key, append(bz, sdk.Uint64ToBigEndian(id)...))
	if err != nil {
		k.Logger().Error("failed to put an active automation", "error", err)
		return err
	}

	return nil
}

// RemoveActiveAutomation removes an active automation ID from KVStore
func (k Keeper) RemoveActiveAutomation(ctx sdk.Context, id uint64) error {
	s := k.storeService.OpenKVStore(ctx)
	key := []byte(types.KeyActiveAutomations)

	bz, err := s.Get(key)
	if err != nil {
		k.Logger().Error("failed to get active automations", "error", err)
		return err
	}
	if bz == nil {
		return nil
	}

	var (
		updatedBytes []byte
		removed      bool
	)
	for i := 0; i < len(bz); i += 8 {
		currentID := sdk.BigEndianToUint64(bz[i : i+8])
		if currentID != id {
			updatedBytes = append(updatedBytes, bz[i:i+8]...)
		} else {
			removed = true
		}
	}

	if !removed {
		return nil
	}

	err = s.Set(key, updatedBytes)
	if err != nil {
		k.Logger().Error("failed to update active automations", "error", err)
		return err
	}

	return nil
}

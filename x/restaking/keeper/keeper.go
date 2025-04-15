package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

type (
	Keeper struct {
		cdc    codec.BinaryCodec
		logger log.Logger
		hooks  types.RestakingHooks

		// the address capable of executing a MsgUpdateValidatorsSet message. Typically, this
		// should be the x/committee module account.
		authority string

		// keeper dependencies
		epochs types.EpochsKeeper

		repository  types.Repository
		mainEpochID string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	logger log.Logger,
	repo types.Repository,
	ak types.AccountKeeper,
	authority string,
	epochs types.EpochsKeeper,
	mainEpochID string,
) *Keeper {
	// ensure that authority is a valid AccAddress
	if _, err := ak.AddressCodec().StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("authority is not valid acc address: %s", authority))
	}

	return &Keeper{
		cdc:         cdc,
		logger:      logger,
		authority:   authority,
		repository:  repo,
		epochs:      epochs,
		mainEpochID: mainEpochID,
	}
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

// GetAuthority returns the x/restaking module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

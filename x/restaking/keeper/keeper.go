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

		repository types.Repository
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	logger log.Logger,
	repo types.Repository,
) *Keeper {
	return &Keeper{
		cdc:        cdc,
		logger:     logger,
		repository: repo,
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

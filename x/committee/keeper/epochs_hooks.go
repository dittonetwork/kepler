package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

var _ types.EpochHooks = EpochsHooks{}

type EpochsHooks struct {
	keeper Keeper
}

// EpochsHooks implements the EpochsHooks interface.
func (k Keeper) EpochsHooks() EpochsHooks {
	return EpochsHooks{k}
}

func (e EpochsHooks) AfterEpochEnd(ctx context.Context, id string, number int64) error {
	// skip if not the main epoch
	if id != e.keeper.epochMainID {
		return nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	_, err := e.keeper.CreateCommittee(sdkCtx, number)
	if err != nil {
		return err
	}

	return nil
}

func (e EpochsHooks) BeforeEpochStart(_ context.Context, _ string, _ int64) error {
	return nil // noop just for save interface
}

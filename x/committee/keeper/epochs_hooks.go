package keeper

import (
	"context"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/committee/types"
)

var _ types.EpochsHooks = EpochsHooks{}

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
	// @TODO: need change type of epoch number to uint32 for avoid this check
	// https://github.com/dittonetwork/kepler/issues/208
	if number < 0 || number > math.MaxUint32 {
		e.keeper.Logger(ctx).With("number", number).Error("invalid epoch number")
		return fmt.Errorf("invalid epoch number: %d", number)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	_, err := e.keeper.CreateCommittee(sdkCtx, uint32(number))
	if err != nil {
		return err
	}

	return nil
}

func (e EpochsHooks) BeforeEpochStart(_ context.Context, _ string, _ int64) error {
	return nil // noop just for save interface
}

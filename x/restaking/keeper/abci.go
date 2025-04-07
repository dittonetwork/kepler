package keeper

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

func (k Keeper) BeginBlocker(_ context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, telemetry.Now(), telemetry.MetricKeyBeginBlocker)

	// Track historical info about validators - this will be implemented in a future update
	// Tracking helps with monitoring validator performance and analyzing historical trends
	return nil
}

func (k Keeper) EndBlocker(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, telemetry.Now(), telemetry.MetricKeyEndBlocker)

	return k.ApplyAndReturnValidatorSetUpdates(ctx)
}

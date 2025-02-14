package keeper

import (
	"context"
	"fmt"
	"kepler/x/epochs/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker runs at the start of every block.
func (k Keeper) BeginBlocker(ctx context.Context) error {
	start := telemetry.Now()
	defer telemetry.ModuleMeasureSince(types.ModuleName, start, telemetry.MetricKeyBeginBlocker)

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	headerInfo := sdkCtx.HeaderInfo()

	return k.EpochInfo.Walk(
		ctx,
		nil,
		func(_ string, epochInfo types.EpochInfo) (bool, error) {
			if headerInfo.Time.Before(epochInfo.StartTime) {
				return false, nil
			}
			// if epoch counting hasn't started, signal we need to start.
			shouldInitialEpochStart := !epochInfo.EpochCountingStarted

			epochEndTime := epochInfo.CurrentEpochStartTime.Add(epochInfo.Duration)
			shouldEpochStart := (headerInfo.Time.After(epochEndTime)) || shouldInitialEpochStart

			if !shouldEpochStart {
				return false, nil
			}

			epochInfo.CurrentEpochStartHeight = headerInfo.Height

			if shouldInitialEpochStart {
				epochInfo.EpochCountingStarted = true
				epochInfo.CurrentEpoch = 1
				epochInfo.CurrentEpochStartTime = epochInfo.StartTime
				k.Logger().Debug(
					fmt.Sprintf("Starting new epoch with identifier %s epoch number %d",
						epochInfo.Identifier,
						epochInfo.CurrentEpoch,
					),
				)
			} else {
				if err := sdkCtx.EventManager().EmitTypedEvent(&types.EventEpochEnd{
					EpochNumber: epochInfo.CurrentEpoch,
				}); err != nil {
					return false, nil //nolint:nilerr // error is logged
				}

				if err := k.AfterEpochEnd(ctx, epochInfo.Identifier, epochInfo.CurrentEpoch); err != nil {
					// purposely ignoring the error here not to halt the chain if the hook fails
					k.Logger().Error(
						fmt.Sprintf(
							"Error after epoch end with identifier %s epoch number %d",
							epochInfo.Identifier,
							epochInfo.CurrentEpoch,
						),
					)
				}

				epochInfo.CurrentEpoch++
				epochInfo.CurrentEpochStartTime = epochInfo.CurrentEpochStartTime.Add(epochInfo.Duration)

				k.Logger().Debug(
					fmt.Sprintf("Starting epoch with identifier %s epoch number %d",
						epochInfo.Identifier,
						epochInfo.CurrentEpoch,
					),
				)
			}

			// emit new epoch start event, set epoch info, and run BeforeEpochStart hook
			err := sdkCtx.EventManager().EmitTypedEvent(&types.EventEpochStart{
				EpochNumber:    epochInfo.CurrentEpoch,
				EpochStartTime: epochInfo.CurrentEpochStartTime.Unix(),
			})
			if err != nil {
				return false, err
			}
			err = k.EpochInfo.Set(ctx, epochInfo.Identifier, epochInfo)
			if err != nil {
				k.Logger().Error(
					fmt.Sprintf(
						"Error set epoch info with identifier %s epoch number %d",
						epochInfo.Identifier,
						epochInfo.CurrentEpoch,
					),
				)
				return false, nil //nolint:nilerr // error is logged
			}

			if err = k.BeforeEpochStart(ctx, epochInfo.Identifier, epochInfo.CurrentEpoch); err != nil {
				// purposely ignoring the error here not to halt the chain if the hook fails
				k.Logger().Error(
					fmt.Sprintf(
						"Error before epoch start with identifier %s epoch number %d",
						epochInfo.Identifier,
						epochInfo.CurrentEpoch,
					),
				)
			}
			return false, nil
		},
	)
}

package committee

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AfterEpochEnd is called when epoch is going to be ended, epochNumber is the number of epoch that is ending.
func (am AppModule) AfterEpochEnd(ctx sdk.Context, id string, number int64) error {
	if id == am.epochMainID {
		// @TODO: need change type of epoch number to uint32 for avoid this check
		// https://github.com/dittonetwork/kepler/issues/208
		if number < 0 || number > math.MaxUint32 {
			am.keeper.Logger().With("number", number).Error("invalid epoch number")
			return fmt.Errorf("invalid epoch number: %d", number)
		}

		_, err := am.keeper.CreateCommittee(ctx, uint32(number))
		if err != nil {
			return err
		}
	}

	return nil
}

// BeforeEpochStart is called when epoch is going to be started, epochNumber is the number of epoch that is starting.
func (am AppModule) BeforeEpochStart(_ sdk.Context, _ string, _ int64) error {
	// noop
	return nil
}

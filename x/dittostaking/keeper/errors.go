package keeper

import (
	"fmt"
)

func WrapCosmosStakingError(err error) error {
	if err != nil {
		return fmt.Errorf("cosmos staking module error: %w", err)
	}

	return nil
}

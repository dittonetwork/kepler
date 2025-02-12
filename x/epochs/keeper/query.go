package keeper

import (
	"kepler/x/epochs/types"
)

var _ types.QueryServer = Keeper{}

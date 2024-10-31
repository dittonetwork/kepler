package keeper

import (
	"kepler/x/beacon/types"
)

var _ types.QueryServer = Keeper{}

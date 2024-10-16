package keeper

import (
	"kepler/x/kepler/types"
)

var _ types.QueryServer = Keeper{}

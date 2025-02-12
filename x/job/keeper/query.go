package keeper

import (
	"kepler/x/job/types"
)

var _ types.QueryServer = Keeper{}

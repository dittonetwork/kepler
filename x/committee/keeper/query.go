package keeper

import (
	"kepler/x/committee/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"github.com/dittonetwork/kepler/x/executors/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"github.com/dittonetwork/kepler/x/job/types"
)

var _ types.QueryServer = Keeper{}

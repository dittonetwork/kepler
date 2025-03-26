package keeper

import (
	"github.com/dittonetwork/kepler/x/restaking/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"github.com/dittonetwork/kepler/x/committee/types"
)

var _ types.QueryServer = Keeper{}

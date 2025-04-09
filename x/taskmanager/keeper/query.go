package keeper

import (
	"github.com/dittonetwork/kepler/x/taskmanager/types"
)

var _ types.QueryServer = Keeper{}

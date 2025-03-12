package keeper

import (
	"github.com/dittonetwork/kepler/x/workflow/types"
)

type Querier struct {
	BaseKeeper
}

var _ types.QueryServer = BaseKeeper{}

func NewQuerier(k *BaseKeeper) Querier {
	return Querier{*k}
}

package keeper

import (
	"github.com/dittonetwork/kepler/x/executors/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{keeper}
}

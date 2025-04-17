package genutil

import (
	"cosmossdk.io/core/genesis"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions.
func InitGenesis(
	ctx sdk.Context, restaking types.RestakingKeeper,
	deliverTx genesis.TxHandler, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) ([]abci.ValidatorUpdate, error) {
	if len(genesisState.GenTxs) > 0 {
		return DeliverGenTxs(ctx, genesisState.GenTxs, restaking, deliverTx, txEncodingConfig)
	}

	return nil, nil
}

package genutil

import (
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/genesis"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/genutil/types"
)

// SetGenTxsInAppGenesisState - sets the genesis transactions in the app genesis state.
func SetGenTxsInAppGenesisState(
	cdc codec.JSONCodec, txJSONEncoder sdk.TxEncoder, appGenesisState map[string]json.RawMessage, genTxs []sdk.Tx,
) (map[string]json.RawMessage, error) {
	genesisState := types.GetGenesisStateFromAppState(cdc, appGenesisState)
	genTxsBz := make([]json.RawMessage, 0, len(genTxs))

	for _, genTx := range genTxs {
		txBz, err := txJSONEncoder(genTx)
		if err != nil {
			return appGenesisState, err
		}

		genTxsBz = append(genTxsBz, txBz)
	}

	genesisState.GenTxs = genTxsBz
	return types.SetGenesisStateInAppState(cdc, appGenesisState, genesisState), nil
}

// DeliverGenTxs iterates over all genesis txs, decodes each into a Tx and
// invokes the provided deliverTxfn with the decoded Tx. It returns the result
// of the staking module's ApplyAndReturnValidatorSetUpdates.
func DeliverGenTxs(
	ctx sdk.Context, genTxs []json.RawMessage,
	restaking types.RestakingKeeper, deliverTx genesis.TxHandler,
	txEncodingConfig client.TxEncodingConfig,
) ([]abci.ValidatorUpdate, error) {
	for _, genTx := range genTxs {
		tx, err := txEncodingConfig.TxJSONDecoder()(genTx)
		if err != nil {
			return nil, fmt.Errorf("failed to decode GenTx '%s': %w", genTx, err)
		}

		bz, err := txEncodingConfig.TxEncoder()(tx)
		if err != nil {
			return nil, fmt.Errorf("failed to encode GenTx '%s': %w", genTx, err)
		}

		err = deliverTx.ExecuteGenesisTx(bz)
		if err != nil {
			return nil, fmt.Errorf("failed to execute DeliverTx for '%s': %w", genTx, err)
		}
	}

	return restaking.ApplyAndReturnValidatorSetUpdates(ctx)
}

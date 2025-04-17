package genutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	cfg "github.com/cometbft/cometbft/config"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankexported "github.com/cosmos/cosmos-sdk/x/bank/exported"
	"github.com/dittonetwork/kepler/x/genutil/types"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"
)

// GenAppStateFromConfig gets the genesis app state from the config.
func GenAppStateFromConfig(
	cdc codec.JSONCodec,
	txEncodingConfig client.TxEncodingConfig,
	config *cfg.Config,
	initCfg types.InitConfig,
	genesis *types.AppGenesis,
	genBalIterator types.GenesisBalancesIterator,
	validator types.MessageValidator,
) (json.RawMessage, error) {
	appGenTxs, persistentPeers, err := CollectTxs(
		cdc, txEncodingConfig.TxJSONDecoder(), config.Moniker, initCfg.GenTxsDir, genesis, genBalIterator, validator,
	)
	if err != nil {
		return nil, err
	}

	config.P2P.PersistentPeers = persistentPeers
	cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

	if len(appGenTxs) == 0 {
		return nil, errors.New("no genesis app genesis txs found")
	}

	// create the app state
	appGenesisState, err := types.GenesisStateFromAppGenesis(genesis)
	if err != nil {
		return nil, err
	}

	appGenesisState, err = SetGenTxsInAppGenesisState(cdc, txEncodingConfig.TxJSONEncoder(), appGenesisState, appGenTxs)
	if err != nil {
		return nil, err
	}

	var appState json.RawMessage
	appState, err = json.MarshalIndent(appGenesisState, "", "  ")
	if err != nil {
		return nil, err
	}

	genesis.AppState = appState

	return appState, ExportGenesisFile(genesis, config.GenesisFile())
}

func CollectTxs(
	cdc codec.JSONCodec,
	txJSONDecoder sdk.TxDecoder,
	moniker, genTxsDir string,
	genesis *types.AppGenesis,
	genBalIterator types.GenesisBalancesIterator,
	validator types.MessageValidator,
) ([]sdk.Tx, string, error) {
	var appGenTxs []sdk.Tx
	var appState map[string]json.RawMessage
	if err := json.Unmarshal(genesis.AppState, &appState); err != nil {
		return nil, "", err
	}

	fos, err := os.ReadDir(genTxsDir)
	if err != nil {
		return nil, "", err
	}

	balancesMap := make(map[string]bankexported.GenesisBalance)

	genBalIterator.IterateGenesisBalances(cdc, appState, func(balance bankexported.GenesisBalance) bool {
		balancesMap[balance.GetAddress()] = balance
		return false
	})

	// addresses and IPs (and port) validator server info
	var addressesIPs []string

	for _, fo := range fos {
		if fo.IsDir() {
			continue
		}

		if !strings.HasSuffix(fo.Name(), ".json") {
			continue
		}

		// get the genTx
		var jsonRawTx []byte
		jsonRawTx, err = os.ReadFile(filepath.Join(genTxsDir, fo.Name()))
		if err != nil {
			return nil, "", err
		}

		var genTx sdk.Tx
		genTx, err = types.ValidateAndGetGenTx(jsonRawTx, txJSONDecoder, validator)
		if err != nil {
			return nil, "", err
		}

		appGenTxs = append(appGenTxs, genTx)

		// the memo flag is used to store
		// the ip and node-id, for example this may be:
		// "528fd3df22b31f4969b05652bfe8f0fe921321d5@192.168.2.37:26656"

		memoTx, ok := genTx.(sdk.TxWithMemo)
		if !ok {
			return nil, "", fmt.Errorf("expected TxWithMemo, got %T", genTx)
		}
		nodeAddrIP := memoTx.GetMemo()

		// genesis transactions must be single-message
		msgs := genTx.GetMsgs()

		msg, ok := msgs[0].(*restakingtypes.MsgBondValidator)
		if !ok {
			return nil, "", fmt.Errorf("expected MsgBondValidator, got %T", msgs[0])
		}

		if msg.Description.Moniker != moniker {
			addressesIPs = append(addressesIPs, nodeAddrIP)
		}
	}

	sort.Strings(addressesIPs)
	persistentPeers := strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}

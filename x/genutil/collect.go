package genutil

import (
	"encoding/json"
	"path/filepath"

	cfg "github.com/cometbft/cometbft/config"
	"github.com/dittonetwork/kepler/x/genutil/types"
)

// GenAppStateFromConfig gets the genesis app state from the config.
func GenAppStateFromConfig(config *cfg.Config, genesis *types.AppGenesis) (json.RawMessage, error) {
	cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

	// create the app state
	appGenesisState, err := types.GenesisStateFromAppGenesis(genesis)
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

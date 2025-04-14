package cli

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math/unsafe"
	cfg "github.com/cometbft/cometbft/config"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/go-bip39"
	"github.com/dittonetwork/kepler/x/genutil"
	"github.com/dittonetwork/kepler/x/genutil/types"
	"github.com/spf13/cobra"
)

const (
	// FlagOverwrite defines a flag to overwrite an existing genesis JSON file.
	FlagOverwrite = "overwrite"

	// FlagRecover defines a flag to initialize the private validator key from a specific seed.
	FlagRecover = "recover"

	// FlagDefaultBondDenom defines the default denom to use in the genesis file.
	FlagDefaultBondDenom = "default-denom"
)

type printInfo struct {
	Moniker    string          `json:"moniker" yaml:"moniker"`
	ChainID    string          `json:"chain_id" yaml:"chain_id"`
	NodeID     string          `json:"node_id" yaml:"node_id"`
	GenTxsDir  string          `json:"gentxs_dir" yaml:"gentxs_dir"`
	AppMessage json.RawMessage `json:"app_message" yaml:"app_message"`
}

func newPrintInfo(moniker, chainID, nodeID, genTxsDir string, appMessage json.RawMessage) printInfo {
	return printInfo{
		Moniker:    moniker,
		ChainID:    chainID,
		NodeID:     nodeID,
		GenTxsDir:  genTxsDir,
		AppMessage: appMessage,
	}
}

func displayInfo(info printInfo) error {
	out, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(os.Stderr, "%s\n", out)

	return err
}

// setupChainID defines and returns the chain ID.
func setupChainID(cmd *cobra.Command, clientCtx client.Context) string {
	chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
	switch {
	case chainID != "":
		return chainID
	case clientCtx.ChainID != "":
		return clientCtx.ChainID
	default:
		const chainIDLength = 6
		return fmt.Sprintf("test-chain-%v", unsafe.Str(chainIDLength))
	}
}

// getMnemonic gets the mnemonic code if needed.
func getMnemonic(cmd *cobra.Command) (string, error) {
	recv, _ := cmd.Flags().GetBool(FlagRecover)
	if !recv {
		return "", nil
	}

	inBuf := bufio.NewReader(cmd.InOrStdin())
	value, err := input.GetString("Enter your bip39 mnemonic", inBuf)
	if err != nil {
		return "", err
	}

	if !bip39.IsMnemonicValid(value) {
		return "", errors.New("invalid mnemonic")
	}

	return value, nil
}

// getInitialHeight gets the initial block height.
func getInitialHeight(cmd *cobra.Command) int64 {
	initHeight, _ := cmd.Flags().GetInt64(flags.FlagInitHeight)
	if initHeight < 1 {
		return 1
	}
	return initHeight
}

// checkGenesisOverwrite checks if the existing genesis file can be overwritten.
func checkGenesisOverwrite(genFile string, overwrite bool) error {
	_, err := os.Stat(genFile)
	if !overwrite && !os.IsNotExist(err) {
		return fmt.Errorf("genesis.json file already exists: %v", genFile)
	}
	return nil
}

// createAppGenesis creates an application genesis object.
func createAppGenesis(
	genFile string,
	cdc codec.Codec,
	mbm module.BasicManager,
	chainID string,
	initHeight int64,
) (*types.AppGenesis, json.RawMessage, error) {
	appGenState := mbm.DefaultGenesis(cdc)

	appState, err := json.MarshalIndent(appGenState, "", " ")
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "Failed to marshal default genesis state")
	}

	appGenesis := &types.AppGenesis{}
	if _, err = os.Stat(genFile); err != nil {
		if !os.IsNotExist(err) {
			return nil, nil, err
		}
	} else {
		appGenesis, err = types.AppGenesisFromFile(genFile)
		if err != nil {
			return nil, nil, errorsmod.Wrap(err, "Failed to read genesis doc from file")
		}
	}

	appGenesis.AppName = version.AppName
	appGenesis.AppVersion = version.Version
	appGenesis.ChainID = chainID
	appGenesis.AppState = appState
	appGenesis.InitialHeight = initHeight
	appGenesis.Consensus = &types.ConsensusGenesis{
		Validators: nil,
	}

	return appGenesis, appState, nil
}

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func InitCmd(mbm module.BasicManager, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [moniker]",
		Short: "Initialize private validator, p2p, genesis, and application configuration files",
		Long:  `Initialize validators's and node's configuration files.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			// setup chain ID
			chainID := setupChainID(cmd, clientCtx)

			// get mnemonic code
			mnemonic, err := getMnemonic(cmd)
			if err != nil {
				return err
			}

			// get initial block height
			initHeight := getInitialHeight(cmd)

			// initialize validator files
			nodeID, _, err := genutil.InitializeNodeValidatorFilesFromMnemonic(config, mnemonic)
			if err != nil {
				return err
			}

			config.Moniker = args[0]

			// check if genesis.json can be overwritten
			genFile := config.GenesisFile()
			overwrite, _ := cmd.Flags().GetBool(FlagOverwrite)
			if errCheck := checkGenesisOverwrite(genFile, overwrite); errCheck != nil {
				return errCheck
			}

			// create application genesis
			appGenesis, appState, err := createAppGenesis(genFile, cdc, mbm, chainID, initHeight)
			if err != nil {
				return err
			}

			// export genesis file
			if err = genutil.ExportGenesisFile(appGenesis, genFile); err != nil {
				return errorsmod.Wrap(err, "Failed to export genesis file")
			}

			toPrint := newPrintInfo(config.Moniker, chainID, nodeID, "", appState)

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
			return displayInfo(toPrint)
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	cmd.Flags().BoolP(FlagOverwrite, "o", false, "overwrite the genesis.json file")
	cmd.Flags().Bool(FlagRecover, false, "provide seed phrase to recover existing key instead of creating")
	cmd.Flags().String(flags.FlagChainID, "",
		"genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(FlagDefaultBondDenom, "",
		"genesis file default denomination, if left blank default value is 'stake'")
	cmd.Flags().Int64(flags.FlagInitHeight, 1, "specify the initial block height at genesis")

	return cmd
}

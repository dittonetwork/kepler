package genutil

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	cfg "github.com/cometbft/cometbft/config"
	tmed25519 "github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cosmos/go-bip39"
	"github.com/dittonetwork/kepler/x/genutil/types"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// ExportGenesisFile creates and writes the genesis configuration to disk. An
// error is returned if building or writing the configuration to file fails.
func ExportGenesisFile(genesis *types.AppGenesis, genFile string) error {
	if err := genesis.ValidateAndComplete(); err != nil {
		return err
	}

	return genesis.SaveAs(genFile)
}

// InitializeNodeValidatorFiles creates private validator and p2p configuration files.
func InitializeNodeValidatorFiles(config *cfg.Config) (string, cryptotypes.PubKey, error) {
	return InitializeNodeValidatorFilesFromMnemonic(config, "")
}

// InitializeNodeValidatorFilesFromMnemonic creates private validator
// and p2p configuration files using the given mnemonic.
// If no valid mnemonic is given, a random one will be used instead.
func InitializeNodeValidatorFilesFromMnemonic(config *cfg.Config, mnemonic string) (string, cryptotypes.PubKey, error) {
	if len(mnemonic) > 0 && !bip39.IsMnemonicValid(mnemonic) {
		return "", nil, errors.New("invalid mnemonic")
	}
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return "", nil, err
	}

	nodeID := string(nodeKey.ID())

	pvKeyFile := config.PrivValidatorKeyFile()
	if err = os.MkdirAll(filepath.Dir(pvKeyFile), 0o777); err != nil {
		return "", nil, fmt.Errorf("could not create directory %q: %w", filepath.Dir(pvKeyFile), err)
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err = os.MkdirAll(filepath.Dir(pvStateFile), 0o777); err != nil {
		return "", nil, fmt.Errorf("could not create directory %q: %w", filepath.Dir(pvStateFile), err)
	}

	// Check if the localnet priv_validator.json file exists
	// and copy it to the pvKeyFile location if it does.
	localKeyFile := CheckLocalnetValidatorKey(config.Moniker)
	if localKeyFile != "" {
		var data []byte
		data, err = os.ReadFile(localKeyFile)
		if err != nil {
			return "", nil, fmt.Errorf("could not read localnet priv_validator.json file: %w", err)
		}

		if err = os.WriteFile(pvKeyFile, data, 0o600); err != nil {
			return "", nil, fmt.Errorf("could not write localnet priv_validator.json file: %w", err)
		}
	}

	var filePV *privval.FilePV
	if len(mnemonic) == 0 {
		filePV = privval.LoadOrGenFilePV(pvKeyFile, pvStateFile)
	} else {
		privKey := tmed25519.GenPrivKeyFromSecret([]byte(mnemonic))
		filePV = privval.NewFilePV(privKey, pvKeyFile, pvStateFile)
		filePV.Save()
	}

	tmValPubKey, err := filePV.GetPubKey()
	if err != nil {
		return "", nil, err
	}

	var valPubKey cryptotypes.PubKey
	valPubKey, err = cryptocodec.FromCmtPubKeyInterface(tmValPubKey)
	if err != nil {
		return "", nil, err
	}

	return nodeID, valPubKey, nil
}

// CheckLocalnetValidatorKey checks if a private validator key exists for the given moniker
// in the .localnet directory and returns its path if found, or nil if not.
func CheckLocalnetValidatorKey(moniker string) string {
	privValidatorFile := filepath.Join(".localnet", fmt.Sprintf("%s_priv_validator_key.json", moniker))
	if _, err := os.Stat(privValidatorFile); os.IsNotExist(err) {
		return ""
	}

	return privValidatorFile
}

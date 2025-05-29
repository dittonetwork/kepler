package cli

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/spf13/cobra"
)

type PrivValidatorKey struct {
	Address string `json:"address"`
	PubKey  struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"pub_key"`
	PrivKey struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"priv_key"`
}

func GenPrivValKeyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gen-priv-val-key [hex-private-key] [home-dir]",
		Short: "Generate a private validator key json file from ed25519 hex private key",
		Long: `Generate a private validator key json file from ed25519 hex private key.
The first argument is the hex-encoded 32-byte ed25519 private key.
The second argument is the path to the home directory (default: ~/.kepler)`,
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			hexPrivKey := args[0]
			privKeyBytes, err := hex.DecodeString(hexPrivKey)
			if err != nil {
				return fmt.Errorf("failed to decode hex private key: %w", err)
			}

			if len(privKeyBytes) != 32 {
				return fmt.Errorf("invalid private key length: expected 32 bytes, got %d", len(privKeyBytes))
			}

			// Create a full ed25519 private key (64 bytes) from the 32-byte seed
			privKey := ed25519.GenPrivKeyFromSecret(privKeyBytes)
			pubKey := privKey.PubKey()

			key := PrivValidatorKey{
				Address: fmt.Sprintf("%X", pubKey.Address()),
				PubKey: struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				}{
					Type:  "tendermint/PubKeyEd25519",
					Value: base64.StdEncoding.EncodeToString(pubKey.Bytes()),
				},
				PrivKey: struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				}{
					Type:  "tendermint/PrivKeyEd25519",
					Value: base64.StdEncoding.EncodeToString(privKey.Bytes()),
				},
			}

			// Get home directory
			var homeDir string
			if len(args) > 1 {
				homeDir = args[1]
			} else {
				homeDir, err = os.UserHomeDir()
				if err != nil {
					return fmt.Errorf("failed to get user home directory: %w", err)
				}
				homeDir = filepath.Join(homeDir, ".kepler")
			}

			// Create config directory if it doesn't exist
			configDir := filepath.Join(homeDir, "config")
			if err := os.MkdirAll(configDir, 0o755); err != nil {
				return fmt.Errorf("failed to create config directory: %w", err)
			}

			// Write the key file
			keyFile := filepath.Join(configDir, "priv_validator_key.json")
			keyJSON, err := json.MarshalIndent(key, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal key: %w", err)
			}

			if err := os.WriteFile(keyFile, keyJSON, 0o600); err != nil {
				return fmt.Errorf("failed to write key file: %w", err)
			}

			fmt.Printf("Generated priv_validator_key.json at %s\n", keyFile)
			return nil
		},
	}
}

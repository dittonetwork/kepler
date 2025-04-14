package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/dittonetwork/kepler/x/genutil"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/spf13/cobra"
)

func AddBulkGenesisOperatorCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bulk-add-genesis-operators [/file/path.json]",
		Short: "Bulk add genesis operators from restaking protocols to genesis.json",
		Args:  cobra.ExactArgs(1),
		Example: `bulk-add-genesis-operators operators.json
where operators.json is:
[
	{
		"operator_address": "0x0",
		"consensus_pubkey": {
			"@type":"/cosmos.crypto.ed25519.PubKey",
			"key":"byefX/uKpgTsyrcAZKrmYYoFiXG0tmTOOaJFziO3D+E="
		},
		"voting_power": 14388,
		"protocol": "Symbiotic",
		"is_emergency": true,
		"status": "bonded"
	}
]
`,
		Long: `Add genesis operators in bulk to genesis.json.
The provided operators must have pre-created Cosmos accounts using the same
ECDSA secp256k1 key as their EVM operator address.
This command reads a JSON file containing operator details and adds them to the genesis file.
Each operator entry in the input file should specify the EVM operator address and corresponding consensus public key.
The command validates each entry before adding it to ensure proper configuration for the network launch.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)

			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			f, err := os.Open(args[0])
			if err != nil {
				return fmt.Errorf("failed to open genesis file: %w", err)
			}
			defer f.Close()

			var jsonData json.RawMessage
			if err = json.NewDecoder(f).Decode(&jsonData); err != nil {
				return fmt.Errorf("failed to decode genesis file: %w", err)
			}

			var genesis restakingtypes.Validators
			clientCtx.Codec.MustUnmarshalJSON(jsonData, &genesis)

			return genutil.AddGenesisOperators(clientCtx.Codec, genesis.Validators, config.GenesisFile())
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

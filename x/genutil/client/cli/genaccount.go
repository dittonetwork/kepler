package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/dittonetwork/kepler/x/genutil"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/address"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	flagAppendMode = "append"
	flagModuleName = "module-name"
)

// getAddressAndPubKeyFromKeyOrAddress retrieves address and public key from address string or key name.
func getAddressAndPubKeyFromKeyOrAddress(
	clientCtx client.Context,
	addressOrKey string,
	addressCodec address.Codec,
	inBuf *bufio.Reader,
	keyringBackend string,
) (sdk.AccAddress, cryptotypes.PubKey, error) {
	// Try to interpret as address
	addr, err := addressCodec.StringToBytes(addressOrKey)
	if err == nil {
		return addr, nil, nil
	}

	// If not an address, try to get key from keyring
	var kr keyring.Keyring
	if keyringBackend != "" && clientCtx.Keyring == nil {
		kr, err = keyring.New(sdk.KeyringServiceName(), keyringBackend,
			clientCtx.HomeDir, inBuf, clientCtx.Codec)
		if err != nil {
			return nil, nil, err
		}
	} else {
		kr = clientCtx.Keyring
	}

	k, err := kr.Key(addressOrKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get address from Keyring: %w", err)
	}

	var pubKey cryptotypes.PubKey
	pubKey, err = k.GetPubKey()
	if err != nil {
		return nil, nil, err
	}

	addr, err = k.GetAddress()
	if err != nil {
		return nil, nil, err
	}

	return addr, pubKey, nil
}

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
// This command is provided as a default, applications are expected to provide their own command
// if custom genesis accounts are needed.
func AddGenesisAccountCmd(defaultNodeHome string, addressCodec address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account [address_or_key_name] [coin][,[coin]]",
		Short: "Add a genesis account to genesis.json",
		Long: `Add a genesis account to genesis.json. The provided account must specify
the account address or key name and a list of initial coins. If a key name is given,
the address will be looked up in the local Keybase. The list of initial tokens must
contain valid denominations. Accounts may optionally be supplied with vesting parameters.
`,
		Args: cobra.ExactArgs(2), //nolint:mnd // just args
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)

			addr, _, err := getAddressAndPubKeyFromKeyOrAddress(
				clientCtx,
				args[0],
				addressCodec,
				inBuf,
				keyringBackend,
			)
			if err != nil {
				return err
			}

			appendflag, _ := cmd.Flags().GetBool(flagAppendMode)
			moduleNameStr, _ := cmd.Flags().GetString(flagModuleName)

			return genutil.AddGenesisAccount(clientCtx.Codec, addr, appendflag,
				config.GenesisFile(), args[1], moduleNameStr)
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend,
		"Select keyring's backend (os|file|kwallet|pass|test)")
	cmd.Flags().Bool(flagAppendMode, false,
		"append the coins to an account already in the genesis.json file")
	cmd.Flags().String(flagModuleName, "", "module account name")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// AddBulkGenesisAccountCmd returns bulk-add-genesis-account cobra Command.
// This command is provided as a default, applications are expected to provide their own command
// if custom genesis accounts are needed.
func AddBulkGenesisAccountCmd(defaultNodeHome string, addressCodec address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bulk-add-genesis-account [/file/path.json]",
		Short: "Bulk add genesis accounts to genesis.json",
		Example: `bulk-add-genesis-account accounts.json
where accounts.json is:
[
   {
       "address": "cosmos139f7kncmglres2nf3h4hc4tade85ekfr8sulz5",
       "coins": [
           { "denom": "umuon", "amount": "100000000" },
       ]
   },
   {
       "address": "cosmos1e0jnq2sun3dzjh8p2xq95kk0expwmd7shwjpfg",
       "coins": [
           { "denom": "umuon", "amount": "500000000" }
       ],
   }
]
`,
		Long: `Add genesis accounts in bulk to genesis.json. The provided account must specify
the account address and a list of initial coins. The list of initial tokens must
contain valid denominations. Accounts may optionally be supplied with vesting parameters.
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			f, err := os.Open(args[0])
			if err != nil {
				return fmt.Errorf("failed to open file: %w", err)
			}
			defer f.Close()

			var accounts []genutil.GenesisAccount
			if err = json.NewDecoder(f).Decode(&accounts); err != nil {
				return fmt.Errorf("failed to decode JSON: %w", err)
			}

			appendflag, _ := cmd.Flags().GetBool(flagAppendMode)

			return genutil.AddGenesisAccounts(clientCtx.Codec, addressCodec, accounts, appendflag, config.GenesisFile())
		},
	}

	cmd.Flags().Bool(flagAppendMode, false, "append the coins to an account already in the genesis.json file")
	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

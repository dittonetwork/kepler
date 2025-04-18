package cli

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"cosmossdk.io/core/address"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/dittonetwork/kepler/x/genutil"
	"github.com/dittonetwork/kepler/x/genutil/types"
	"github.com/dittonetwork/kepler/x/restaking/client/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// GenTxCmd generates a restaking protocol operator. !!!!Intended for internal use only.
//
//nolint:gocognit,funlen // no matter
func GenTxCmd(defaultNodeHome string, accAddressCodec address.Codec) *cobra.Command {
	ipDefault, _ := server.ExternalIP()
	fsCreateValidator := cli.CreateValidatorMsgFlagSet(ipDefault)

	cmd := &cobra.Command{
		Use:   "gentx [key_name] [amount]",
		Short: "Generates a restaking protocol operator. !!!!Intended for internal use only.",
		Args:  cobra.ExactArgs(2), //nolint:mnd // no matter
		Long: `This command generates an operator from the restaking protocol.

> [!WARNING]
It is not intended for manual use.
The command exists solely for backward compatibility during development using ignite chain serve.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(serverCtx.Config)
			if err != nil {
				return errors.Wrap(err, "failed to initialize node validator files")
			}

			// read --nodeID, if empty take it from priv_validator.json
			if nodeIDString, _ := cmd.Flags().GetString(cli.FlagNodeID); nodeIDString != "" {
				nodeID = nodeIDString
			}

			// read --pubkey, if empty take it from priv_validator.json
			if pkStr, _ := cmd.Flags().GetString(cli.FlagPubKey); pkStr != "" {
				if err = clientCtx.Codec.UnmarshalInterfaceJSON([]byte(pkStr), &valPubKey); err != nil {
					return errors.Wrap(err, "failed to unmarshal validator public key")
				}
			}

			appGenesis, err := types.AppGenesisFromFile(config.GenesisFile())
			if err != nil {
				return errors.Wrapf(err, "failed to read genesis doc file %s", config.GenesisFile())
			}

			var genesisState map[string]json.RawMessage
			if err = json.Unmarshal(appGenesis.AppState, &genesisState); err != nil {
				return errors.Wrap(err, "failed to unmarshal genesis state")
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())

			name := args[0]
			key, err := clientCtx.Keyring.Key(name)
			if err != nil {
				return errors.Wrapf(err, "failed to fetch '%s' from the keyring", name)
			}

			moniker := config.Moniker
			if m, _ := cmd.Flags().GetString(cli.FlagMoniker); m != "" {
				moniker = m
			}

			// set flags for creating a gentx
			createValCfg, err := cli.PrepareConfigForTxBondValidator(
				cmd.Flags(), moniker, nodeID, appGenesis.ChainID,
			)
			if err != nil {
				return errors.Wrap(err, "error creating configuration to create validator msg")
			}

			txFactory, err := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			pub, err := key.GetAddress()
			if err != nil {
				return errors.Wrap(err, "failed to get validator address")
			}
			clientCtx = clientCtx.WithInput(inBuf).WithFromAddress(pub)

			txBuilder, msg, err := cli.BuildBondValidatorMsg(
				clientCtx,
				createValCfg,
				txFactory,
				true,
				accAddressCodec,
			)
			if err != nil {
				return errors.Wrap(err, "failed to build bond validator message")
			}

			if key.GetType() == keyring.TypeOffline || key.GetType() == keyring.TypeMulti {
				cmd.PrintErrln("Offline key passed in. Use `tx sign` command to sign.")
				return txBuilder.PrintUnsignedTx(clientCtx, msg)
			}

			// write the unsigned transaction to the buffer
			w := bytes.NewBuffer([]byte{})
			clientCtx = clientCtx.WithOutput(w)

			if m, ok := msg.(sdk.HasValidateBasic); ok {
				if err = m.ValidateBasic(); err != nil {
					return err
				}
			}

			if err = txBuilder.PrintUnsignedTx(clientCtx, msg); err != nil {
				return errors.Wrap(err, "failed to print unsigned std tx")
			}

			// read the transaction
			stdTx, err := readUnsignedGenTxFile(clientCtx, w)
			if err != nil {
				return errors.Wrap(err, "failed to read unsigned gen tx file")
			}

			txb, err := clientCtx.TxConfig.WrapTxBuilder(stdTx)
			if err != nil {
				return fmt.Errorf("error creating tx builder: %w", err)
			}

			err = authclient.SignTx(txFactory, clientCtx, name, txb, true, true)
			if err != nil {
				return errors.Wrap(err, "failed to sign std tx")
			}

			outputDocument, _ := cmd.Flags().GetString(flags.FlagOutputDocument)
			if outputDocument == "" {
				outputDocument, err = makeOutputFilepath(config.RootDir, nodeID)
				if err != nil {
					return errors.Wrap(err, "failed to create output file path")
				}
			}

			if err = writeSignedGenTx(clientCtx, outputDocument, stdTx); err != nil {
				return errors.Wrap(err, "failed to write signed gen tx")
			}

			cmd.PrintErrf("Genesis transaction written to %q\n", outputDocument)
			return nil
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagOutputDocument, "",
		"Write the genesis transaction JSON document to the given file instead of the default location")
	cmd.Flags().AddFlagSet(fsCreateValidator)
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.Flags().MarkHidden(flags.FlagOutput) // signing makes sense to output only json

	return cmd
}

func readUnsignedGenTxFile(clientCtx client.Context, r io.Reader) (sdk.Tx, error) {
	bz, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	aTx, err := clientCtx.TxConfig.TxJSONDecoder()(bz)
	if err != nil {
		return nil, err
	}

	return aTx, err
}

func writeSignedGenTx(clientCtx client.Context, outputDocument string, tx sdk.Tx) error {
	outputFile, err := os.OpenFile(outputDocument, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	jsonData, err := clientCtx.TxConfig.TxJSONEncoder()(tx)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(outputFile, "%s\n", jsonData)

	return err
}

func makeOutputFilepath(rootDir, nodeID string) (string, error) {
	writePath := filepath.Join(rootDir, "config", "gentx")
	if err := os.MkdirAll(writePath, 0o700); err != nil {
		return "", fmt.Errorf("could not create directory %q: %w", writePath, err)
	}

	return filepath.Join(writePath, fmt.Sprintf("gentx-%v.json", nodeID)), nil
}

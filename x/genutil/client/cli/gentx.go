package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/dittonetwork/kepler/x/genutil"
	"github.com/dittonetwork/kepler/x/genutil/types"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// GenTxCmd generates a restaking protocol operator. !!!!Intended for internal use only.
//
//nolint:gocognit,funlen // no matter
func GenTxCmd(defaultNodeHome string) *cobra.Command {
	ipDefault, _ := server.ExternalIP()
	fsCreateValidator, _ := cli.CreateValidatorMsgFlagSet(ipDefault)

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
			cdc := clientCtx.Codec

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
			createValCfg, err := cli.PrepareConfigForTxCreateValidator(
				cmd.Flags(), moniker, nodeID, appGenesis.ChainID, valPubKey,
			)
			if err != nil {
				return errors.Wrap(err, "error creating configuration to create validator msg")
			}

			amount := args[1]
			coins, err := sdk.ParseCoinsNormalized(amount)
			if err != nil {
				return errors.Wrap(err, "failed to parse coins")
			}

			outputDocument, _ := cmd.Flags().GetString(flags.FlagOutputDocument)
			if outputDocument == "" {
				outputDocument, err = makeOutputFilepath(config.RootDir, nodeID)
				if err != nil {
					return errors.Wrap(err, "failed to create output file path")
				}
			}

			var valPkAny *codectypes.Any
			if createValCfg.PubKey != nil {
				if valPkAny, err = codectypes.NewAnyWithValue(createValCfg.PubKey); err != nil {
					return errors.Wrap(err, "failed to create any from pub key")
				}
			}

			addr, err := key.GetAddress()
			if err != nil {
				return errors.Wrap(err, "failed to get address")
			}

			votingPower := coins.AmountOf("power")

			operators := []restakingtypes.Validator{
				{
					OperatorAddress: addr.String(),
					ConsensusPubkey: valPkAny,
					IsEmergency:     true,
					Status:          restakingtypes.Bonded,
					VotingPower:     votingPower.Int64(),
					Protocol:        restakingtypes.Ditto,
				},
			}

			if err = genutil.AddGenesisOperators(cdc, operators, config.GenesisFile()); err != nil {
				return errors.Wrap(err, "failed to add genesis operators")
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

func makeOutputFilepath(rootDir, nodeID string) (string, error) {
	writePath := filepath.Join(rootDir, "config", "gentx")
	if err := os.MkdirAll(writePath, 0o700); err != nil {
		return "", fmt.Errorf("could not create directory %q: %w", writePath, err)
	}

	return filepath.Join(writePath, fmt.Sprintf("gentx-%v.json", nodeID)), nil
}

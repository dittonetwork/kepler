package cli

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/dittonetwork/kepler/x/genutil"
	genutiltypes "github.com/dittonetwork/kepler/x/genutil/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const (
	suggestionMinimumDistance = 2
)

// Commands adds core sdk's sub-commands into genesis command.
func Commands(txConfig client.TxConfig, moduleBasics module.BasicManager, defaultNodeHome string) *cobra.Command {
	return CommandsWithCustomMigrationMap(txConfig, moduleBasics, defaultNodeHome)
}

// CommandsWithCustomMigrationMap adds core sdk's sub-commands into genesis command with custom migration map.
// This custom migration map can be used by the application to add its own migration map.
func CommandsWithCustomMigrationMap(
	txConfig client.TxConfig, moduleBasics module.BasicManager, defaultNodeHome string,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "genesis",
		Short:                      "Application's genesis-related subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: suggestionMinimumDistance,
		RunE:                       client.ValidateCmd,
	}
	genTxModule, ok := moduleBasics[genutiltypes.ModuleName].(genutil.AppModuleBasic)
	if !ok {
		panic("genutil module is not registered")
	}

	cmd.AddCommand(
		AddGenesisAccountCmd(defaultNodeHome, txConfig.SigningContext().AddressCodec()),
		AddBulkGenesisAccountCmd(defaultNodeHome, txConfig.SigningContext().AddressCodec()),
		AddBulkGenesisOperatorCmd(defaultNodeHome),
		GenTxCmd(defaultNodeHome, txConfig.SigningContext().ValidatorAddressCodec()),
		CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, defaultNodeHome, genTxModule.GenTxValidator),
		ValidateGenesisCmd(moduleBasics),
	)

	return cmd
}

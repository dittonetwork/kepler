package cli

import (
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

	cmd.AddCommand(
		AddGenesisAccountCmd(defaultNodeHome, txConfig.SigningContext().AddressCodec()),
		AddBulkGenesisAccountCmd(defaultNodeHome, txConfig.SigningContext().AddressCodec()),
		AddBulkGenesisOperatorCmd(defaultNodeHome),
		ValidateGenesisCmd(moduleBasics),
	)

	return cmd
}

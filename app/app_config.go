package app

import (
	epochsmodulev1 "github.com/dittonetwork/kepler/api/kepler/epochs/module"
	_ "github.com/dittonetwork/kepler/x/epochs/module" // import for side effects
	epochstypes "github.com/dittonetwork/kepler/x/epochs/types"
	genutiltypes "github.com/dittonetwork/kepler/x/genutil/types"

	committeemodulev1 "github.com/dittonetwork/kepler/api/kepler/committee/module"
	_ "github.com/dittonetwork/kepler/x/committee/module" // import for side-effects
	committeemoduletypes "github.com/dittonetwork/kepler/x/committee/types"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	"cosmossdk.io/core/appconfig"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	restakingmodulev1 "github.com/dittonetwork/kepler/api/kepler/restaking/module"
	_ "github.com/dittonetwork/kepler/x/restaking/module" // import for side-effects
	restakingmoduletypes "github.com/dittonetwork/kepler/x/restaking/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport
)

var (
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.

	genesisModuleOrder = []string{
		// cosmos-sdk modules
		authtypes.ModuleName,
		banktypes.ModuleName,
		restakingmoduletypes.ModuleName,
		genutiltypes.ModuleName,
		// chain modules
		epochstypes.ModuleName,
		committeemoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	}

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	beginBlockers = []string{
		// cosmos sdk modules
		// chain modules
		epochstypes.ModuleName,
		committeemoduletypes.ModuleName,
		restakingmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	}

	endBlockers = []string{
		// cosmos sdk modules
		// chain modules
		epochstypes.ModuleName,
		committeemoduletypes.ModuleName,
		restakingmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers
	}

	preBlockers = []string{
		upgradetypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/preBlockers
	}

	// module account permissions
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: committeemoduletypes.ModuleName, Permissions: []string{authtypes.Minter}},
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
	}

	// appConfig application configuration (used by depinject)
	appConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName:       Name,
					PreBlockers:   preBlockers,
					BeginBlockers: beginBlockers,
					EndBlockers:   endBlockers,
					InitGenesis:   genesisModuleOrder,
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					// When ExportGenesis is not specified, the export genesis module order
					// is equal to the init genesis order
					// ExportGenesis: genesisModuleOrder,
					// Uncomment if you want to set a custom migration order here.
					// OrderMigrations: nil,
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             AccountAddressPrefix,
					ModuleAccountPermissions: moduleAccPerms,
					// By default modules authority is the governance module. This is configurable with the following:
					// Authority: "group", // A custom module authority can be set using a module name
					// Authority: "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv", // or a specific address
				}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   epochstypes.ModuleName,
				Config: appconfig.WrapAny(&epochsmodulev1.Module{}),
			},
			{
				Name: committeemoduletypes.ModuleName,
				Config: appconfig.WrapAny(&committeemodulev1.Module{
					EpochId: MainEpochID,
				}),
			},
			{
				Name: restakingmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&restakingmodulev1.Module{
					// By default, main epoch id is "hour". This is configurable with the following:
					// MainEpochId: "main"
					MainEpochId: MainEpochID,

					// By default, modules authority is the committee module. This is configurable with the following:
					// Authority: "ditto1kprsj8y2x9d7pyzvfnurkuygsttnjy90fv6gkk", // specific address (alice for exam)
					// Authority: "abracadabra" // or a custom module authority can be set using a module name
				}),
			},
			// this line is used by starport scaffolding # stargate/app/moduleConfig
		},
	})
)

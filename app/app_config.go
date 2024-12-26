package app

import (
	accountsmodulev1 "cosmossdk.io/api/cosmos/accounts/module/v1"
	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	epochsmodulev1 "cosmossdk.io/api/cosmos/epochs/module/v1"
	poolmodulev1 "cosmossdk.io/api/cosmos/protocolpool/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"

	"cosmossdk.io/depinject/appconfig"
	"cosmossdk.io/x/accounts"

	_ "cosmossdk.io/x/bank"
	banktypes "cosmossdk.io/x/bank/types"

	_ "cosmossdk.io/x/consensus"
	consensustypes "cosmossdk.io/x/consensus/types"

	_ "cosmossdk.io/x/epochs"
	epochstypes "cosmossdk.io/x/epochs/types"

	minttypes "cosmossdk.io/x/mint/types"

	_ "cosmossdk.io/x/protocolpool"
	pooltypes "cosmossdk.io/x/protocolpool/types"

	_ "kepler/x/genutil/module"
	genutilmoduletypes "kepler/x/genutil/types"

	_ "kepler/x/horizon/module"
	horizonmoduletypes "kepler/x/horizon/types"

	_ "kepler/x/staking/module"
	stakingmoduletypes "kepler/x/staking/types"

	_ "github.com/cosmos/cosmos-sdk/testutil/x/counter"
	_ "github.com/cosmos/cosmos-sdk/x/auth"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

var (
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: pooltypes.ModuleName},
		{Account: pooltypes.StreamAccount},
		{Account: pooltypes.ProtocolPoolDistrAccount},
		{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
		// pooltypes.ModuleName
	}

	// application configuration (used by depinject)
	appConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName: Name,
					// NOTE: upgrade module is required to be prioritized
					PreBlockers: []string{
						// this line is used by starport scaffolding # stargate/app/preBlockers
					},
					// During begin block slashing happens after distr.BeginBlocker so that
					// there is nothing left over in the validator fee pool, so as to keep the
					// CanWithdrawInvariant invariant.
					// NOTE: staking module is required if HistoricalEntries param > 0
					BeginBlockers: []string{
						pooltypes.ModuleName,
						// chain modules
						horizonmoduletypes.ModuleName,
						epochstypes.ModuleName,
						stakingmoduletypes.ModuleName,
						// this line is used by starport scaffolding # stargate/app/beginBlockers
					},
					EndBlockers: []string{
						pooltypes.ModuleName,
						// chain modules
						horizonmoduletypes.ModuleName,
						stakingmoduletypes.ModuleName,
						// this line is used by starport scaffolding # stargate/app/endBlockers
					},
					// The following is mostly only needed when ModuleName != StoreKey name.
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
						{
							ModuleName: accounts.ModuleName,
							KvStoreKey: accounts.StoreKey,
						},
					},
					// NOTE: The genutils module must occur after staking so that pools are
					// properly initialized with tokens from genesis accounts.
					// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
					InitGenesis: []string{
						consensustypes.ModuleName,
						accounts.ModuleName,
						authtypes.ModuleName,
						banktypes.ModuleName,
						pooltypes.ModuleName,
						epochstypes.ModuleName,
						// chain modules
						genutilmoduletypes.ModuleName,
						horizonmoduletypes.ModuleName,
						stakingmoduletypes.ModuleName,
						// this line is used by starport scaffolding # stargate/app/initGenesis
					},
					// SkipStoreKeys is an optional list of store keys to skip when constructing the
					// module's keeper. This is useful when a module does not have a store key.
					SkipStoreKeys: []string{
						"tx",
					},
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
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			{
				Name:   pooltypes.ModuleName,
				Config: appconfig.WrapAny(&poolmodulev1.Module{}),
			},
			{
				Name:   accounts.ModuleName,
				Config: appconfig.WrapAny(&accountsmodulev1.Module{}),
			},
			{
				Name:   epochstypes.ModuleName,
				Config: appconfig.WrapAny(&epochsmodulev1.Module{}),
			},
			{
				Name:   genutilmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmoduletypes.Module{}),
			},
			{
				Name:   horizonmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&horizonmoduletypes.Module{}),
			},
			{
				Name: stakingmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&stakingmoduletypes.Module{
					Bech32PrefixValidator: AccountAddressPrefix + "valoper",
					Bech32PrefixConsensus: AccountAddressPrefix + "valcons",
				}),
			},
			// this line is used by starport scaffolding # stargate/app/moduleConfig
		},
	})
)

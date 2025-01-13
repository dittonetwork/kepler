package dittostaking

import (
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	bankkeeper "cosmossdk.io/x/bank/keeper"
	consensuskeeper "cosmossdk.io/x/consensus/keeper"
	stakingkeeper "cosmossdk.io/x/staking/keeper"

	"kepler/x/dittostaking/keeper"
	"kepler/x/dittostaking/types"
)

var _ depinject.OnePerModuleType = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModule) IsOnePerModuleType() {}

func init() {
	appconfig.Register(
		&types.Module{},
		appconfig.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Config          *types.Module
	Environment     appmodule.Environment
	Cdc             codec.Codec
	BankKeeper      bankkeeper.Keeper
	ConsensusKeeper consensuskeeper.Keeper
	StakingKeeper   *stakingkeeper.Keeper
	AddressCodec    address.Codec
}

type ModuleOutputs struct {
	depinject.Out

	DittoStakingKeeper keeper.Keeper
	Module             appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(types.GovModuleName)

	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(
		in.Environment,
		in.Cdc,
		in.BankKeeper,
		in.ConsensusKeeper,
		in.StakingKeeper,
		authority,
		in.AddressCodec,
	)

	m := NewAppModule(in.Cdc, k)

	return ModuleOutputs{DittoStakingKeeper: k, Module: m}
}

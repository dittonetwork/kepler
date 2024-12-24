package genutil

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"kepler/x/genutil/types"
)

var _ depinject.OnePerModuleType = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

func init() {
	appconfig.RegisterModule(&types.Module{},
		appconfig.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	AccountKeeper  types.AccountKeeper
	StakingKeeper  types.StakingKeeper
	Config         client.TxConfig
	Cdc            codec.Codec
	DeliverTx      TxHandler              `optional:"true"` // Only used in server v0 applications
	GenTxValidator types.MessageValidator `optional:"true"`
}

func ProvideModule(in ModuleInputs) appmodule.AppModule {
	if in.GenTxValidator == nil {
		in.GenTxValidator = types.DefaultMessageValidator
	}

	return NewAppModule(in.Cdc, in.AccountKeeper, in.StakingKeeper, in.DeliverTx, in.Config, in.GenTxValidator)
}

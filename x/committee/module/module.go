package committee

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkruntime "github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	epochstypes "github.com/dittonetwork/kepler/x/epochs/types"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	// this line is used by starport scaffolding # 1.

	modulev1 "github.com/dittonetwork/kepler/api/kepler/committee/module"
	"github.com/dittonetwork/kepler/x/committee/keeper"
	"github.com/dittonetwork/kepler/x/committee/repository"
	"github.com/dittonetwork/kepler/x/committee/types"
)

var (
	_ module.AppModuleBasic      = (*AppModule)(nil)
	_ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasInvariants       = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)

	_ appmodule.AppModule       = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker = (*AppModule)(nil)
	_ appmodule.HasEndBlocker   = (*AppModule)(nil)
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the
// independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used
// to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore.
func (AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage.
// The default GenesisState need to be defined by the module developer and is primarily used for testing.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement.
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper

	epochMainID string
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	epochMainID string,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		epochMainID:    epochMainID,
	}
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServerImpl(am.keeper))
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value,
// the InvariantRegistry triggers appropriate logic (most often the chain will be halted).
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	am.keeper.InitGenesis(ctx, genState)
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module.
// It should be incremented on each consensus-breaking change introduced by the module.
// To avoid wrong/empty versions, the initial version should be set to 1.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block.
// The begin block implementation is optional.
func (am AppModule) BeginBlock(_ context.Context) error {
	return nil
}

// EndBlock contains the logic that is automatically triggered at the end of each block.
// The end block implementation is optional.
func (am AppModule) EndBlock(_ context.Context) error {
	return nil
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// ----------------------------------------------------------------------------
// App Wiring Setup
// ----------------------------------------------------------------------------

//nolint:gochecknoinits // generated
func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Router baseapp.MessageRouter

	StoreService store.KVStoreService
	Cdc          codec.Codec
	Amino        *codec.LegacyAmino
	Config       *modulev1.Module
	Logger       log.Logger

	ValidatorAddressCodec sdkruntime.ValidatorAddressCodec

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper

	RestakingKeeper types.RestakingKeeper
}

type ModuleOutputs struct {
	depinject.Out

	CommitteeKeeper keeper.Keeper
	Module          appmodule.AppModule
	EpochsHooks     epochstypes.EpochHooksWrapper
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	mainEpochID := "hour"
	if in.Config.EpochId != "" {
		mainEpochID = in.Config.EpochId
	}

	repo := repository.NewCommitteeRepository(in.StoreService, in.Cdc)

	k := keeper.NewKeeper(
		authority.String(),
		in.AccountKeeper,
		in.BankKeeper,
		in.RestakingKeeper,
		repo,
		in.Logger,
		in.Router,
		in.Amino,
		in.Cdc,
		mainEpochID,
		in.ValidatorAddressCodec,
	)
	m := NewAppModule(
		in.Cdc,
		k,
		in.AccountKeeper,
		in.BankKeeper,
		in.Config.EpochId,
	)

	return ModuleOutputs{
		CommitteeKeeper: k,
		Module:          m,
		EpochsHooks:     epochstypes.EpochHooksWrapper{EpochHooks: k.EpochsHooks()},
	}
}

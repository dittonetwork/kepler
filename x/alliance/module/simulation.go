package alliance

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"kepler/testutil/sample"
	alliancesimulation "kepler/x/alliance/simulation"
	"kepler/x/alliance/types"
)

// avoid unused import issue
var (
	_ = alliancesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgAddEntropy = "op_weight_msg_add_entropy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddEntropy int = 100

	opWeightMsgCreateSharedEntropy = "op_weight_msg_shared_entropy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSharedEntropy int = 100

	opWeightMsgUpdateSharedEntropy = "op_weight_msg_shared_entropy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSharedEntropy int = 100

	opWeightMsgDeleteSharedEntropy = "op_weight_msg_shared_entropy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSharedEntropy int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	allianceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&allianceGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddEntropy int
	simState.AppParams.GetOrGenerate(opWeightMsgAddEntropy, &weightMsgAddEntropy, nil,
		func(_ *rand.Rand) {
			weightMsgAddEntropy = defaultWeightMsgAddEntropy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddEntropy,
		alliancesimulation.SimulateMsgAddEntropy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSharedEntropy int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSharedEntropy, &weightMsgCreateSharedEntropy, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSharedEntropy = defaultWeightMsgCreateSharedEntropy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSharedEntropy,
		alliancesimulation.SimulateMsgCreateSharedEntropy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSharedEntropy int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSharedEntropy, &weightMsgUpdateSharedEntropy, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSharedEntropy = defaultWeightMsgUpdateSharedEntropy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSharedEntropy,
		alliancesimulation.SimulateMsgUpdateSharedEntropy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSharedEntropy int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSharedEntropy, &weightMsgDeleteSharedEntropy, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSharedEntropy = defaultWeightMsgDeleteSharedEntropy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSharedEntropy,
		alliancesimulation.SimulateMsgDeleteSharedEntropy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddEntropy,
			defaultWeightMsgAddEntropy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				alliancesimulation.SimulateMsgAddEntropy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSharedEntropy,
			defaultWeightMsgCreateSharedEntropy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				alliancesimulation.SimulateMsgCreateSharedEntropy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateSharedEntropy,
			defaultWeightMsgUpdateSharedEntropy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				alliancesimulation.SimulateMsgUpdateSharedEntropy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteSharedEntropy,
			defaultWeightMsgDeleteSharedEntropy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				alliancesimulation.SimulateMsgDeleteSharedEntropy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

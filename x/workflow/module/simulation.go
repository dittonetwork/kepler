package workflow

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"kepler/testutil/sample"
	workflowsimulation "kepler/x/workflow/simulation"
	"kepler/x/workflow/types"
)

// avoid unused import issue.
var (
	_ = workflowsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	//nolint:gosec // TODO: Determine the operation weight value.
	opWeightMsgAddAutomation = "op_weight_msg_add_automation"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgAddAutomation int = 100

	// this line is used by starport scaffolding # simapp/module/const.
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	workflowGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&workflowGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddAutomation int
	simState.AppParams.GetOrGenerate(opWeightMsgAddAutomation, &weightMsgAddAutomation, nil,
		func(_ *rand.Rand) {
			weightMsgAddAutomation = defaultWeightMsgAddAutomation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddAutomation,
		workflowsimulation.SimulateMsgAddAutomation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddAutomation,
			defaultWeightMsgAddAutomation,
			func(_ *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
				workflowsimulation.SimulateMsgAddAutomation(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

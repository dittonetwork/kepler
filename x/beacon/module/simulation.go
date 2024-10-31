package beacon

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"kepler/testutil/sample"
	beaconsimulation "kepler/x/beacon/simulation"
	"kepler/x/beacon/types"
)

// avoid unused import issue
var (
	_ = beaconsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateFinalizedBlockInfo = "op_weight_msg_finalized_block_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFinalizedBlockInfo int = 100

	opWeightMsgUpdateFinalizedBlockInfo = "op_weight_msg_finalized_block_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateFinalizedBlockInfo int = 100

	opWeightMsgDeleteFinalizedBlockInfo = "op_weight_msg_finalized_block_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFinalizedBlockInfo int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	beaconGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&beaconGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateFinalizedBlockInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateFinalizedBlockInfo, &weightMsgCreateFinalizedBlockInfo, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFinalizedBlockInfo = defaultWeightMsgCreateFinalizedBlockInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFinalizedBlockInfo,
		beaconsimulation.SimulateMsgCreateFinalizedBlockInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateFinalizedBlockInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateFinalizedBlockInfo, &weightMsgUpdateFinalizedBlockInfo, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateFinalizedBlockInfo = defaultWeightMsgUpdateFinalizedBlockInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateFinalizedBlockInfo,
		beaconsimulation.SimulateMsgUpdateFinalizedBlockInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFinalizedBlockInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteFinalizedBlockInfo, &weightMsgDeleteFinalizedBlockInfo, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFinalizedBlockInfo = defaultWeightMsgDeleteFinalizedBlockInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFinalizedBlockInfo,
		beaconsimulation.SimulateMsgDeleteFinalizedBlockInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateFinalizedBlockInfo,
			defaultWeightMsgCreateFinalizedBlockInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				beaconsimulation.SimulateMsgCreateFinalizedBlockInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateFinalizedBlockInfo,
			defaultWeightMsgUpdateFinalizedBlockInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				beaconsimulation.SimulateMsgUpdateFinalizedBlockInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteFinalizedBlockInfo,
			defaultWeightMsgDeleteFinalizedBlockInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				beaconsimulation.SimulateMsgDeleteFinalizedBlockInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

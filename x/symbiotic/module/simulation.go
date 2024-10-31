package symbiotic

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"kepler/testutil/sample"
	symbioticsimulation "kepler/x/symbiotic/simulation"
	"kepler/x/symbiotic/types"
)

// avoid unused import issue
var (
	_ = symbioticsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateStakedAmountInfo = "op_weight_msg_staked_amount_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStakedAmountInfo int = 100

	opWeightMsgUpdateStakedAmountInfo = "op_weight_msg_staked_amount_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStakedAmountInfo int = 100

	opWeightMsgDeleteStakedAmountInfo = "op_weight_msg_staked_amount_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteStakedAmountInfo int = 100

	opWeightMsgCreateContractAddress = "op_weight_msg_contract_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContractAddress int = 100

	opWeightMsgUpdateContractAddress = "op_weight_msg_contract_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateContractAddress int = 100

	opWeightMsgDeleteContractAddress = "op_weight_msg_contract_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteContractAddress int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	symbioticGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		StakedAmountInfoList: []types.StakedAmountInfo{
			{
				Creator:         sample.AccAddress(),
				EthereumAddress: "0",
			},
			{
				Creator:         sample.AccAddress(),
				EthereumAddress: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&symbioticGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateStakedAmountInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateStakedAmountInfo, &weightMsgCreateStakedAmountInfo, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStakedAmountInfo = defaultWeightMsgCreateStakedAmountInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStakedAmountInfo,
		symbioticsimulation.SimulateMsgCreateStakedAmountInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStakedAmountInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateStakedAmountInfo, &weightMsgUpdateStakedAmountInfo, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStakedAmountInfo = defaultWeightMsgUpdateStakedAmountInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStakedAmountInfo,
		symbioticsimulation.SimulateMsgUpdateStakedAmountInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteStakedAmountInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteStakedAmountInfo, &weightMsgDeleteStakedAmountInfo, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStakedAmountInfo = defaultWeightMsgDeleteStakedAmountInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteStakedAmountInfo,
		symbioticsimulation.SimulateMsgDeleteStakedAmountInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateContractAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateContractAddress, &weightMsgCreateContractAddress, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContractAddress = defaultWeightMsgCreateContractAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContractAddress,
		symbioticsimulation.SimulateMsgCreateContractAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateContractAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateContractAddress, &weightMsgUpdateContractAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateContractAddress = defaultWeightMsgUpdateContractAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateContractAddress,
		symbioticsimulation.SimulateMsgUpdateContractAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteContractAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteContractAddress, &weightMsgDeleteContractAddress, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteContractAddress = defaultWeightMsgDeleteContractAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteContractAddress,
		symbioticsimulation.SimulateMsgDeleteContractAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateStakedAmountInfo,
			defaultWeightMsgCreateStakedAmountInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				symbioticsimulation.SimulateMsgCreateStakedAmountInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateStakedAmountInfo,
			defaultWeightMsgUpdateStakedAmountInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				symbioticsimulation.SimulateMsgUpdateStakedAmountInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteStakedAmountInfo,
			defaultWeightMsgDeleteStakedAmountInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				symbioticsimulation.SimulateMsgDeleteStakedAmountInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateContractAddress,
			defaultWeightMsgCreateContractAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				symbioticsimulation.SimulateMsgCreateContractAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateContractAddress,
			defaultWeightMsgUpdateContractAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				symbioticsimulation.SimulateMsgUpdateContractAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteContractAddress,
			defaultWeightMsgDeleteContractAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				symbioticsimulation.SimulateMsgDeleteContractAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

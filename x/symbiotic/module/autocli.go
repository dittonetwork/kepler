package symbiotic

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "kepler/api/kepler/symbiotic"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "StakedAmountInfoAll",
					Use:       "list-staked-amount-info",
					Short:     "List all StakedAmountInfo",
				},
				{
					RpcMethod:      "StakedAmountInfo",
					Use:            "show-staked-amount-info [id]",
					Short:          "Shows a StakedAmountInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ethereumAddress"}},
				},
				{
					RpcMethod: "ContractAddress",
					Use:       "show-contract-address",
					Short:     "show contractAddress",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateStakedAmountInfo",
					Use:            "create-staked-amount-info [ethereumAddress] [stakedAmount] [lastUpdateTs]",
					Short:          "Create a new StakedAmountInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ethereumAddress"}, {ProtoField: "stakedAmount"}, {ProtoField: "lastUpdateTs"}},
				},
				{
					RpcMethod:      "UpdateStakedAmountInfo",
					Use:            "update-staked-amount-info [ethereumAddress] [stakedAmount] [lastUpdateTs]",
					Short:          "Update StakedAmountInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ethereumAddress"}, {ProtoField: "stakedAmount"}, {ProtoField: "lastUpdateTs"}},
				},
				{
					RpcMethod:      "DeleteStakedAmountInfo",
					Use:            "delete-staked-amount-info [ethereumAddress]",
					Short:          "Delete StakedAmountInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ethereumAddress"}},
				},
				{
					RpcMethod:      "CreateContractAddress",
					Use:            "create-contract-address [address]",
					Short:          "Create contractAddress",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "UpdateContractAddress",
					Use:            "update-contract-address [address]",
					Short:          "Update contractAddress",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod: "DeleteContractAddress",
					Use:       "delete-contract-address",
					Short:     "Delete contractAddress",
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

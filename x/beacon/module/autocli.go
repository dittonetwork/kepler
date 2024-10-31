package beacon

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "kepler/api/kepler/beacon"
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
					RpcMethod: "FinalizedBlockInfo",
					Use:       "show-finalized-block-info",
					Short:     "show FinalizedBlockInfo",
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
					RpcMethod:      "CreateFinalizedBlockInfo",
					Use:            "create-finalized-block-info [slotNum] [blockTimestamp] [blockNum] [blockHash]",
					Short:          "Create FinalizedBlockInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "slotNum"}, {ProtoField: "blockTimestamp"}, {ProtoField: "blockNum"}, {ProtoField: "blockHash"}},
				},
				{
					RpcMethod:      "UpdateFinalizedBlockInfo",
					Use:            "update-finalized-block-info [slotNum] [blockTimestamp] [blockNum] [blockHash]",
					Short:          "Update FinalizedBlockInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "slotNum"}, {ProtoField: "blockTimestamp"}, {ProtoField: "blockNum"}, {ProtoField: "blockHash"}},
				},
				{
					RpcMethod: "DeleteFinalizedBlockInfo",
					Use:       "delete-finalized-block-info",
					Short:     "Delete FinalizedBlockInfo",
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

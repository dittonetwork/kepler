package alliance

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "kepler/api/kepler/alliance"
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
					RpcMethod: "SharedEntropy",
					Use:       "show-shared-entropy",
					Short:     "show SharedEntropy",
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
					RpcMethod:      "AddEntropy",
					Use:            "add-entropy [entropy]",
					Short:          "Send a AddEntropy tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "entropy"}},
				},
				{
					RpcMethod:      "CreateSharedEntropy",
					Use:            "create-shared-entropy [entropy]",
					Short:          "Create SharedEntropy",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "entropy"}},
				},
				{
					RpcMethod:      "UpdateSharedEntropy",
					Use:            "update-shared-entropy [entropy]",
					Short:          "Update SharedEntropy",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "entropy"}},
				},
				{
					RpcMethod: "DeleteSharedEntropy",
					Use:       "delete-shared-entropy",
					Short:     "Delete SharedEntropy",
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

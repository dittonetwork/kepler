package job

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "kepler/api/kepler/job"
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
					RpcMethod:      "GetJob",
					Use:            "get-job [id]",
					Short:          "Query get-job",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateJob",
					Use:            "create-job [job]",
					Short:          "Send a create-job tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "job"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

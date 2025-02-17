package workflow

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "kepler/api/kepler/workflow"
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
					RpcMethod:      "GetActiveAutomations",
					Use:            "get-active-automations",
					Short:          "Query GetActiveAutomations",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod: "AddAutomation",
					Use:       "add-automation [automation]",
					Short:     "Send a AddAutomation tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "triggers"},
						{ProtoField: "actions"},
						{ProtoField: "expire_at"},
					},
				},
				{
					RpcMethod: "CancelAutomation",
					Use:       "cancel-automation [id]",
					Short:     "Send a CancelAutomation tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

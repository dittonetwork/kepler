package epochs

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/dittonetwork/kepler/api/kepler/epochs"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "EpochInfos",
					Use:       "epoch-infos",
					Short:     "Query running epoch infos",
				},
				{
					RpcMethod:      "CurrentEpoch",
					Use:            "current-epoch",
					Short:          "Query current epoch by specified identifier",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "identifier"}},
				},
			},
		},
	}
}

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
				{
					RpcMethod: "QuorumParams",
					Use:       "show-quorum-params",
					Short:     "show QuorumParams",
				},
				{
					RpcMethod: "AlliancesTimelineAll",
					Use:       "list-alliances-timeline",
					Short:     "List all AlliancesTimeline",
				},
				{
					RpcMethod:      "AlliancesTimeline",
					Use:            "show-alliances-timeline [id]",
					Short:          "Shows a AlliancesTimeline by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
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
				{
					RpcMethod:      "CreateQuorumParams",
					Use:            "create-quorum-params [maxParticipants] [thresholdPercent] [lifetimeInBlocks]",
					Short:          "Create QuorumParams",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "maxParticipants"}, {ProtoField: "thresholdPercent"}, {ProtoField: "lifetimeInBlocks"}},
				},
				{
					RpcMethod:      "UpdateQuorumParams",
					Use:            "update-quorum-params [maxParticipants] [thresholdPercent] [lifetimeInBlocks]",
					Short:          "Update QuorumParams",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "maxParticipants"}, {ProtoField: "thresholdPercent"}, {ProtoField: "lifetimeInBlocks"}},
				},
				{
					RpcMethod: "DeleteQuorumParams",
					Use:       "delete-quorum-params",
					Short:     "Delete QuorumParams",
				},
				{
					RpcMethod:      "CreateAlliancesTimeline",
					Use:            "create-alliances-timeline [participants] [startBlock] [endBlock]",
					Short:          "Create AlliancesTimeline",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "participants"}, {ProtoField: "startBlock"}, {ProtoField: "endBlock"}},
				},
				{
					RpcMethod:      "UpdateAlliancesTimeline",
					Use:            "update-alliances-timeline [id] [participants] [startBlock] [endBlock]",
					Short:          "Update AlliancesTimeline",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "participants"}, {ProtoField: "startBlock"}, {ProtoField: "endBlock"}},
				},
				{
					RpcMethod:      "DeleteAlliancesTimeline",
					Use:            "delete-alliances-timeline [id]",
					Short:          "Delete AlliancesTimeline",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

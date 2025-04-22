package restaking

import (
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"github.com/cosmos/cosmos-sdk/version"

	modulev1 "github.com/dittonetwork/kepler/api/kepler/restaking"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "PendingOperators",
					Use:       "pending-operators",
					Short:     "Get all pending operators",
				},
				{
					RpcMethod: "Validators",
					Use:       "validators",
					Short:     "Get all validators",
					Long:      "Get all validators with pagination and status filter",
					Example: fmt.Sprintf(`$ %s query restaking validators --status bonded --page-offset 2`,
						version.AppName,
					),
				},
				{
					RpcMethod: "OperatorStatus",
					Use:       "operator-status",
					Short:     "Get operator status by EVM address",
					Long:      "Get operator status by EVM address",
					Example: fmt.Sprintf(`$ %s query restaking operator-status --evm-address 0x...`,
						version.AppName,
					),
				},
				// this line is used by ignite scaffolding # autocli/query
				{
					RpcMethod: "NeedValidatorsUpdate",
					Use:       "need-update",
					Short:     "Need validators update",
					Long:      "Check whether the validators set has been updated in the current epoch.",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "BondValidator",
					Use:       "bond-validator",
					Short:     "Bond a validator",
					Long:      "Bond a new validator already initialized on L1",
					Example:   fmt.Sprintf(`$ %s tx restaking bond-validator --from [mykey]`, version.AppName),
				},
				{
					RpcMethod: "UpdateValidatorsSet",
					Use:       "update-validators",
					Short:     "Update validators set",
					Long:      "Update validators set for the current epochs",
				},
			},
		},
	}
}

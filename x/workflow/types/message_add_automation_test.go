package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	// Adjust the import path as needed.
	"github.com/dittonetwork/kepler/x/workflow/types"
)

// validCosmosAddress is a sample valid Cosmos bech32 address.
const validCosmosAddress = "cosmos1ghekyjucln7y67ntx7cf27m9dpuxxemn4c8g4r"

// validEthAddress is a sample valid Ethereum address.
const validEthAddress = "0x1234567890abcdef1234567890abcdef12345678"

// buildValidOnChainCallTrigger constructs a valid on-chain-call trigger.
func buildValidOnChainCallTrigger(chainID string) *types.Trigger {
	return &types.Trigger{
		// Simulate oneof field: OnChainCall
		Trigger: &types.Trigger_OnChainCall{
			OnChainCall: &types.OnChainCallTrigger{
				Contract: validEthAddress,
				ChainId:  chainID,
				MethodAbi: &types.MethodABI{
					Name: "checkCondition",
					Abi:  []byte("{\"name\":\"checkCondition\",\"type\":\"function\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\"}"),
				},
				Args: []string{
					"0xabcdef1234567890abcdef1234567890abcdef12",
					"1000",
				},
			},
		},
	}
}

// buildValidScheduleTrigger returns a valid schedule trigger.
func buildValidScheduleTrigger() *types.Trigger {
	return &types.Trigger{
		Trigger: &types.Trigger_Schedule{
			Schedule: &types.ScheduleTrigger{
				Cron: "0 12 * * 1", // Valid 5-field cron expression.
			},
		},
	}
}

// buildValidCountTrigger returns a valid count trigger.
func buildValidCountTrigger() *types.Trigger {
	return &types.Trigger{
		Trigger: &types.Trigger_Count{
			Count: &types.CountTrigger{
				RepeatCount: 1,
			},
		},
	}
}

func buildValidGasLimitTrigger() *types.Trigger {
	return &types.Trigger{
		Trigger: &types.Trigger_GasLimit{
			GasLimit: &types.GasLimitTrigger{
				MaxFeePerGas:         "1",
				MaxPriorityFeePerGas: "1",
			},
		},
	}
}

func buildValidValidUntilTrigger() *types.Trigger {
	return &types.Trigger{
		Trigger: &types.Trigger_ValidUntil{
			ValidUntil: &types.ValidUntilTrigger{
				Timestamp: time.Now().Add(time.Hour).Unix(),
			},
		},
	}
}

// buildValidUserOp returns a valid user op.
func buildValidUserOp(chainID string) *types.UserOp {
	return &types.UserOp{
		ContractAddress: []byte(validEthAddress),
		ChainId:         chainID,
	}
}

// buildValidMsg creates a valid MsgAddAutomation.
func buildValidMsg() *types.MsgAddAutomation {
	chainID := "1"
	return &types.MsgAddAutomation{
		Creator: validCosmosAddress,
		Triggers: []*types.Trigger{
			buildValidOnChainCallTrigger(chainID),
			buildValidScheduleTrigger(),
			buildValidCountTrigger(),
			buildValidGasLimitTrigger(),
			buildValidValidUntilTrigger(),
		},
		UserOp: buildValidUserOp(chainID),
	}
}

func TestMsgAddAutomation_ValidateBasic_Valid(t *testing.T) {
	msg := buildValidMsg()
	err := msg.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgAddAutomation_ValidateBasic_InvalidCreator(t *testing.T) {
	msg := buildValidMsg()
	msg.Creator = "invalid_address"
	err := msg.ValidateBasic()
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid creator address")
}

func TestValidateValidUntilTrigger_PastExpireAt(t *testing.T) {
	msg := buildValidMsg()
	for _, trig := range msg.Triggers {
		if vut := trig.GetValidUntil(); vut != nil {
			vut.Timestamp = time.Now().Add(-time.Hour).Unix()
		}
	}
	err := msg.ValidateBasic()
	require.Error(t, err)
	require.Contains(t, err.Error(), "valid until time must be in the future")
}

func TestMsgAddAutomation_ValidateBasic_MismatchedChainID(t *testing.T) {
	msg := buildValidMsg()
	// Change the chain id in the on-chain action to mismatch.
	msg.UserOp.ChainId = "137"
	err := msg.ValidateBasic()
	require.Error(t, err)
	require.Contains(t, err.Error(), "triggers and user_op must be on the same chain")
}

func TestValidateCountTriggers_InvalidRepeatCount(t *testing.T) {
	msg := buildValidMsg()
	// Set a count trigger with repeat count 0.
	for _, trig := range msg.Triggers {
		if cnt := trig.GetCount(); cnt != nil {
			cnt.RepeatCount = 0
		}
	}
	err := msg.ValidateCountTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "repeat count must be greater than 0")
}

func TestValidateGasLimitTriggers_InvalidValue(t *testing.T) {
	msg := buildValidMsg()
	// Set a count trigger with repeat count 0.
	for _, trig := range msg.Triggers {
		if cnt := trig.GetGasLimit(); cnt != nil {
			cnt.MaxFeePerGas = "fsdf"
		}
	}
	err := msg.ValidateGasLimitTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid max fee per gas")

	msg = buildValidMsg()
	// Set a count trigger with repeat count 0.
	for _, trig := range msg.Triggers {
		if cnt := trig.GetGasLimit(); cnt != nil {
			cnt.MaxPriorityFeePerGas = "fsdf"
		}
	}
	err = msg.ValidateGasLimitTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid max priority fee per gas")

	msg = buildValidMsg()
	// Set a count trigger with repeat count 0.
	for _, trig := range msg.Triggers {
		if cnt := trig.GetGasLimit(); cnt != nil {
			cnt.MaxPriorityFeePerGas = "-1"
		}
	}
	err = msg.ValidateGasLimitTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "gas limit must be greater than or equal to 1")
}

func TestValidateScheduleTriggers_InvalidCron(t *testing.T) {
	msg := buildValidMsg()
	// Replace schedule trigger with invalid cron expression.
	for i, trig := range msg.Triggers {
		if sch := trig.GetSchedule(); sch != nil {
			sch.Cron = "0 12 * * * *" // 6-field expression, invalid for our validation.
			t.Logf("Trigger %d has invalid cron: %s", i, sch.Cron)
		}
	}
	err := msg.ValidateScheduleTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid cron expression")
}

func TestValidateOnChainCallTriggers_InvalidContract(t *testing.T) {
	msg := buildValidMsg()
	// Set invalid contract address in on-chain call trigger.
	for i, trig := range msg.Triggers {
		if occ := trig.GetOnChainCall(); occ != nil {
			occ.Contract = "123456" // missing "0x", invalid.
			t.Logf("Trigger %d has invalid contract: %s", i, occ.Contract)
		}
	}
	_, err := msg.ValidateOnChainCallTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid contract address")
}

func TestValidateOnChainCallTriggers_MissingMethodABI(t *testing.T) {
	msg := buildValidMsg()
	// Remove method_abi from on-chain call trigger.
	for i, trig := range msg.Triggers {
		if occ := trig.GetOnChainCall(); occ != nil {
			occ.MethodAbi = nil
			t.Logf("Trigger %d missing method_abi", i)
		}
	}
	_, err := msg.ValidateOnChainCallTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "method_abi must be provided")
}

func TestValidateOnChainCallTriggers_MismatchedArgsCount(t *testing.T) {
	msg := buildValidMsg()
	// Remove one argument from on-chain call trigger.
	for i, trig := range msg.Triggers {
		if occ := trig.GetOnChainCall(); occ != nil {
			occ.Args = occ.Args[:1]
			t.Logf("Trigger %d has mismatched args count", i)
		}
	}
	_, err := msg.ValidateOnChainCallTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "number of args")
}

func TestValidateOnChainCallTriggers_InvalidArgType(t *testing.T) {
	msg := buildValidMsg()
	// Change an address argument to an invalid value.
	for i, trig := range msg.Triggers {
		if occ := trig.GetOnChainCall(); occ != nil {
			// For input "address", set an invalid value.
			occ.Args[0] = "invalid_address"
			t.Logf("Trigger %d has invalid arg type", i)
		}
	}
	_, err := msg.ValidateOnChainCallTriggers()
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected valid address")
}

func TestValidateUserOp_UnsupportedChain(t *testing.T) {
	msg := buildValidMsg()
	// Set unsupported chain id in on-chain action.
	msg.UserOp.ChainId = "unsupported"
	err := msg.ValidateUserOp()
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported chain id")
}

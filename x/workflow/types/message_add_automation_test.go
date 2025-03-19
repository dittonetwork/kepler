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

// buildValidOnChainAction returns a valid on-chain action.
func buildValidOnChainAction(chainID string) *types.Action {
	return &types.Action{
		Action: &types.Action_OnChain{
			OnChain: &types.OnChainAction{
				ContractAddress: validEthAddress,
				ChainId:         chainID,
			},
		},
	}
}

// buildValidMsg creates a valid MsgAddAutomation.
func buildValidMsg() *types.MsgAddAutomation {
	chainID := "1"
	return &types.MsgAddAutomation{
		Creator:  validCosmosAddress,
		ExpireAt: time.Now().Add(time.Hour).Unix(),
		Triggers: []*types.Trigger{
			buildValidOnChainCallTrigger(chainID),
			buildValidScheduleTrigger(),
			buildValidCountTrigger(),
		},
		Actions: []*types.Action{
			buildValidOnChainAction(chainID),
		},
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

func TestMsgAddAutomation_ValidateBasic_PastExpireAt(t *testing.T) {
	msg := buildValidMsg()
	msg.ExpireAt = time.Now().Add(-time.Minute).Unix()
	err := msg.ValidateBasic()
	require.Error(t, err)
	require.Contains(t, err.Error(), "expire at time must be in the future")
}

func TestMsgAddAutomation_ValidateBasic_MismatchedChainID(t *testing.T) {
	msg := buildValidMsg()
	// Change the chain id in the on-chain action to mismatch.
	msg.Actions[0].GetOnChain().ChainId = "137"
	err := msg.ValidateBasic()
	require.Error(t, err)
	require.Contains(t, err.Error(), "triggers and actions must be on the same chain")
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

func TestValidateOnChainActions_InvalidContract(t *testing.T) {
	msg := buildValidMsg()
	// Set invalid contract address in on-chain action.
	for i, act := range msg.Actions {
		if oca := act.GetOnChain(); oca != nil {
			oca.ContractAddress = "invalid_contract"
			t.Logf("Action %d has invalid contract address", i)
		}
	}
	_, err := msg.ValidateOnChainActions()
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid contract address")
}

func TestValidateOnChainActions_UnsupportedChain(t *testing.T) {
	msg := buildValidMsg()
	// Set unsupported chain id in on-chain action.
	for i, act := range msg.Actions {
		if oca := act.GetOnChain(); oca != nil {
			oca.ChainId = "unsupported"
			t.Logf("Action %d has unsupported chain id", i)
		}
	}
	_, err := msg.ValidateOnChainActions()
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported chain id")
}

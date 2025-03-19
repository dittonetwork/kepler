package types

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robfig/cron/v3"
)

const defaultIntegerBase = 10

var _ sdk.Msg = &MsgAddAutomation{}

func (msg *MsgAddAutomation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	triggersChainID, err := msg.ValidateOnChainCallTriggers()
	if err != nil {
		return err
	}

	if err = msg.ValidateScheduleTriggers(); err != nil {
		return err
	}

	if err = msg.ValidateCountTriggers(); err != nil {
		return err
	}

	actionsChainID, err := msg.ValidateOnChainActions()
	if err != nil {
		return err
	}

	if triggersChainID != actionsChainID {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"triggers and actions must be on the same chain",
		)
	}

	if time.Unix(msg.ExpireAt, 0).Before(time.Now()) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "expire at time must be in the future")
	}

	return nil
}

// ValidateCountTriggers validates all count triggers of an automation.
func (msg *MsgAddAutomation) ValidateCountTriggers() error {
	for i, t := range msg.Triggers {
		cnt := t.GetCount()
		if cnt != nil {
			if t.GetCount().RepeatCount < 1 {
				return errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: repeat count must be greater than 0", i),
				)
			}
		}
	}

	return nil
}

// ValidateScheduleTriggers validates all schedule triggers of an automation.
func (msg *MsgAddAutomation) ValidateScheduleTriggers() error {
	for i, t := range msg.Triggers {
		schd := t.GetSchedule()
		if schd != nil {
			if _, parseErr := cron.ParseStandard(t.GetSchedule().Cron); parseErr != nil {
				return errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: invalid cron expression", i),
				)
			}
		}
	}

	return nil
}

// ValidateOnChainCallTriggers validates all on-chain-call triggers of an automation.
//
//nolint:gocognit // this is a validation function of a big msg, it must be this complex
func (msg *MsgAddAutomation) ValidateOnChainCallTriggers() (string, error) {
	chainID := ""
	for i, t := range msg.Triggers {
		occ := t.GetOnChainCall()
		//nolint:nestif // trigger is big
		if occ != nil {
			if !common.IsHexAddress(occ.Contract) {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: invalid contract address", i),
				)
			}

			if !SupportedChainIDs.IsSupported(occ.ChainId) {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: unsupported chain id %s", i, occ.ChainId),
				)
			}

			if chainID == "" {
				chainID = occ.ChainId
			}

			if chainID != occ.ChainId {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					"only one chain id is supported per automation",
				)
			}

			if occ.MethodAbi == nil {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: method_abi must be provided", i),
				)
			}
			if occ.MethodAbi.Name == "" {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: method name cannot be empty", i),
				)
			}
			if len(occ.Args) != len(occ.MethodAbi.Inputs) {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf(
						"trigger %d: number of args (%d) does not match number of inputs in method ABI (%d)",
						i,
						len(occ.Args),
						len(occ.MethodAbi.Inputs),
					),
				)
			}
			for j, input := range occ.MethodAbi.Inputs {
				arg := occ.Args[j]
				if err := validateArgAgainstInputType(arg, input.Type); err != nil {
					return "", errorsmod.Wrap(
						sdkerrors.ErrInvalidRequest,
						fmt.Sprintf("trigger %d, argument %d: %v", i, j, err),
					)
				}
			}
		}
	}

	return chainID, nil
}

func (msg *MsgAddAutomation) ValidateOnChainActions() (string, error) {
	chainID := ""
	for i, a := range msg.Actions {
		oca := a.GetOnChain()
		if oca != nil {
			if !common.IsHexAddress(oca.ContractAddress) {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("action %d: invalid contract address", i),
				)
			}

			if !SupportedChainIDs.IsSupported(oca.ChainId) {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("action %d: unsupported chain id %s", i, oca.ChainId),
				)
			}

			if chainID == "" {
				chainID = oca.ChainId
			}

			if chainID != oca.ChainId {
				return "", errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					"only one chain id is supported per automation",
				)
			}
		}
	}

	return chainID, nil
}

// validateArgAgainstInputType validates that an argument (as a string)
// is appropriate for the expected input type. Supported types are "address", "uint256", and "string".
func validateArgAgainstInputType(arg string, expectedType string) error {
	switch strings.ToLower(expectedType) {
	case "address":
		if !common.IsHexAddress(arg) {
			return fmt.Errorf("expected address, got %q", arg)
		}
	case "uint256":
		// Try to parse the argument as a base-10 big integer.
		if _, ok := new(big.Int).SetString(arg, defaultIntegerBase); !ok {
			return fmt.Errorf("expected uint256, got %q", arg)
		}
	case "string":
		// No validation needed for generic strings.
	default:
		// For unsupported types, return an error.
		return fmt.Errorf("unsupported input type %q", expectedType)
	}
	return nil
}

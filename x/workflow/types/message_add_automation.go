package types

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robfig/cron/v3"
)

const (
	abiFunctionTypeFunction = "function"
)

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
func (msg *MsgAddAutomation) ValidateOnChainCallTriggers() (string, error) {
	var chainID string
	for i, t := range msg.Triggers {
		occ := t.GetOnChainCall()
		if occ == nil {
			continue
		}

		// Validate the contract address.
		if err := validateContractAddress(occ, i); err != nil {
			return "", err
		}

		// Validate that the chain id is supported.
		if err := validateSupportedChainID(occ, i); err != nil {
			return "", err
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

		// Validate the method ABI, inputs, outputs, and arguments.
		if err := validateMethodABI(occ, i); err != nil {
			return "", err
		}
	}
	return chainID, nil
}

// validateContractAddress ensures the contract address is a valid Ethereum hex address.
func validateContractAddress(occ *OnChainCallTrigger, idx int) error {
	if !common.IsHexAddress(occ.Contract) {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("trigger %d: invalid contract address", idx),
		)
	}
	return nil
}

// validateSupportedChainID ensures that the chain id is supported.
func validateSupportedChainID(occ *OnChainCallTrigger, idx int) error {
	if !SupportedChainIDs.IsSupported(occ.ChainId) {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("trigger %d: unsupported chain id %s", idx, occ.ChainId),
		)
	}
	return nil
}

// validateMethodABI validates the provided method ABI along with its inputs, outputs, and arguments.
func validateMethodABI(occ *OnChainCallTrigger, idx int) error {
	if occ.MethodAbi == nil {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("trigger %d: method_abi must be provided", idx),
		)
	}
	if occ.MethodAbi.Name == "" {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("trigger %d: method name cannot be empty", idx),
		)
	}
	if occ.MethodAbi.Type != abiFunctionTypeFunction {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("trigger %d: method type must be %q", idx, abiFunctionTypeFunction),
		)
	}
	if len(occ.Args) != len(occ.MethodAbi.Inputs) {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf(
				"trigger %d: number of args (%d) does not match number of inputs in method ABI (%d)",
				idx,
				len(occ.Args),
				len(occ.MethodAbi.Inputs),
			),
		)
	}
	if err := validateMethodABIInputs(occ, idx); err != nil {
		return err
	}
	if err := validateMethodABIOutputs(occ, idx); err != nil {
		return err
	}
	return nil
}

// validateMethodABIInputs validates each input parameter and its corresponding argument.
func validateMethodABIInputs(occ *OnChainCallTrigger, idx int) error {
	for j, input := range occ.MethodAbi.Inputs {
		if input.Name == "" {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, input %d: input name cannot be empty", idx, j),
			)
		}
		if input.Type == "" {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, input %d: input type cannot be empty", idx, j),
			)
		}
		it, err := abi.NewType(input.Type, "", nil)
		if err != nil {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, input %d: %v", idx, j, err),
			)
		}
		convertedArg, err := ConvertArgAgainstType(occ.Args[j], input.Type)
		if err != nil {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, argument %d: %v", idx, j, err),
			)
		}
		args := abi.Arguments{
			{
				Type: it,
				Name: input.Name,
			},
		}
		if _, packErr := args.Pack(convertedArg); packErr != nil {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, argument %d: %v", idx, j, err),
			)
		}
	}
	return nil
}

// validateMethodABIOutputs validates each output parameter in the method ABI.
func validateMethodABIOutputs(occ *OnChainCallTrigger, idx int) error {
	for k, output := range occ.MethodAbi.Outputs {
		if output.Name == "" {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, output %d: output name cannot be empty", idx, k),
			)
		}
		if output.Type == "" {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, output %d: output type cannot be empty", idx, k),
			)
		}
		if _, err := abi.NewType(output.Type, "", nil); err != nil {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, output %d: %v", idx, k, err),
			)
		}
	}
	return nil
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

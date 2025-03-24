package types

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dittonetwork/kepler/utils/converter"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robfig/cron/v3"
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

	if err = msg.ValidateGasLimitTriggers(); err != nil {
		return err
	}

	if err = msg.ValidateValidUntilTrigger(); err != nil {
		return err
	}

	err = msg.ValidateUserOp()
	if err != nil {
		return err
	}

	if triggersChainID != "" && triggersChainID != msg.UserOp.GetChainId() {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"triggers and user_op must be on the same chain",
		)
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
			if _, parseErr := cron.ParseStandard(schd.Cron); parseErr != nil {
				return errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: invalid cron expression", i),
				)
			}
		}
	}

	return nil
}

func (msg *MsgAddAutomation) ValidateGasLimitTriggers() error {
	for i, t := range msg.Triggers {
		gl := t.GetGasLimit()
		if gl != nil {
			maxFeePerGas, ok := new(big.Int).SetString(gl.MaxFeePerGas, converter.DefaultIntegerBase)
			if !ok {
				return errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: invalid max fee per gas", i),
				)
			}
			maxPriorityFeePerGas, ok := new(big.Int).SetString(gl.MaxPriorityFeePerGas, converter.DefaultIntegerBase)
			if !ok {
				return errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: invalid max priority fee per gas", i),
				)
			}

			one := new(big.Int).SetUint64(1)
			if maxFeePerGas.Cmp(one) < 0 || maxPriorityFeePerGas.Cmp(one) < 0 {
				return errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("trigger %d: gas limit must be greater than or equal to 1", i),
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

	wrappedABI := fmt.Sprintf("[%s]", occ.MethodAbi.Abi)
	abiJSON, err := abi.JSON(strings.NewReader(wrappedABI))
	if err != nil {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("trigger %d: %v", idx, err),
		)
	}

	if len(occ.Args) != len(abiJSON.Methods[occ.MethodAbi.Name].Inputs) {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf(
				"trigger %d: number of args (%d) does not match number of inputs in method ABI (%d)",
				idx,
				len(occ.Args),
				len(abiJSON.Methods[occ.MethodAbi.Name].Inputs),
			),
		)
	}

	if err = validateMethodABIArguments(abiJSON.Methods[occ.MethodAbi.Name].Inputs, occ.Args, idx); err != nil {
		return err
	}

	return nil
}

// validateMethodABIArguments validates each input parameter and its corresponding argument.
func validateMethodABIArguments(abiArguments abi.Arguments, arguments []string, idx int) error {
	for j, arg := range abiArguments {
		convertedArg, err := converter.StrToABICompatible(arguments[j], arg.Type.String())
		if err != nil {
			return errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("trigger %d, argument %d: %v", idx, j, err),
			)
		}
		args := abi.Arguments{
			{
				Type: arg.Type,
				Name: arg.Name,
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

func (msg *MsgAddAutomation) ValidateUserOp() error {
	if msg.UserOp == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "user operation is required")
	}

	if !SupportedChainIDs.IsSupported(msg.UserOp.GetChainId()) {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("user_op: unsupported chain id %s", msg.UserOp.GetChainId()),
		)
	}

	return nil
}

func (msg *MsgAddAutomation) ValidateValidUntilTrigger() error {
	hasValidUntilTrigger := false

	for _, t := range msg.Triggers {
		if t.GetValidUntil() != nil {
			if hasValidUntilTrigger {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "only one expire time trigger is allowed")
			}

			expireTime := time.Unix(t.GetValidUntil().Timestamp, 0)
			currentTime := time.Now()

			// Mark that we have this trigger already
			hasValidUntilTrigger = true

			if expireTime.Before(currentTime) {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "valid until time must be in the future")
			}
		}
	}

	return nil
}

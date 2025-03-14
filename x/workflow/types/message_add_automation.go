package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ChainIDEth     = "1"
	ChainIDPolygon = "137"
)

var _ sdk.Msg = &MsgAddAutomation{}

func (msg *MsgAddAutomation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	chainIDs := map[string]struct{}{}
	for _, t := range msg.Triggers {
		if t.GetOnChain() != nil {
			if t.GetOnChain().ChainId != ChainIDEth && t.GetOnChain().ChainId != ChainIDPolygon {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "unsupported chain id")
			}

			chainIDs[t.GetOnChain().ChainId] = struct{}{}
		}
	}

	if err = msg.CheckExpireTimeTrigger(); err != nil {
		return err
	}

	if err = msg.ValidateUserOp(); err != nil {
		return err
	}

	if len(chainIDs) > 1 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "we only support one chain id per automation")
	}

	return nil
}

func (msg *MsgAddAutomation) ValidateUserOp() error {
	if msg.UserOp == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "user operation is required")
	}

	if msg.UserOp.GetChainId() != ChainIDEth && msg.UserOp.GetChainId() != ChainIDPolygon {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "unsupported chain id")
	}

	return nil
}

func (msg *MsgAddAutomation) CheckExpireTimeTrigger() error {
	hasExpireTimeTrigger := false

	for _, t := range msg.Triggers {
		if t.GetExpireAt() != nil {
			if hasExpireTimeTrigger {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "only one expire time trigger is allowed")
			}

			expireTime := time.Unix(t.GetExpireAt().Timestamp, 0)
			currentTime := time.Now()

			// Mark that we have an expire time trigger
			hasExpireTimeTrigger = true

			if expireTime.Before(currentTime) {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "expire at time must be in the future")
			}
		}
	}

	return nil
}

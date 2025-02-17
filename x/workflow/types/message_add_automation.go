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

	for _, a := range msg.Actions {
		if a.GetOnChain() != nil {
			if a.GetOnChain().ChainId != ChainIDEth && a.GetOnChain().ChainId != ChainIDPolygon {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "unsupported chain id")
			}

			chainIDs[a.GetOnChain().ChainId] = struct{}{}
		}
	}

	if len(chainIDs) > 1 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "we only support one chain id per automation")
	}

	if time.Unix(msg.ExpireAt, 0).Before(time.Now()) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "expire at time must be in the future")
	}

	return nil
}

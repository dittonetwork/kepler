package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	blockHashLength = 66
)

var (
	_ sdk.Msg = &MsgUpdateValidatorsSet{}
)

func (msg MsgUpdateValidatorsSet) Validate() error {
	if len(msg.Operators) == 0 {
		return sdkerrors.Wrap(ErrUpdateValidator, "no operators provided")
	}

	if len(msg.Info.BlockHash) != blockHashLength {
		return sdkerrors.Wrap(ErrUpdateValidator, "invalid block hash")
	}

	if msg.Info.EpochNum <= 0 {
		return sdkerrors.Wrap(ErrUpdateValidator, "epoch number must be greater than 0")
	}

	for _, operator := range msg.Operators {
		if len(operator.Address) == 0 {
			return sdkerrors.Wrap(ErrUpdateValidator, "operator address cannot be empty")
		}

		if !common.IsHexAddress(operator.Address) {
			return sdkerrors.Wrap(
				ErrUpdateValidator,
				"operator address is not a valid Ethereum address",
			)
		}

		if operator.ConsensusPubkey == nil {
			return sdkerrors.Wrap(
				ErrUpdateValidator,
				"operator consensus pubkey cannot be empty",
			)
		}
	}

	return nil
}

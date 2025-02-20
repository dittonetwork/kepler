package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitJobResult{}

func NewMsgSubmitJobResult(creator string, status string, committeeId string, chainId string, automationId int32, txHash string, executorAddress string, createdAt int32, executedAt int32, signedAt int32, signs string, payload string) *MsgSubmitJobResult {
	return &MsgSubmitJobResult{
		Creator:         creator,
		Status:          status,
		CommitteeId:     committeeId,
		ChainId:         chainId,
		AutomationId:    automationId,
		TxHash:          txHash,
		ExecutorAddress: executorAddress,
		CreatedAt:       createdAt,
		ExecutedAt:      executedAt,
		SignedAt:        signedAt,
		Signs:           signs,
		Payload:         payload,
	}
}

func (msg *MsgSubmitJobResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

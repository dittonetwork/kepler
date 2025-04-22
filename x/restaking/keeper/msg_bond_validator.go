package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// BondValidator complete the bonding process for a validator recognized in Bonding status.
func (s msgServer) BondValidator(
	ctx context.Context,
	msg *types.MsgBondValidator,
) (*types.MsgBondValidatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	valAddr, err := s.validatorAddressCodec.StringToBytes(msg.Owner)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid validator address: %s", err)
	}

	s.logger.With("acc_address", sdk.AccAddress(valAddr), "owner", msg.Owner, "val_addr", valAddr).
		Debug("BondValidator", "msg", msg)

	acc := s.accounts.GetAccount(ctx, valAddr)

	if acc.GetPubKey() == nil {
		s.logger.Error("failed to get pubkey for account", "address", msg.Owner)
		return nil, types.ErrGenesisInit
	}

	evmaddr, err := types.ToKeccakLast20(acc.GetPubKey())
	if err != nil {
		return nil, err
	}

	operator, err := s.repository.GetPendingOperator(sdkCtx, evmaddr.String())
	if err != nil {
		return nil, err
	}

	validator := &types.Validator{
		OperatorAddress:    msg.Owner,
		ConsensusPubkey:    operator.ConsensusPubkey,
		EvmOperatorAddress: evmaddr.String(),
		Description:        msg.Description,
		VotingPower:        operator.VotingPower,
		IsEmergency:        operator.IsEmergency,
		Protocol:           operator.Protocol,
		Status:             operator.Status,
	}

	err = s.repository.AddValidatorsChange(sdkCtx, *validator, types.ValidatorChangeTypeCreate)
	if err != nil {
		return nil, err
	}

	if err = s.repository.RemovePendingOperator(sdkCtx, evmaddr.String()); err != nil {
		return nil, err
	}

	return &types.MsgBondValidatorResponse{}, nil
}

package keeper

import (
	"context"

	"kepler/x/committees/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) UpdateValidatorsStakes(ctx context.Context, msg *types.MsgUpdateValidatorsStakes) (*types.MsgUpdateValidatorsStakesResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgUpdateValidatorsStakesResponse{}, nil
}

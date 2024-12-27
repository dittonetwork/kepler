package keeper

import (
	"context"
	"fmt"

	"kepler/x/committees/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) UpdateValidatorsStakes(ctx context.Context, msg *types.MsgUpdateValidatorsStakes) (*types.MsgUpdateValidatorsStakesResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	if params.TrustedAddress != msg.GetCreator() {
		return nil, fmt.Errorf("message creator (%s) does not match with trusted address (%s)", msg.GetCreator(), params.TrustedAddress)
	}

	stakes := make(map[string]uint64, len(msg.GetAaddresses()))
	for i, address := range msg.GetAaddresses() {
		stakes[address] = msg.GetStakes()[i]
	}

	k.stakingKeeper.UpdateValidatorsStakes(ctx, stakes)

	return &types.MsgUpdateValidatorsStakesResponse{}, nil
}

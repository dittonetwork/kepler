package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

func (s msgServer) UpdateValidatorsSet(
	ctx context.Context,
	msg *types.MsgUpdateValidatorsSet,
) (*types.MsgUpdateValidatorsSetResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if s.authority != msg.Authority {
		return nil, sdkerrors.Wrapf(
			govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", s.authority, msg.Authority,
		)
	}

	// Check if the block height is higher than the last update
	lastUpdate, err := s.repository.GetLastUpdate(sdkCtx)
	if err != nil {
		return nil, err
	}

	if lastUpdate.BlockHeight >= msg.Info.BlockHeight {
		return nil, sdkerrors.Wrap(types.ErrUpdateValidator, "block height is lower than last update")
	}

	if lastUpdate.EpochNum >= msg.Info.EpochNum {
		return nil, sdkerrors.Wrap(types.ErrUpdateValidator, "epoch number is lower than last update")
	}

	delta, err := s.makeDeltaUpdates(sdkCtx, msg.Operators)
	if err != nil {
		return nil, err
	}

	if err = s.processCreatedOperators(sdkCtx, delta.Created); err != nil {
		return nil, err
	}

	if err = s.processDeletedOperators(sdkCtx, delta.Deleted); err != nil {
		return nil, err
	}

	if err = s.processUpdatedValidators(sdkCtx, delta.Updated); err != nil {
		return nil, err
	}

	return &types.MsgUpdateValidatorsSetResponse{}, nil
}

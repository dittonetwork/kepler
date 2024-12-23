package types

import (
	"testing"

	"kepler/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateAlliancesTimeline_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAlliancesTimeline
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAlliancesTimeline{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateAlliancesTimeline{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateAlliancesTimeline_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAlliancesTimeline
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAlliancesTimeline{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAlliancesTimeline{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteAlliancesTimeline_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAlliancesTimeline
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteAlliancesTimeline{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteAlliancesTimeline{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

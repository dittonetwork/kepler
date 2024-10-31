package types

import (
	"testing"

	"kepler/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateFinalizedBlockInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateFinalizedBlockInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateFinalizedBlockInfo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateFinalizedBlockInfo{
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

func TestMsgUpdateFinalizedBlockInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateFinalizedBlockInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateFinalizedBlockInfo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateFinalizedBlockInfo{
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

func TestMsgDeleteFinalizedBlockInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteFinalizedBlockInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteFinalizedBlockInfo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteFinalizedBlockInfo{
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

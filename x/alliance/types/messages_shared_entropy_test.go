package types

import (
	"testing"

	"kepler/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateSharedEntropy_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSharedEntropy
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSharedEntropy{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSharedEntropy{
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

func TestMsgUpdateSharedEntropy_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSharedEntropy
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSharedEntropy{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSharedEntropy{
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

func TestMsgDeleteSharedEntropy_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSharedEntropy
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSharedEntropy{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSharedEntropy{
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

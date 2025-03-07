package types

import (
	"testing"
	"time"

	"github.com/dittonetwork/kepler/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAddAutomation_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddAutomation
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddAutomation{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "unsupported chain id in trigger",
			msg: MsgAddAutomation{
				Creator:  sample.AccAddress(),
				ExpireAt: time.Now().Add(time.Hour).Unix(),
				Triggers: []*Trigger{
					{
						Trigger: &Trigger_OnChain{OnChain: &OnChainTrigger{
							ChainId: "unsupported_chain_id",
						},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "unsupported chain id in action",
			msg: MsgAddAutomation{
				Creator:  sample.AccAddress(),
				ExpireAt: time.Now().Add(time.Hour).Unix(),
				Triggers: []*Trigger{
					{
						Trigger: &Trigger_OnChain{OnChain: &OnChainTrigger{
							ChainId: "1",
						},
						},
					},
				},
				Actions: []*Action{
					{
						Action: &Action_OnChain{OnChain: &OnChainAction{
							ChainId: "unsupported_chain_id",
						},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "both chain ids are supported, but more than one chain id",
			msg: MsgAddAutomation{
				Creator:  sample.AccAddress(),
				ExpireAt: time.Now().Add(time.Hour).Unix(),
				Triggers: []*Trigger{
					{
						Trigger: &Trigger_OnChain{OnChain: &OnChainTrigger{
							ChainId: "1",
						},
						},
					},
				},
				Actions: []*Action{
					{
						Action: &Action_OnChain{OnChain: &OnChainAction{
							ChainId: "137",
						},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "expire at in the past",
			msg: MsgAddAutomation{
				Creator:  sample.AccAddress(),
				ExpireAt: time.Now().Add(-1 * time.Hour).Unix(),
				Triggers: []*Trigger{
					{
						Trigger: &Trigger_OnChain{OnChain: &OnChainTrigger{
							ChainId: "1",
						},
						},
					},
				},
				Actions: []*Action{
					{
						Action: &Action_OnChain{OnChain: &OnChainAction{
							ChainId: "1",
						},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
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

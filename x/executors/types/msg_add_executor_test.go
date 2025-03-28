package types_test

import (
	"testing"

	"github.com/dittonetwork/kepler/x/executors/types"
	"github.com/stretchr/testify/require"
)

// Use known valid bech32 addresses for testing. These should pass sdk.AccAddressFromBech32.
const validCreator = "cosmos1ghekyjucln7y67ntx7cf27m9dpuxxemn4c8g4r"
const validOwner = "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv"

func TestMsgAddExecutor_ValidateBasic_Valid(t *testing.T) {
	msg := &types.MsgAddExecutor{
		Creator:      validCreator,
		OwnerAddress: validOwner,
		PublicKey:    "cosmospub1addwnpep9ym0rv96yl0xxyzz", // dummy pubkey
	}

	err := msg.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgAddExecutor_ValidateBasic_InvalidCreator(t *testing.T) {
	msg := &types.MsgAddExecutor{
		Creator:      "invalidaddress", // not a valid bech32 address
		OwnerAddress: validOwner,
		PublicKey:    "cosmospub1addwnpep9ym0rv96yl0xxyzz",
	}

	err := msg.ValidateBasic()
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid creator address")
}

func TestMsgAddExecutor_ValidateBasic_InvalidOwner(t *testing.T) {
	msg := &types.MsgAddExecutor{
		Creator:      validCreator,
		OwnerAddress: "invalidaddress", // not a valid bech32 address
		PublicKey:    "cosmospub1addwnpep9ym0rv96yl0xxyzz",
	}

	err := msg.ValidateBasic()
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid owner address")
}

func TestMsgAddExecutor_ValidateBasic_CreatorOwnerSame(t *testing.T) {
	msg := &types.MsgAddExecutor{
		Creator:      validCreator,
		OwnerAddress: validCreator, // same as creator
		PublicKey:    "cosmospub1addwnpep9ym0rv96yl0xxyzz",
	}

	err := msg.ValidateBasic()
	require.Error(t, err)
	require.Contains(t, err.Error(), "creator and owner cannot be the same")
}

package types_test

import (
	codectestutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"kepler/x/staking/types"
	"testing"
)

func TestValidatorTestEquivalent(t *testing.T) {
	val1 := newValidator(t, valAddr1, pk1)
	val2 := newValidator(t, valAddr1, pk1)
	require.Equal(t, val1.String(), val2.String())

	val2 = newValidator(t, valAddr2, pk1)
	require.NotEqual(t, val1.String(), val2.String())
}

func TestUpdateDescription(t *testing.T) {
	d1 := types.Description{
		Website: "https://validator.cosmos",
		Details: "Test validator",
	}

	d2 := types.Description{
		Moniker:  types.DoNotModifyDesc,
		Identity: types.DoNotModifyDesc,
		Website:  types.DoNotModifyDesc,
		Details:  types.DoNotModifyDesc,
	}

	d3 := types.Description{
		Moniker:  "",
		Identity: "",
		Website:  "",
		Details:  "",
	}

	d, err := d1.UpdateDescription(d2)
	require.Nil(t, err)
	require.Equal(t, d, d1)

	d, err = d1.UpdateDescription(d3)
	require.Nil(t, err)
	require.Equal(t, d, d3)
}

func TestUpdateStatus(t *testing.T) {
	validator := newValidator(t, valAddr1, pk1)
	require.Equal(t, types.Unbonded, validator.Status)
	require.Equal(t, int64(0), validator.Tokens.Int64())

	// Unbonded to Bonded
	validator.UpdateStatus(types.Bonded)
	require.Equal(t, types.Bonded, validator.Status)

	// Bonded to Unbonding
	validator.UpdateStatus(types.Unbonding)
	require.Equal(t, types.Unbonding, validator.Status)

	// Unbonding to Bonded
	validator.UpdateStatus(types.Bonded)
	require.Equal(t, types.Bonded, validator.Status)
}

// Creates a new validators and asserts the error check.
func newValidator(t *testing.T, operator sdk.ValAddress, pubKey cryptotypes.PubKey) types.Validator {
	t.Helper()
	addr, err := codectestutil.CodecOptions{}.GetValidatorCodec().BytesToString(operator)
	require.NoError(t, err)
	v, err := types.NewValidator(addr, pubKey, types.Description{})
	require.NoError(t, err)
	return v
}

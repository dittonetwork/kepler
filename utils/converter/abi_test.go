package converter_test

import (
	"math/big"
	"testing"

	"github.com/dittonetwork/kepler/utils/converter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestConvertArg_Address_Valid(t *testing.T) {
	arg := "0x1234567890abcdef1234567890abcdef12345678"
	result, err := converter.StrToABICompatible(arg, "address")
	require.NoError(t, err)
	addr, ok := result.(common.Address)
	require.True(t, ok)
	require.Equal(t, common.HexToAddress(arg), addr)
}

func TestConvertArg_Address_Invalid(t *testing.T) {
	arg := "123456"
	_, err := converter.StrToABICompatible(arg, "address")
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected valid address")
}

func TestConvertArg_Uint256_Valid(t *testing.T) {
	arg := "123456789"
	result, err := converter.StrToABICompatible(arg, "uint256")
	require.NoError(t, err)
	bi, ok := result.(*big.Int)
	require.True(t, ok)
	expected := new(big.Int)
	expected.SetString(arg, 10)
	require.Equal(t, 0, bi.Cmp(expected))
}

func TestConvertArg_Uint256_Invalid(t *testing.T) {
	arg := "abc"
	_, err := converter.StrToABICompatible(arg, "uint256")
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected uint256")
}

func TestConvertArg_String(t *testing.T) {
	arg := "hello world"
	result, err := converter.StrToABICompatible(arg, "string")
	require.NoError(t, err)
	s, ok := result.(string)
	require.True(t, ok)
	require.Equal(t, arg, s)
}

func TestConvertArg_Bool_Valid(t *testing.T) {
	result, err := converter.StrToABICompatible("true", "bool")
	require.NoError(t, err)
	b, ok := result.(bool)
	require.True(t, ok)
	require.True(t, b)

	result, err = converter.StrToABICompatible("false", "bool")
	require.NoError(t, err)
	b, ok = result.(bool)
	require.True(t, ok)
	require.False(t, b)
}

func TestConvertArg_Bool_Invalid(t *testing.T) {
	_, err := converter.StrToABICompatible("notabool", "bool")
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected bool")
}

func TestConvertArg_Bytes_Hex(t *testing.T) {
	arg := "0xabcdef"
	result, err := converter.StrToABICompatible(arg, "bytes")
	require.NoError(t, err)
	b, ok := result.([]byte)
	require.True(t, ok)
	require.Equal(t, common.FromHex(arg), b)
}

func TestConvertArg_Bytes_String(t *testing.T) {
	arg := "abcdef"
	result, err := converter.StrToABICompatible(arg, "bytes")
	require.NoError(t, err)
	b, ok := result.([]byte)
	require.True(t, ok)
	require.Equal(t, []byte(arg), b)
}

func TestConvertArg_Uint8(t *testing.T) {
	arg := "255"
	result, err := converter.StrToABICompatible(arg, "uint8")
	require.NoError(t, err)
	u, ok := result.(uint8)
	require.True(t, ok)
	require.Equal(t, uint8(255), u)
}

func TestConvertArg_Uint16(t *testing.T) {
	arg := "65535"
	result, err := converter.StrToABICompatible(arg, "uint16")
	require.NoError(t, err)
	u, ok := result.(uint16)
	require.True(t, ok)
	require.Equal(t, uint16(65535), u)
}

func TestConvertArg_Uint32(t *testing.T) {
	arg := "4294967295"
	result, err := converter.StrToABICompatible(arg, "uint32")
	require.NoError(t, err)
	u, ok := result.(uint32)
	require.True(t, ok)
	require.Equal(t, uint32(4294967295), u)
}

func TestConvertArg_Uint64(t *testing.T) {
	arg := "18446744073709551615" // max uint64
	result, err := converter.StrToABICompatible(arg, "uint64")
	require.NoError(t, err)
	u, ok := result.(uint64)
	require.True(t, ok)
	require.Equal(t, uint64(18446744073709551615), u)
}

func TestConvertArg_Uint(t *testing.T) {
	arg := "1000"
	result, err := converter.StrToABICompatible(arg, "uint")
	require.NoError(t, err)
	u, ok := result.(uint)
	require.True(t, ok)
	require.Equal(t, uint(1000), u)
}

func TestConvertArg_Int8(t *testing.T) {
	arg := "127"
	result, err := converter.StrToABICompatible(arg, "int8")
	require.NoError(t, err)
	i, ok := result.(int8)
	require.True(t, ok)
	require.Equal(t, int8(127), i)
}

func TestConvertArg_Int16(t *testing.T) {
	arg := "32767"
	result, err := converter.StrToABICompatible(arg, "int16")
	require.NoError(t, err)
	i, ok := result.(int16)
	require.True(t, ok)
	require.Equal(t, int16(32767), i)
}

func TestConvertArg_Int32(t *testing.T) {
	arg := "2147483647"
	result, err := converter.StrToABICompatible(arg, "int32")
	require.NoError(t, err)
	i, ok := result.(int32)
	require.True(t, ok)
	require.Equal(t, int32(2147483647), i)
}

func TestConvertArg_Int64(t *testing.T) {
	arg := "9223372036854775807"
	result, err := converter.StrToABICompatible(arg, "int64")
	require.NoError(t, err)
	i, ok := result.(int64)
	require.True(t, ok)
	require.Equal(t, int64(9223372036854775807), i)
}

func TestConvertArg_Int(t *testing.T) {
	arg := "100"
	result, err := converter.StrToABICompatible(arg, "int")
	require.NoError(t, err)
	i, ok := result.(int)
	require.True(t, ok)
	require.Equal(t, int(100), i)
}

func TestConvertArg_UnsupportedType(t *testing.T) {
	_, err := converter.StrToABICompatible("123.45", "float")
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported input type")
}

func TestConvertArg_SliceUint256(t *testing.T) {
	// JSON array string of uint256 values.
	arg := `["123", "456", "789"]`
	result, err := converter.StrToABICompatible(arg, "[]uint256")
	require.NoError(t, err)
	slice, ok := result.([]interface{})
	require.True(t, ok)
	require.Len(t, slice, 3)
	expectedValues := []string{"123", "456", "789"}
	for i, exp := range expectedValues {
		bi, ok := slice[i].(*big.Int)
		require.True(t, ok)
		expected := new(big.Int)
		expected.SetString(exp, 10)
		require.Equal(t, 0, bi.Cmp(expected))
	}
}

func TestConvertArg_FixedArrayAddress(t *testing.T) {
	// JSON array string for a fixed-length array of 2 addresses.
	arg := `["0x1234567890abcdef1234567890abcdef12345678", "0xabcdef1234567890abcdef1234567890abcdef12"]`
	result, err := converter.StrToABICompatible(arg, "[2]address")
	require.NoError(t, err)
	array, ok := result.([]interface{})
	require.True(t, ok)
	require.Len(t, array, 2)
	expected := []string{
		"0x1234567890abcdef1234567890abcdef12345678",
		"0xabcdef1234567890abcdef1234567890abcdef12",
	}
	for i, exp := range expected {
		addr, ok := array[i].(common.Address)
		require.True(t, ok)
		require.Equal(t, common.HexToAddress(exp), addr)
	}
}

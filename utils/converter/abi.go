package converter

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// DefaultIntegerBase defines the number base for integer conversions.
const (
	DefaultIntegerBase = 10
	bitSize64          = 64
	bitSize32          = 32
	bitSize16          = 16
	bitSize8           = 8
	bitSizeDefault     = 0
)

// StrToABICompatible validates and converts an argument (as a string)
// to the proper Go type expected by the ABI for the given expectedType.
// It supports slices, fixed-length arrays, and primitive types.
func StrToABICompatible(arg string, expectedType string) (interface{}, error) {
	if strings.HasPrefix(expectedType, "[]") {
		elemType := expectedType[2:]
		return convertSliceArg(arg, elemType)
	}

	if strings.HasPrefix(expectedType, "[") {
		return convertArrayArg(arg, expectedType)
	}

	return convertPrimitiveArg(arg, expectedType)
}

// convertSliceArg handles dynamic slices. It expects the argument to be a JSON array.
func convertSliceArg(arg, elemType string) (interface{}, error) {
	var raw []interface{}
	if err := json.Unmarshal([]byte(arg), &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal slice: %w", err)
	}

	resultSlice := make([]interface{}, len(raw))
	for i, elem := range raw {
		// Convert the element to string and recursively convert it.
		elemStr := fmt.Sprintf("%v", elem)
		conv, err := StrToABICompatible(elemStr, elemType)
		if err != nil {
			return nil, fmt.Errorf("failed to convert slice element %d: %w", i, err)
		}
		resultSlice[i] = conv
	}
	return resultSlice, nil
}

// convertArrayArg handles fixed-length arrays (syntax: "[N]<elemType>").
// It expects the argument to be a JSON array with exactly N elements.
func convertArrayArg(arg, expectedType string) (interface{}, error) {
	closingIndex := strings.Index(expectedType, "]")
	if closingIndex == -1 {
		return nil, fmt.Errorf("invalid array type syntax: %s", expectedType)
	}
	lengthStr := expectedType[1:closingIndex]
	arrayLength, err := strconv.Atoi(lengthStr)
	if err != nil {
		return nil, fmt.Errorf("invalid array length in type %s: %w", expectedType, err)
	}
	elemType := expectedType[closingIndex+1:]

	var raw []interface{}
	if err = json.Unmarshal([]byte(arg), &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal array: %w", err)
	}
	if len(raw) != arrayLength {
		return nil, fmt.Errorf("expected array of length %d, got %d", arrayLength, len(raw))
	}

	resultArray := make([]interface{}, len(raw))
	for i, elem := range raw {
		elemStr := fmt.Sprintf("%v", elem)
		conv, convErr := StrToABICompatible(elemStr, elemType)
		if convErr != nil {
			return nil, fmt.Errorf("failed to convert array element %d: %w", i, convErr)
		}
		resultArray[i] = conv
	}
	return resultArray, nil
}

// convertPrimitiveArg dispatches conversion based on the expected type.
func convertPrimitiveArg(arg, expectedType string) (interface{}, error) {
	switch strings.ToLower(expectedType) {
	case "address":
		return convertAddress(arg)
	case "uint256":
		return convertUint256(arg)
	case "string":
		return convertString(arg)
	case "bool":
		return convertBool(arg)
	case "bytes":
		return convertBytes(arg)
	case "uint8":
		return convertUintGeneric(arg, bitSize8)
	case "uint16":
		return convertUintGeneric(arg, bitSize16)
	case "uint32":
		return convertUintGeneric(arg, bitSize32)
	case "uint64":
		return convertUintGeneric(arg, bitSize64)
	case "uint":
		return convertUintGeneric(arg, bitSizeDefault)
	case "int8":
		return convertIntGeneric(arg, bitSize8)
	case "int16":
		return convertIntGeneric(arg, bitSize16)
	case "int32":
		return convertIntGeneric(arg, bitSize32)
	case "int64":
		return convertIntGeneric(arg, bitSize64)
	case "int":
		return convertIntGeneric(arg, bitSizeDefault)
	default:
		return nil, fmt.Errorf("unsupported input type %q", expectedType)
	}
}

// convertAddress validates and converts an argument to an Ethereum address.
func convertAddress(arg string) (interface{}, error) {
	if !common.IsHexAddress(arg) {
		return nil, fmt.Errorf("expected valid address, got %q", arg)
	}
	return common.HexToAddress(arg), nil
}

// convertUint256 converts a string to a *big.Int representing a uint256.
func convertUint256(arg string) (interface{}, error) {
	bi, ok := new(big.Int).SetString(arg, DefaultIntegerBase)
	if !ok {
		return nil, fmt.Errorf("expected uint256, got %q", arg)
	}
	return bi, nil
}

// convertString returns the string unchanged.
func convertString(arg string) (interface{}, error) {
	return arg, nil
}

// convertBool parses a string as a boolean.
func convertBool(arg string) (interface{}, error) {
	b, err := strconv.ParseBool(arg)
	if err != nil {
		return nil, fmt.Errorf("expected bool, got %q: %w", arg, err)
	}
	return b, nil
}

// convertBytes converts the string to a byte slice.
// If the string starts with "0x", it decodes it as hex.
func convertBytes(arg string) (interface{}, error) {
	if strings.HasPrefix(arg, "0x") {
		return common.FromHex(arg), nil
	}
	return []byte(arg), nil
}

// convertUintGeneric converts a string to an unsigned integer of the specified bitSize.
//
//nolint:gosec // This function is used to convert unsigned integers.
func convertUintGeneric(arg string, bitSize int) (interface{}, error) {
	u, err := strconv.ParseUint(arg, 10, bitSize)
	if err != nil {
		return nil, fmt.Errorf("expected uint%d, got %q: %w", bitSize, arg, err)
	}
	switch bitSize {
	case bitSize8:
		return uint8(u), nil
	case bitSize16:
		return uint16(u), nil
	case bitSize32:
		return uint32(u), nil
	case bitSize64:
		return u, nil
	default:
		return uint(u), nil
	}
}

// convertIntGeneric converts a string to a signed integer of the specified bitSize.
//
//nolint:gosec // This function is used to convert integers.
func convertIntGeneric(arg string, bitSize int) (interface{}, error) {
	i, err := strconv.ParseInt(arg, 10, bitSize)
	if err != nil {
		return nil, fmt.Errorf("expected int%d, got %q: %w", bitSize, arg, err)
	}
	switch bitSize {
	case bitSize8:
		return int8(i), nil
	case bitSize16:
		return int16(i), nil
	case bitSize32:
		return int32(i), nil
	case bitSize64:
		return i, nil
	default:
		return int(i), nil
	}
}

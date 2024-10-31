package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StakedAmountInfoKeyPrefix is the prefix to retrieve all StakedAmountInfo
	StakedAmountInfoKeyPrefix = "StakedAmountInfo/value/"
)

// StakedAmountInfoKey returns the store key to retrieve a StakedAmountInfo from the index fields
func StakedAmountInfoKey(
	ethereumAddress string,
) []byte {
	var key []byte

	ethereumAddressBytes := []byte(ethereumAddress)
	key = append(key, ethereumAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}

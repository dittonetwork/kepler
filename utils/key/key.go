package key

import "encoding/binary"

const byteSize = 8

func GetBytesKeyFromUint64(id uint64) []byte {
	bz := make([]byte, byteSize)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

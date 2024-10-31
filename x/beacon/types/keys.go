package types

const (
	// ModuleName defines the module name
	ModuleName = "beacon"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_beacon"
)

var (
	ParamsKey = []byte("p_beacon")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	FinalizedBlockInfoKey = "FinalizedBlockInfo/value/"
)

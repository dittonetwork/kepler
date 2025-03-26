package types

const (
	// ModuleName defines the module name.
	ModuleName = "executors"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_executors"
)

var (
	ParamsKey = []byte("p_executors")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

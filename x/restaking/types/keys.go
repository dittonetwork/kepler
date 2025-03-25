package types

const (
	// ModuleName defines the module name.
	ModuleName = "restaking"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_restaking"
)

var (
	ParamsKey = []byte("p_restaking")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

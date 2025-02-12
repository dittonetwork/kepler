package types

const (
	// ModuleName defines the module name
	ModuleName = "instant"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_instant"
)

var (
	ParamsKey = []byte("p_instant")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

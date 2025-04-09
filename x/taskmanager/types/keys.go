package types

const (
	// ModuleName defines the module name.
	ModuleName = "taskmanager"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_taskmanager"
)

var (
	ParamsKey = []byte("p_taskmanager")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "job"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_job"
)

var (
	ParamsKey = []byte("p_job")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

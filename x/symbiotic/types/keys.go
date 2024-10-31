package types

const (
	// ModuleName defines the module name
	ModuleName = "symbiotic"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_symbiotic"
)

var (
	ParamsKey = []byte("p_symbiotic")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ContractAddressKey = "ContractAddress/value/"
)

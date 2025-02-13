package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name.
	ModuleName = "job"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_job"
)

var (
	ParamsKey = []byte("p_job")

	JobsPrefix = collections.NewPrefix(0)
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

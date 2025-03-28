package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name.
	ModuleName = "executors"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// CollectionNameExecutors is the name of executors collection.
	CollectionNameExecutors                = "executors"
	CollectionIndexExecutorsByOwnerAddress = "executors_by_owner_address"

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_executors"
)

var (
	ParamsKey = []byte("p_executors")

	ExecutorsPrefix = collections.NewPrefix(0)
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

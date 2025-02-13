package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name.
	ModuleName = "committee"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_committee"
)

var (
	RandaoCommitmentsKey = collections.NewPrefix(0)
	RandaoRevealsKey     = collections.NewPrefix(1)

	ParamsKey = []byte("p_committee")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

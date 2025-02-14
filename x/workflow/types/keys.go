package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name.
	ModuleName = "workflow"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_workflow"
)

var (
	ParamsKey                                     = []byte("p_workflow")
	KeyPrefixAutomation        collections.Prefix = collections.NewPrefix(CollectionNameAutomations)
	KeyPrefixActiveAutomations                    = collections.NewPrefix(CollectionNameActiveAutomations)
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

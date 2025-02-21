package types

import (
	"cosmossdk.io/collections"
)

const (
	// ModuleName defines the module name.
	ModuleName = "committee"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_committee"
)

var (
	// CommitteeStoreKeyPrefix is the prefix for the committee store key.
	CommitteeStoreKeyPrefix = collections.NewPrefix(1)

	// ChainIDStoreKeyPrefix is the prefix for the chain id store key.
	ChainIDStoreKeyPrefix = collections.NewPrefix("chain_id")

	// ActiveCommitteeStoreKeyPrefix is the prefix for the active committee store key.
	ActiveCommitteeStoreKeyPrefix = collections.NewPrefix("active_committee")
)

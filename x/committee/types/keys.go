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

	latestEpochStorePrefixIndex = 2
)

var (
	// CommitteesStoreKeyPrefix is the prefix for the committee store key.
	CommitteesStoreKeyPrefix = collections.NewPrefix(0)

	// CommitteesEmergencyIdxPrefix is the prefix for the emergency committees emergency index key.
	CommitteesEmergencyIdxPrefix = collections.NewPrefix(1)

	// LatestEpochStorePrefix is the prefix for the latest epoch store key.
	LatestEpochStorePrefix = collections.NewPrefix(latestEpochStorePrefixIndex)
)

const (
	EventKeyReportGot = "report_got"
)

package types

import (
	"cosmossdk.io/collections"
)

const (
	// ModuleName defines the module name.
	ModuleName = "epochs"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName
)

var (
	// KeyPrefixEpoch defines the prefix for the epoch info store.
	KeyPrefixEpoch = collections.NewPrefix(1)
)

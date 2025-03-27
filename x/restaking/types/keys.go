package types

import (
	"cosmossdk.io/collections"
)

const (
	// ModuleName defines the module name.
	ModuleName = "restaking"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_restaking"

	// CollectionNameValidators is the name of validators collection.
	CollectionNameValidators = "validators"
	// CollectionNameLastUpdate is the name of last update collection.
	CollectionNameLastUpdate = "last_update"
	// CollectionIndexValidatorsByEmergency is the name of validators by emergency index.
	CollectionIndexValidatorsByEmergency = "validators_by_emergency"
	// CollectionIndexValidatorsByOperatorAddress is the name of validators by operator address index.
	CollectionIndexValidatorsByOperatorAddress = "validators_by_operator_address"
)

var (
	// ParamsKey is the key for params.
	ParamsKey = []byte("p_restaking")
	// KeyPrefixValidator is the prefix for validator keys.
	KeyPrefixValidator = collections.NewPrefix(CollectionNameValidators)
	// KeyPrefixLastUpdate is the prefix for last update keys.
	KeyPrefixLastUpdate = collections.NewPrefix(CollectionNameLastUpdate)
)

// KeyPrefix returns the key prefix.
func KeyPrefix(p string) []byte {
	return []byte(p)
}

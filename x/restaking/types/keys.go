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
)

var (
	// ParamsKey is the key for params.
	ParamsKey = collections.NewPrefix(0)
	// KeyPrefixValidators is the prefix for validator keys.
	KeyPrefixValidators = collections.NewPrefix(1)
	// KeyPrefixLastUpdate is the prefix for last update keys.
	KeyPrefixLastUpdate = collections.NewPrefix(2)
	// KeyPrefixPendingOperators is the prefix for pending validators keys.
	KeyPrefixPendingOperators = collections.NewPrefix(3)

	// KeyPrefixBondedValidators is the prefix for bonded validators keys.
	KeyPrefixBondedValidators = collections.NewPrefix(4)
	// KeyPrefixEmergencyValidators is the prefix for emergency validators keys.
	KeyPrefixEmergencyValidators = collections.NewPrefix(5)
	// KeyPrefixEvmAddressValidators is the prefix for EVM address pending validators keys.
	KeyPrefixEvmAddressValidators = collections.NewPrefix(6)
	// KeyPrefixDeltaUpdates is the prefix for validator updates keys.
	KeyPrefixDeltaUpdates = collections.NewPrefix(7)
)

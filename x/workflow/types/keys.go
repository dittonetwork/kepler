package types

const (
	// ModuleName defines the module name.
	ModuleName = "workflow"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_workflow"

	KeyPrefixAutomation  = "automation:"
	KeyActiveAutomations = "active_automations"
)

var (
	ParamsKey = []byte("p_workflow")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

package types

type OperatorStatus string

const (
	OperatorStatusBonded    OperatorStatus = "bonded"
	OperatorStatusUnbonded  OperatorStatus = "unbonded"
	OperatorStatusUnbonding OperatorStatus = "unbonding"
)

type Operator struct {
	// Address of the operator from L1 chain
	Address string `json:"address"`

	// PublicKey of the operator from L1 chain (can be rotated)
	PublicKey string `json:"public_key"`

	// Status of the operator
	Status OperatorStatus `json:"status"`

	// IsEmergency is true if the operator is in emergency mode
	IsEmergency bool `json:"is_emergency"`

	// Tokens is the amount of tokens staked by the operator
	Tokens uint64 `json:"tokens"`
}

type UpdateValidatorSetParams struct {
	Operators   []Operator `json:"validator_set"`
	BlockHeight int64      `json:"block_height"`
	BlockHash   []byte     `json:"block_hash"`
}

// RestakingValidator is a struct that contains the additional information of a validator
type RestakingValidator struct {
	OperatorAddress string `json:"operator_address"`
	IsEmergency     bool   `json:"is_emergency"`
}

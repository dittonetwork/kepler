package types

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// OperatorStatus defines the status of an operator.
type OperatorStatus string

const (
	// OperatorStatusBonded represents a bonded operator.
	OperatorStatusBonded OperatorStatus = "bonded"
	// OperatorStatusUnbonded represents an unbonded operator.
	OperatorStatusUnbonded OperatorStatus = "unbonded"
	// OperatorStatusUnbonding represents an unbonding operator.
	OperatorStatusUnbonding OperatorStatus = "unbonding"
)

// ToStakingBondStatus converts restaking operator status to staking bond status.
// This mapping is used when updating the staking module with validator information.
func (x OperatorStatus) ToStakingBondStatus() stakingtypes.BondStatus {
	switch x {
	case OperatorStatusBonded:
		return stakingtypes.Bonded
	case OperatorStatusUnbonded:
		return stakingtypes.Unbonded
	case OperatorStatusUnbonding:
		return stakingtypes.Unbonding
	default:
		return stakingtypes.Unspecified
	}
}

// ToRestakingValidatorStatus converts operator status to restaking validator status
// This mapping is used when saving validator information to the restaking module's store
// Note: isNew is a flag that indicates if the operator is a new validator.
// This flag is set to true when the operator is a new validator coming from L1 chain.
func (x OperatorStatus) ToRestakingValidatorStatus(isNew bool) ValidatorStatus {
	switch x {
	case OperatorStatusBonded:
		if isNew {
			return ValidatorStatus_VALIDATOR_STATUS_BONDING
		}

		return ValidatorStatus_VALIDATOR_STATUS_BONDED
	case OperatorStatusUnbonded:
		return ValidatorStatus_VALIDATOR_STATUS_UNBONDED
	case OperatorStatusUnbonding:
		return ValidatorStatus_VALIDATOR_STATUS_UNBONDING
	default:
		return ValidatorStatus_VALIDATOR_STATUS_UNSPECIFIED
	}
}

// Operator represents an operator from the L1 chain.
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

// UpdateValidatorSetParams defines the parameters for updating the validator set.
type UpdateValidatorSetParams struct {
	Operators   []Operator `json:"validator_set"`
	BlockHeight int64      `json:"block_height"`
	BlockHash   []byte     `json:"block_hash"`
}

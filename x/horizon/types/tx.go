package types

import (
	"fmt"
	"math/big"
)

func (m *MsgScheduleAutomationParams) Validate() error {
	// Trigger
	if m.TriggerExpressionVersion != 1 {
		return fmt.Errorf(`trigger expression version [%d] is invalid`, m.TriggerExpressionVersion)
	}

	// TODO: Pre-validate trigger expression

	activeTargetsCount := len(m.ActionTargets)
	actionTargetCallsCount := len(m.ActionTargetCalls)

	if activeTargetsCount == 0 {
		return fmt.Errorf(`automation should contain at least one action`)
	}

	if activeTargetsCount != actionTargetCallsCount {
		return fmt.Errorf(`action targets [%d] and action calls [%d] count mismatch`, activeTargetsCount, actionTargetCallsCount)
	}

	preconditionCheckersCount := len(m.PreconditionCheckers)
	preconditionCheckCallsCount := len(m.PreconditionCheckCalls)

	if len(m.PreconditionCheckers) != len(m.PreconditionCheckCalls) {
		return fmt.Errorf(`pre-condition checkers [%d] and precondition check [%d] calls count mismatch`, preconditionCheckersCount, preconditionCheckCallsCount)
	}

	// Time range
	activeFrom, parseActiveFromSuccess := big.NewInt(0).SetString(m.ActiveFrom, 10)
	if !parseActiveFromSuccess {
		return fmt.Errorf(`invalid value for ActiveFrom: '%s' is not a valid integer`, m.ActiveFrom)
	}

	if activeFrom.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf(`invalid value for ActiveFrom [%s]: it should be greater than zero`, m.ActiveFrom)
	}

	activeTo, parseActiveToSuccess := big.NewInt(0).SetString(m.ActiveTo, 10)
	if !parseActiveToSuccess {
		return fmt.Errorf(`invalid value for ActiveTo: '%s' is not a valid integer`, m.ActiveTo)
	}

	if activeTo.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf(`invalid value for ActiveTo [%s]: it should be greater than zero`, m.ActiveTo)
	}

	if activeTo.Cmp(activeFrom) <= 0 {
		return fmt.Errorf(`invalid range: ActiveTo (%s) must be greater than ActiveFrom [%s]`, m.ActiveTo, m.ActiveFrom)
	}

	// Execution conditions
	_, parseTargetChainIdSuccess := big.NewInt(0).SetString(m.TargetChainId, 10)
	if !parseTargetChainIdSuccess {
		return fmt.Errorf(`invalid value for TargetChainId: '%s' is not a valid integer`, m.TargetChainId)
	}

	// TODO: Validate target chain id

	_, parseMaxGasPriceSuccess := big.NewInt(0).SetString(m.MaxGasPrice, 10)
	if !parseMaxGasPriceSuccess {
		return fmt.Errorf(`invalid value for MaxGasPrice: '%s' is not a valid integer`, m.MaxGasPrice)
	}

	// TODO: Validate max gas price if necessary

	if m.MaxNumberOfExecutions < 0 {
		return fmt.Errorf(`invalid value for MaxNumberOfExecutions [%d]: it should be greater than zero`, m.MaxNumberOfExecutions)
	}

	// TODO: Define max value for MaxNumberOfExecutions, probably via module params
	// if m.MaxNumberOfExecutions > 0 {
	// 	 return fmt.Errorf(`invalid value for MaxNumberOfExecutions [%d]: it should be lower than X`, m.MaxNumberOfExecutions)
	// }

	// Signature

	// TODO: Validate signer and signature

	// Positive scenario

	return nil
}

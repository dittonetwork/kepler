package types

import (
	"context"
	"errors"
)

type EpochHooks interface {
	// AfterEpochEnd the first block whose timestamp is after the duration is counted as the end of the epoch.
	AfterEpochEnd(ctx context.Context, epochID string, epochNumber int64) error

	// BeforeEpochStart new epoch is next block of epoch end block.
	BeforeEpochStart(ctx context.Context, epochID string, epochNumber int64) error
}

var _ EpochHooks = MultiEpochHooks{}

type MultiEpochHooks []EpochHooks

func NewMultiEpochHooks(hooks ...EpochHooks) MultiEpochHooks {
	return hooks
}

// AfterEpochEnd is called when epoch is going to be ended, epochNumber is the number of epoch that is ending.
func (m MultiEpochHooks) AfterEpochEnd(ctx context.Context, epochID string, epochNumber int64) error {
	var errs error

	for i := range m {
		errs = errors.Join(errs, m[i].AfterEpochEnd(ctx, epochID, epochNumber))
	}

	return errs
}

// BeforeEpochStart is called when epoch is going to be started, epochNumber is the number of epoch that is starting.
func (m MultiEpochHooks) BeforeEpochStart(ctx context.Context, epochID string, epochNumber int64) error {
	var errs error

	for i := range m {
		errs = errors.Join(errs, m[i].BeforeEpochStart(ctx, epochID, epochNumber))
	}

	return errs
}

// EpochHooksWrapper is a wrapper for modules to inject EpochHooks using depinject.
type EpochHooksWrapper struct{ EpochHooks }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (EpochHooksWrapper) IsOnePerModuleType() {}

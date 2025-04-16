package types

import (
	"context"
	"errors"
)

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

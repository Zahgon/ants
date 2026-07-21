package ants

import (
	"context"
	"time"
)

type LoadBalancingStrategy int

const (
	RoundRobin LoadBalancingStrategy = 1 << (iota + 1)

	LeastTasks
)

type contextReleaser interface {
	ReleaseContext(ctx context.Context) error
}

func releasePools(ctx context.Context, pools []contextReleaser) error {
	_ = "STUB: not implemented"
	return nil
}

type MultiPool struct {
	pools []*Pool
	index uint32
	state int32
	lbs   LoadBalancingStrategy
}

func NewMultiPool(size, sizePerPool int, lbs LoadBalancingStrategy, options ...Option) (*MultiPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mp *MultiPool) next(lbs LoadBalancingStrategy) (idx int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPool) Submit(task func()) (err error) { _ = "STUB: not implemented"; return nil }

func (mp *MultiPool) Running() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPool) RunningByIndex(idx int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (mp *MultiPool) Free() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPool) FreeByIndex(idx int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (mp *MultiPool) Waiting() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPool) WaitingByIndex(idx int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (mp *MultiPool) Cap() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPool) Tune(size int) { _ = "STUB: not implemented"; return }

func (mp *MultiPool) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (mp *MultiPool) ReleaseTimeout(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *MultiPool) ReleaseContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *MultiPool) Reboot() { _ = "STUB: not implemented"; return }

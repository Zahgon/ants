package ants

import (
	"context"
	"time"
)

type MultiPoolWithFunc struct {
	pools []*PoolWithFunc
	index uint32
	state int32
	lbs   LoadBalancingStrategy
}

func NewMultiPoolWithFunc(size, sizePerPool int, fn func(any), lbs LoadBalancingStrategy, options ...Option) (*MultiPoolWithFunc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mp *MultiPoolWithFunc) next(lbs LoadBalancingStrategy) (idx int) {
	_ = "STUB: not implemented"
	return 0
}

func (mp *MultiPoolWithFunc) Invoke(args any) (err error) { _ = "STUB: not implemented"; return nil }

func (mp *MultiPoolWithFunc) Running() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPoolWithFunc) RunningByIndex(idx int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mp *MultiPoolWithFunc) Free() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPoolWithFunc) FreeByIndex(idx int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mp *MultiPoolWithFunc) Waiting() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPoolWithFunc) WaitingByIndex(idx int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mp *MultiPoolWithFunc) Cap() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPoolWithFunc) Tune(size int) { _ = "STUB: not implemented"; return }

func (mp *MultiPoolWithFunc) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (mp *MultiPoolWithFunc) ReleaseTimeout(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *MultiPoolWithFunc) ReleaseContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *MultiPoolWithFunc) Reboot() { _ = "STUB: not implemented"; return }

package ants

import (
	"context"
	"time"
)

type MultiPoolWithFuncGeneric[T any] struct {
	pools []*PoolWithFuncGeneric[T]
	index uint32
	state int32
	lbs   LoadBalancingStrategy
}

func NewMultiPoolWithFuncGeneric[T any](size, sizePerPool int, fn func(T), lbs LoadBalancingStrategy, options ...Option) (*MultiPoolWithFuncGeneric[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mp *MultiPoolWithFuncGeneric[T]) next(lbs LoadBalancingStrategy) (idx int) {
	_ = "STUB: not implemented"
	return 0
}

func (mp *MultiPoolWithFuncGeneric[T]) Invoke(args T) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (mp *MultiPoolWithFuncGeneric[T]) Running() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPoolWithFuncGeneric[T]) RunningByIndex(idx int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mp *MultiPoolWithFuncGeneric[T]) Free() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPoolWithFuncGeneric[T]) FreeByIndex(idx int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mp *MultiPoolWithFuncGeneric[T]) Waiting() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPoolWithFuncGeneric[T]) WaitingByIndex(idx int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mp *MultiPoolWithFuncGeneric[T]) Cap() (n int) { _ = "STUB: not implemented"; return 0 }

func (mp *MultiPoolWithFuncGeneric[T]) Tune(size int) { _ = "STUB: not implemented"; return }

func (mp *MultiPoolWithFuncGeneric[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (mp *MultiPoolWithFuncGeneric[T]) ReleaseTimeout(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *MultiPoolWithFuncGeneric[T]) ReleaseContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *MultiPoolWithFuncGeneric[T]) Reboot() { _ = "STUB: not implemented"; return }

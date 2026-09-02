package ants

type goWorkerWithFuncGeneric[T any] struct {
	worker

	pool *PoolWithFuncGeneric[T]

	arg chan T

	exit chan struct{}

	lastUsed int64
}

func (w *goWorkerWithFuncGeneric[T]) run() { _ = "STUB: not implemented"; return }

func (w *goWorkerWithFuncGeneric[T]) finish() { _ = "STUB: not implemented"; return }

func (w *goWorkerWithFuncGeneric[T]) lastUsedTime() int64 { _ = "STUB: not implemented"; return 0 }

func (w *goWorkerWithFuncGeneric[T]) setLastUsedTime(t int64) { _ = "STUB: not implemented"; return }

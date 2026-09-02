package ants

type goWorkerWithFunc struct {
	worker

	pool *PoolWithFunc

	arg chan any

	lastUsed int64
}

func (w *goWorkerWithFunc) run() { _ = "STUB: not implemented"; return }

func (w *goWorkerWithFunc) finish() { _ = "STUB: not implemented"; return }

func (w *goWorkerWithFunc) lastUsedTime() int64 { _ = "STUB: not implemented"; return 0 }

func (w *goWorkerWithFunc) setLastUsedTime(t int64) { _ = "STUB: not implemented"; return }

func (w *goWorkerWithFunc) inputArg(arg any) { _ = "STUB: not implemented"; return }

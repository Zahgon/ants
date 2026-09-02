package ants

type goWorker struct {
	worker

	pool *Pool

	task chan func()

	lastUsed int64
}

func (w *goWorker) run() { _ = "STUB: not implemented"; return }

func (w *goWorker) finish() { _ = "STUB: not implemented"; return }

func (w *goWorker) lastUsedTime() int64 { _ = "STUB: not implemented"; return 0 }

func (w *goWorker) setLastUsedTime(t int64) { _ = "STUB: not implemented"; return }

func (w *goWorker) inputFunc(fn func()) { _ = "STUB: not implemented"; return }

package ants

import "time"

type loopQueue struct {
	items  []worker
	expiry []worker
	head   int
	tail   int
	size   int
	isFull bool
}

func newWorkerLoopQueue(size int) *loopQueue { _ = "STUB: not implemented"; return nil }

func (wq *loopQueue) len() int { _ = "STUB: not implemented"; return 0 }

func (wq *loopQueue) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (wq *loopQueue) insert(w worker) error { _ = "STUB: not implemented"; return nil }

func (wq *loopQueue) detach() worker { _ = "STUB: not implemented"; return *new(worker) }

func (wq *loopQueue) refresh(duration time.Duration) []worker {
	_ = "STUB: not implemented"
	return nil
}

func (wq *loopQueue) binarySearch(expiryTime int64) int { _ = "STUB: not implemented"; return 0 }

func (wq *loopQueue) reset() { _ = "STUB: not implemented"; return }

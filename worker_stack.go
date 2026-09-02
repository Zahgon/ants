package ants

import "time"

type workerStack struct {
	items  []worker
	expiry []worker
}

func newWorkerStack(size int) *workerStack { _ = "STUB: not implemented"; return nil }

func (ws *workerStack) len() int { _ = "STUB: not implemented"; return 0 }

func (ws *workerStack) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (ws *workerStack) insert(w worker) error { _ = "STUB: not implemented"; return nil }

func (ws *workerStack) detach() worker { _ = "STUB: not implemented"; return *new(worker) }

func (ws *workerStack) refresh(duration time.Duration) []worker {
	_ = "STUB: not implemented"
	return nil
}

func (ws *workerStack) binarySearch(l, r int, expiryTime int64) int {
	_ = "STUB: not implemented"
	return 0
}

func (ws *workerStack) reset() { _ = "STUB: not implemented"; return }

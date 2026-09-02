package ants

import (
	"errors"
	"time"
)

var errQueueIsFull = errors.New("the queue is full")

type worker interface {
	run()
	finish()
	lastUsedTime() int64
	setLastUsedTime(t int64)
	inputFunc(func())
	inputArg(any)
}

type workerQueue interface {
	len() int
	isEmpty() bool
	insert(worker) error
	detach() worker
	refresh(duration time.Duration) []worker
	reset()
}

type queueType int

const (
	queueTypeStack queueType = 1 << iota
	queueTypeLoopQueue
)

func newWorkerQueue(qType queueType, size int) workerQueue {
	_ = "STUB: not implemented"
	return *new(workerQueue)
}

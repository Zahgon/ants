package ants

import (
	"context"
	"errors"
	"log"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	DefaultAntsPoolSize = math.MaxInt32

	DefaultCleanIntervalTime = time.Second
)

const (
	OPENED = iota

	CLOSED
)

var (
	ErrLackPoolFunc = errors.New("must provide function for pool")

	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")

	ErrPoolClosed = errors.New("this pool has been closed")

	ErrPoolOverload = errors.New("too many goroutines blocked on submit or Nonblocking is set")

	ErrInvalidPreAllocSize = errors.New("can not set up a negative capacity under PreAlloc mode")

	ErrTimeout = errors.New("operation timed out")

	ErrInvalidPoolIndex = errors.New("invalid pool index")

	ErrInvalidLoadBalancingStrategy = errors.New("invalid load-balancing strategy")

	ErrInvalidMultiPoolSize = errors.New("invalid size for multiple pool")

	workerChanCap = func() int {

		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}

		return 1
	}()

	defaultLogger = Logger(log.New(os.Stderr, "[ants]: ", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds))

	defaultAntsPool, _ = NewPool(DefaultAntsPoolSize)
)

func Submit(task func()) error { _ = "STUB: not implemented"; return nil }

func Running() int { _ = "STUB: not implemented"; return 0 }

func Cap() int { _ = "STUB: not implemented"; return 0 }

func Free() int { _ = "STUB: not implemented"; return 0 }

func Release() { _ = "STUB: not implemented"; return }

func ReleaseTimeout(timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

func ReleaseContext(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func Reboot() { _ = "STUB: not implemented"; return }

type Logger interface {
	Printf(format string, args ...any)
}

type poolCommon struct {
	capacity int32

	running int32

	lock sync.Locker

	workers workerQueue

	state int32

	cond *sync.Cond

	allDone chan struct{}

	once *sync.Once

	workerCache sync.Pool

	waiting int32

	purgeDone int32
	purgeCtx  context.Context
	stopPurge context.CancelFunc

	ticktockDone int32
	ticktockCtx  context.Context
	stopTicktock context.CancelFunc

	now int64

	options *Options
}

func newPool(size int, options ...Option) (*poolCommon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *poolCommon) purgeStaleWorkers() { _ = "STUB: not implemented"; return }

const nowTimeUpdateInterval = 500 * time.Millisecond

func (p *poolCommon) ticktock() { _ = "STUB: not implemented"; return }

func (p *poolCommon) goPurge() { _ = "STUB: not implemented"; return }

func (p *poolCommon) goTicktock() { _ = "STUB: not implemented"; return }

func (p *poolCommon) nowTime() int64 { _ = "STUB: not implemented"; return 0 }

func (p *poolCommon) Running() int { _ = "STUB: not implemented"; return 0 }

func (p *poolCommon) Free() int { _ = "STUB: not implemented"; return 0 }

func (p *poolCommon) Waiting() int { _ = "STUB: not implemented"; return 0 }

func (p *poolCommon) Cap() int { _ = "STUB: not implemented"; return 0 }

func (p *poolCommon) Tune(size int) { _ = "STUB: not implemented"; return }

func (p *poolCommon) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (p *poolCommon) Release() { _ = "STUB: not implemented"; return }

func (p *poolCommon) ReleaseTimeout(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *poolCommon) ReleaseContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *poolCommon) Reboot() { _ = "STUB: not implemented"; return }

func (p *poolCommon) addRunning(delta int) int { _ = "STUB: not implemented"; return 0 }

func (p *poolCommon) addWaiting(delta int) { _ = "STUB: not implemented"; return }

func (p *poolCommon) retrieveWorker() (w worker, err error) {
	_ = "STUB: not implemented"
	return *new(worker), nil
}

func (p *poolCommon) revertWorker(worker worker) bool { _ = "STUB: not implemented"; return false }

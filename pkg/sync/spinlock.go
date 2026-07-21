package sync

import (
	"sync"
)

type spinLock uint32

const maxBackoff = 16

func (sl *spinLock) Lock() { _ = "STUB: not implemented"; return }

func (sl *spinLock) Unlock() { _ = "STUB: not implemented"; return }

func NewSpinLock() sync.Locker { _ = "STUB: not implemented"; return *new(sync.Locker) }

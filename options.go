package ants

import "time"

type Option func(opts *Options)

func loadOptions(options ...Option) *Options { _ = "STUB: not implemented"; return nil }

type Options struct {
	ExpiryDuration time.Duration

	PreAlloc bool

	MaxBlockingTasks int

	Nonblocking bool

	PanicHandler func(any)

	Logger Logger

	DisablePurge bool
}

func WithOptions(options Options) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithExpiryDuration(expiryDuration time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPreAlloc(preAlloc bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxBlockingTasks(maxBlockingTasks int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithNonblocking(nonblocking bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPanicHandler(panicHandler func(any)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLogger(logger Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDisablePurge(disable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

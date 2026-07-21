package ants

type Pool struct {
	*poolCommon
}

func (p *Pool) Submit(task func()) error { _ = "STUB: not implemented"; return nil }

func NewPool(size int, options ...Option) (*Pool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

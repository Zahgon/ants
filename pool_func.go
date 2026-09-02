package ants

type PoolWithFunc struct {
	*poolCommon

	fn func(any)
}

func (p *PoolWithFunc) Invoke(arg any) error { _ = "STUB: not implemented"; return nil }

func NewPoolWithFunc(size int, pf func(any), options ...Option) (*PoolWithFunc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

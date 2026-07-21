package ants

type PoolWithFuncGeneric[T any] struct {
	*poolCommon

	fn func(T)
}

func (p *PoolWithFuncGeneric[T]) Invoke(arg T) error { _ = "STUB: not implemented"; return nil }

func NewPoolWithFuncGeneric[T any](size int, pf func(T), options ...Option) (*PoolWithFuncGeneric[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

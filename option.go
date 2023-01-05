package myhig

type Option[T any] struct {
	v *T
}

func Some[T any](v T) *Option[T] {
	nv := v
    return &Option[T]{v:&nv}
}

func None[T any]() *Option[T] {
	return &Option[T]{v: nil}
}

func (o *Option[T]) IsNone() bool {
	return o.v == nil
}

func (o *Option[T]) Unwrap() (T, bool) {
	return *o.v, o.v != nil
}

func (o *Option[T]) SetSome(v T) {
    nv := v
	o.v = &nv
}

func (o *Option[T]) SetNone() {
	o.v = nil
}

func (o *Option[T]) setFail(err error) {
    if err != nil {
        o.v = nil
    }
}

func (o *Option[T]) MustOrReturnTo(ra failReturnAble) T {
    if o.v == nil {
        ra.ErrorReturn(ErrIsFalse)
    }
    return *o.v
}

func (o *Option[T]) MustOrDoReturnTo(fn func(error), ra failReturnAble) T {
    if o.v == nil {
        fn(ErrIsFalse)
        ra.ErrorReturn(ErrIsFalse)
    }
    return *o.v
}

func (o *Option[T]) MustOrFunc(fn func() T ) T {
	if o.v == nil {
		return fn()
	}
	return *o.v
}

func (o *Option[T]) MustOr(v T) T {
	if o.v == nil {
        return v
	}
	return *o.v
}


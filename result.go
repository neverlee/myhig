package myhig

type Result[T any] struct {
	v T
	err error
}

func (r *Result[T]) IsOK() bool {
	return r.err == nil
}

func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

func (r *Result[T]) Expect(msg string) T {
	if r.err != nil {
		panic(msg)
	}
	return r.v
}

func (r *Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.v
}

func (r *Result[T]) UnwrapOr(v T) T {
	if r.err != nil {
		return v
	}
	return r.v
}

func (r *Result[T]) UnwrapOrDefault() T {
	if r.err != nil {
		var v T
		return v
	}
	return r.v
}

func (r *Result[T]) UnwrapOrElse(fn func() T) T {
	if r.err != nil {
		return fn()
	}
	return r.v
}

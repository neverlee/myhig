package myhig

type Result[T any] struct {
	v T
	err error
}

func OK[T any](v T) *Result[T] {
    return &Result[T]{v:v, err:nil}
}

func Err[T any](e error) *Result[T] {
    return &Result[T]{err:e}
}

func (r *Result[T]) IsOK() bool {
	return r.err == nil
}

func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

func (r *Result[T]) Unwrap() (T, error) {
	return r.v, r.err
}

func (r *Result[T]) SetOK(v T) {
    r.v = v
    r.err = nil
}

func (r *Result[T]) SetErr(err error) {
    if err != nil {
        var v T
        r.v = v
        r.err = err
    }
}

func (r *Result[T]) setFail(err error) {
    r.SetErr(err)
}

func (r *Result[T]) MustOrReturnTo(ra failReturnAble) T {
    if r.err != nil {
        ra.ErrorReturn(r.err)
    }
    return r.v
}

func (r *Result[T]) MustOrDoReturnTo(fn func(error), ra failReturnAble) T {
    if r.err != nil {
        fn(r.err)
        ra.ErrorReturn(r.err)
    }
    return r.v
}

func (r *Result[T]) MustOrFunc(fn func() T ) T {
	if r.err != nil {
		return fn()
	}
	return r.v
}

func (r *Result[T]) MustOr(v T) T {
	if r.err != nil {
        return v 
	}
	return r.v
}

// func (r *Result[T]) Expect(msg string) T {
// 	if r.err != nil {
// 		panic(msg)
// 	}
// 	return r.v
// }
// 
// func (r *Result[T]) Unwrap() T {
// 	if r.err != nil {
// 		panic(r.err)
// 	}
// 	return r.v
// }
// 
// func (r *Result[T]) UnwrapOr(v T) T {
// 	if r.err != nil {
// 		return v
// 	}
// 	return r.v
// }
// 
// func (r *Result[T]) UnwrapOrDefault() T {
// 	if r.err != nil {
// 		var v T
// 		return v
// 	}
// 	return r.v
// }
// 
// func (r *Result[T]) UnwrapOrElse(fn func() T) T {
// 	if r.err != nil {
// 		return fn()
// 	}
// 	return r.v
// }

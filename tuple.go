package myhig

var Unit = Tuple0{}
type Tuple0 struct {}

type Tuple1 [T1 any]struct {
	V1 T1
}

func Tw1[T1 any](v1 T1) Tuple1[T1] {
	return Tuple1[T1]{v1}
}

func (t *Tuple1[T1]) Unwrap() T1 {
	return t.V1
}

func (t *Tuple1[T1]) From(v1 T1) returnDone {
	t.V1 = v1
	return returnDone{}
}


func (t *Tuple1[T1]) Must(ra failReturnAble) T1 {
	tupleMust(ra, t.V1, nil)
	return t.V1
}

func (t *Tuple1[T1]) MustFunc(ra failReturnAble, fn func ()) T1 {
	tupleMust(ra, t.V1, fn)
	return t.V1
}

func (t *Tuple1[T1]) setLast(err error) {
	setError(&t.V1, err)
}


type Tuple2 [T1, T2 any]struct {
	V1 T1
	V2 T2
}

func Tw2[T1 any, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{v1, v2}
}

func (t *Tuple2[T1, T2]) Unwrap() (T1, T2) {
	return t.V1, t.V2
}

func (t *Tuple2[T1, T2]) From(v1 T1, v2 T2) returnDone {
	t.V1 = v1
	t.V2 = v2
	return returnDone{}
}

func (t *Tuple2[T1, T2]) Must(ra failReturnAble) (T1, T2) {
	tupleMust(ra, t.V2, nil)
	return t.V1, t.V2
}

func (t *Tuple2[T1, T2]) MustFunc(ra failReturnAble, fn func ()) (T1, T2) {
	tupleMust(ra, t.V2, fn)
	return t.V1, t.V2
}

func (t *Tuple2[T1, T2]) setLast(err error) {
	setError(&t.V2, err)
}


func setError[T any](v *T, err error) {
	switch v := any(v).(type) {
	case *error:
		(*v) = err
	case *bool:
		(*v) = err == nil
	default:
		panic("last type must be error or bool")
	}
}


type failReturnAble interface {
	ErrorReturn(err error)
	FalseReturn(b bool)
}

func tupleMust[T any](ra failReturnAble, v T, fn func()) {
	switch vt := any(v).(type) {
	case error:
		if vt != nil {
			if fn != nil {
				fn()
			}
			ra.ErrorReturn(vt)
		}
	case bool:
		if !vt {
			if fn != nil {
				fn()
			}
			ra.FalseReturn(vt)
		}
	default:
		panic("last type must be error or bool")
	}
}

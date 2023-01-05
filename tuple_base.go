package myhig

var Unit = Tuple0{}
type Tuple0 struct {}

func Tw0() Tuple0 {
	return Unit
}

func (t *Tuple0) Unwrap() {
}

func (t *Tuple0) Set() {
}

func (t *Tuple0) Must(ra failReturnAble) () {
	panic("you should not use this func")
}

func (t *Tuple0) MustFunc(ra failReturnAble, fn func ()) () {
	panic("you should not use this func")
}

func (t *Tuple0) setFail(err error) {
	panic("you should not use this func")
}

type RetProc0 struct {
	Tuple0
	Proc
}

func NewRetProc0() *RetProc0 {
	return &RetProc0{}
}

func (rt *RetProc0) ErrorReturn(err error) {
	if err != nil {
		rt.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc0) FalseReturn(b bool) {
	if !b {
		rt.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc0) Return() {
	rt.ReturnOnly()
}

func (rt *RetProc0) IfReturn(b bool) {
	if b {
		rt.ReturnOnly()
	}
}

func (rt *RetProc0) Dov(fn func()) *Tuple0 {
	rt.Do(fn)
	return &rt.Tuple0
}

func (rt *RetProc0) Dow(fn func()) {
	rt.Do(fn)
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

func tupleMust[T any](ra failReturnAble, v T, fn func(err error)) {
	switch vt := any(&v).(type) {
	case *error:
	if *vt != nil {
		if fn != nil {
			fn(*vt)
		}
		ra.ErrorReturn(*vt)
	}
	case *bool:
	if !*vt {
		if fn != nil {
			fn(ErrIsFalse)
		}
		ra.FalseReturn(*vt)
	}
	default:
		panic("last type must be error or bool")
	}
}

func isFail[T any](v T) bool {
	switch vt := any(&v).(type) {
	case *error:
	if *vt != nil {
		return true
	}
	case *bool:
	if !*vt {
		return true
	}
	default:
		panic("last type must be error or bool")
	}
	return false
}

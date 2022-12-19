package myhig

import (
	"errors"
)

type Myhig [T any]struct {
	V T
	returned bool
}

func New[T any](p *T) *Myhig[T] {
	return &Myhig[T]{}
}

func (m *Myhig[T]) Recover() {
	if m.returned {
		recover()
		m.returned = false
	}
}

func (m *Myhig[T]) VReturn(d returnDone) {
	m.returned = true
	panic(d)
}

func (m *Myhig[T]) TReturn(r T) {
	m.V = r
	m.returned = true
	panic(returnDone{})
}

func (m *Myhig[T]) VIfReturn(cond bool, d returnDone) {
	if cond {
		m.VReturn(d)
	}
}

func (m *Myhig[T]) TIfReturn(cond bool, r T) {
	if cond {
		m.TReturn(r)
	}
}

var errFalse = errors.New("ErrFalse")

type setLastor interface{
	setLast(error)
}

func (m *Myhig[T]) ErrorReturn(err error) {
	if err != nil {
		if sl, ok := any(&m.V).(setLastor); ok {
			sl.setLast(err)
			m.VReturn(returnDone{})
		} else {
			panic("last value type is not error")
		}
	}
}

func (m *Myhig[T]) FalseReturn(b bool) {
	if !b {
		if sl, ok := any(&m.V).(setLastor); ok {
			sl.setLast(errFalse)
			m.VReturn(returnDone{})
		} else {
			panic("last value type is not bool")
		}
	}
}

func (m *Myhig[T]) Do(fn func ()) (ret *T) {
	ret = &m.V
	defer m.Recover()
	fn()
	return &m.V
}



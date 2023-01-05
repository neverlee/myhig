package myhig

import (
	"errors"
)

var ErrIsFalse = errors.New("ErrFalse")

type setFailer interface{
	setFail(error)
}

type Proc struct {
	returned bool
}

func (m *Proc) RecoverTheReturn() {
	if m.returned {
		recover()
		m.returned = false
	}
}

func (m *Proc) ReturnOnly() {
	m.returned = true
	panic(nil)
}

func (m *Proc) ReturnWith(fn func()) {
	if fn != nil {
		fn()
	}
	m.returned = true
	panic(nil)
}

func (m *Proc) IfReturnWith(cond bool, fn func()) {
	if cond {
		if fn != nil {
			fn()
		}
		m.ReturnOnly()
	}
}

func (m *Proc) Do(fn func ()) {
	defer m.RecoverTheReturn()
	fn()
}


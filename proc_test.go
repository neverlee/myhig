package myhig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Proc_Return(t *testing.T) {
	p := Proc{}
	assert.NotPanics(t, func() {
		p.Do(func() {
			p.ReturnOnly()
		})
	})

	assert.NotPanics(t, func() {
		p.Do(func() {
			p.ReturnWith(nil)
		})
	})

	assert.NotPanics(t, func() {
		p.Do(func() {
			p.IfReturnWith(true, nil)
		})
	})
}


func Test_Proc_Panic(t *testing.T) {
	p := Proc{}
	assert.Panics(t, func() {
		p.Do(func() {
			panic("panic")
		})
	})
}

func Test_Proc_ReturnFunc(t *testing.T) {
	p := Proc{}
	a := 9
	p.Do(func() {
		p.ReturnWith(func() {
			a = 100
		})
	})
	assert.Equal(t, 100, a)

	p.Do(func() {
		p.IfReturnWith(true, func() {
			a = 200
		})
		a = 250
	})
	assert.Equal(t, 200, a)

	p.Do(func() {
		p.IfReturnWith(false, func() {
			a = 300
		})
		a = 350
	})
	assert.Equal(t, 350, a)
}


package myhig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RetPorc(t *testing.T) {
	var f func(a int, b int) (int, error)
	rp := GetProcOfFunc2r2(f)
	{
		ri, re := rp.Dov(func() {
			rp.Return(9, nil)
		}).Unwrap()

		assert.Equal(t, ri, 9)
		assert.NoError(t, re)
	}

	{
		ri, re := rp.Dow(func() {
			rp.IfReturn(false, 10, fmt.Errorf("x"))
		})

		assert.Equal(t, ri, 9)
		assert.NoError(t, re)
	}

	{
		ri, re := rp.Dow(func() {
			rp.IfReturn(true, 10, fmt.Errorf("x"))
		})

		assert.Equal(t, ri, 10)
		assert.Error(t, re)
	}
}

func Test_RetPorc_ErrorReturn(t *testing.T) {
	var f func(a int, b int) (int, error)
	{
		rp := GetProcOfFunc2r2(f)
		k := 0
		ri, re := rp.Dow(func() {
			rp.ErrorReturn(nil)
			k = 9
		})

		assert.Equal(t, ri, 0)
		assert.NoError(t, re)
		assert.Equal(t, k, 9)
	}

	{
		rp := GetProcOfFunc2r2(f)
		k := 0
		ri, re := rp.Dow(func() {
			rp.ErrorReturn(fmt.Errorf("x"))
			k = 9
		})

		assert.Equal(t, ri, 0)
		assert.Error(t, re)
		assert.Equal(t, k, 0)
	}
}

func Test_RetPorc_FalseReturn(t *testing.T) {
	var f func(a int, b int) (int, error)
	{
		rp := GetProcOfFunc2r2(f)
		k := 0
		ri, re := rp.Dow(func() {
			rp.FalseReturn(true)
			k = 9
		})

		assert.Equal(t, ri, 0)
		assert.NoError(t, re)
		assert.Equal(t, k, 9)
	}

	{
		rp := GetProcOfFunc2r2(f)
		k := 0
		ri, re := rp.Dow(func() {
			rp.FalseReturn(false)
			k = 9
		})

		assert.Equal(t, ri, 0)
		assert.Error(t, re)
		assert.Equal(t, k, 0)
	}
}

func Test_RetProc_Panic(t *testing.T) {
	p := Proc{}
	assert.Panics(t, func() {
		p.Do(func() {
			panic("panic")
		})
	})
	var f func(a int) (int, error)
	{
		rp := GetProcOfFunc1r2(f)
		k := 0
		assert.Panics(t, func() {
			rp.Dow(func() {
				defer func() {
					k = 9
				}()
				panic("hello")
			})
		})

		assert.Equal(t, k, 9)
	}
}

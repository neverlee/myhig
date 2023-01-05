package myhig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RetTuple2(t *testing.T) {
	fmt.Println("test RetTuple")
	{
		m := Must2(9, false)
		assert.Equal(t, 10, m.Or(10))
		assert.Equal(t, 10, m.OrFunc(func() int {
			return 10
		}))
	}

	{
		m := Must2(9, true)
		assert.Equal(t, 9, m.Or(10))
		assert.Equal(t, 9, m.OrFunc(func() int {
			return 10
		}))
	}
}

func Test_RetTupleWithProc(t *testing.T) {
	{
		var f func(a int) (int, error)
		rp := GetProcOfFunc1r2(f)
		ri, re := rp.Dow(func() {
			_ = Must2(9, true).OrReturnTo(rp)
			rp.Return(10, nil)
		})
		assert.Equal(t, ri, 10)
		assert.NoError(t, re)
	}

	{
		var f func(a int) (int, error)
		rp := GetProcOfFunc1r2(f)
		ri, re := rp.Dow(func() {
			_ = Must2(9, false).OrReturnTo(rp)
			rp.Return(10, nil)
		})
		assert.Equal(t, ri, 0)
		assert.Error(t, re)
	}

	{
		var f func(a int) (int, error)
		rp := GetProcOfFunc1r2(f)
		ri, re := rp.Dow(func() {
			_ = Must2(9, error(nil)).OrReturnTo(rp)
			rp.Return(10, nil)
		})
		assert.Equal(t, ri, 10)
		assert.NoError(t, re)
	}

	{
		var f func(a int) (int, error)
		rp := GetProcOfFunc1r2(f)
		ri, re := rp.Dow(func() {
			_ = Must2(9, fmt.Errorf("e")).OrReturnTo(rp)
			rp.Return(10, nil)
		})
		assert.Equal(t, ri, 0)
		assert.Error(t, re)
	}

	{
		var f func(a int) (int, bool)
		rp := GetProcOfFunc1r2(f)
		ri, re := rp.Dow(func() {
			_ = Must2(9, true).OrReturnTo(rp)
			rp.Return(10, true)
		})
		assert.Equal(t, ri, 10)
		assert.Equal(t, re, true)
	}

	{
		var f func(a int) (int, bool)
		rp := GetProcOfFunc1r2(f)
		ri, re := rp.Dow(func() {
			_ = Must2(9, false).OrReturnTo(rp)
			rp.Return(10, true)
		})
		assert.Equal(t, ri, 0)
		assert.Equal(t, re, false)
	}

	{
		var f func(a int) (int, bool)
		rp := GetProcOfFunc1r2(f)
		ri, re := rp.Dow(func() {
			_ = Must2(9, fmt.Errorf("e")).OrReturnTo(rp)
			rp.Return(10, true)
		})
		assert.Equal(t, ri, 0)
		assert.Equal(t, re, false)
	}

	{
		var f func(a int) (int, bool)
		rp := GetProcOfFunc1r2(f)
		ri, re := rp.Dow(func() {
			_ = Must2(9, error(nil)).OrReturnTo(rp)
			rp.Return(10, true)
		})
		assert.Equal(t, ri, 10)
		assert.Equal(t, re, true)
	}

}


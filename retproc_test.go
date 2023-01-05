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


// func Test_RetProc_Return(t *testing.T) {
// 	p := Proc{}
// 	assert.NotPanics(t, func() {
// 		p.Do(func() {
// 			p.ReturnOnly()
// 		})
// 	})
// 
// 	assert.NotPanics(t, func() {
// 		p.Do(func() {
// 			p.ReturnWith(nil)
// 		})
// 	})
// 
// 	assert.NotPanics(t, func() {
// 		p.Do(func() {
// 			p.IfReturnWith(true, nil)
// 		})
// 	})
// }
// 
// 
// func Test_RetProc_Panic(t *testing.T) {
// 	p := Proc{}
// 	assert.Panics(t, func() {
// 		p.Do(func() {
// 			panic("panic")
// 		})
// 	})
// }
// 
// func Test_RetProc_ReturnFunc(t *testing.T) {
// 	p := Proc{}
// 	a := 9
// 	p.Do(func() {
// 		p.ReturnWith(func() {
// 			a = 100
// 		})
// 	})
// 	assert.Equal(t, 100, a)
// 
// 	p.Do(func() {
// 		p.IfReturnWith(true, func() {
// 			a = 200
// 		})
// 		a = 250
// 	})
// 	assert.Equal(t, 200, a)
// 
// 	p.Do(func() {
// 		p.IfReturnWith(false, func() {
// 			a = 300
// 		})
// 		a = 350
// 	})
// 	assert.Equal(t, 350, a)
// }



package myhig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Myhig1(t *testing.T) {
	hg := &Myhig[Tuple1[int]]{}
	res := hg.Do(func() {
		hg.TReturn(Tw1(9))
	}).Unwrap()

	assert.Equal(t, res, 9)
}

func Test_Myhig2(t *testing.T) {
	hg := &Myhig[Tuple2[int, error]]{}
	{
		ri, re := hg.Do(func() {
			hg.VReturn(hg.V.From(9, nil))
		}).Unwrap()

		assert.Equal(t, ri, 9)
		assert.NoError(t, re)
	}
}

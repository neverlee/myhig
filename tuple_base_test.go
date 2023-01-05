package myhig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setError(t *testing.T) {
	var e error
	{
		e = nil
		setError(&e, fmt.Errorf("hello"))
		assert.Error(t, e)

		setError(&e, error(nil))
		assert.NoError(t, e)
	}
}

func Test_isFail(t *testing.T) {
	{
		assert.True(t, isFail(fmt.Errorf("e")))
		assert.False(t, isFail(error(nil)))
	}
	{
		assert.True(t, isFail(false))
		assert.False(t, isFail(true))
	}
}


package calc

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	var exp Expression
	// assert equality

	val, err := exp.Calc("5+2")
	if assert.Nil(t, err) {
		assert.Equal(t, float64(7), val)
	}

	val, err = exp.Calc("9+9")
	if assert.Nil(t, err) {
		assert.Equal(t, float64(18), val)
	}
}

func TestMultiply(t *testing.T) {
	var exp Expression

	val, err := exp.Calc("5*2")
	if assert.Nil(t, err) {
		assert.Equal(t, float64(10), val)
	}

	val, err = exp.Calc("9*3")
	if assert.Nil(t, err) {
		assert.Equal(t, float64(27), val)
	}
}

func TestDivide(t *testing.T) {
	var exp Expression

	val, err := exp.Calc("5/2")
	if assert.Nil(t, err) {
		assert.Equal(t, float64(2.5), val)
	}

	val, err = exp.Calc("1/3")
	if assert.Nil(t, err) {
		assert.Equal(t, float64(0.3333333333333333), val)
	}
}

func TestDiff(t *testing.T) {
	var exp Expression

	val, err := exp.Calc("5-9")
	if assert.Nil(t, err) {
		assert.Equal(t, float64(-4), val)
	}

	val, err = exp.Calc("0-1")
	if assert.Nil(t, err) {
		assert.Equal(t, float64(-1), val)
	}
}

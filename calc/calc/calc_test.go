package calc

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	var exp Expression
	// assert equality

	val, err := exp.Calc("8*5+2")
	if assert.Nil(t, err) {
		assert.Equal(t, val, float64(42), "they should be equal")
	}

	_, err = exp.Calc("4+2*")
	assert.NotNil(t, err)
}

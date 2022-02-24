package uniq

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	var lines []string
	var err error
	var options Options
	optionsDefault := GetDefaultOptions()

	options = optionsDefault
	lines, err = Uniq(
		[]string {
			"I love music",
			"I love music",
		}, options,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, lines,
		[]string {
			"I love music",
		}, "they should be equal")
	}

	options = optionsDefault
	options.IgnoreCase = true
	lines, err = Uniq(
		[]string {
			"I LOVE MUSIC",
			"I love music",
		}, options,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, lines,
		[]string {
			"I LOVE MUSIC",
		}, "they should be equal")
	}
}

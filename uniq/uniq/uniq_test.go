package uniq

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNoOptions(t *testing.T) {
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
}

func TestIgnoreCase(t *testing.T) {
	var lines []string
	var err error
	var options Options
	optionsDefault := GetDefaultOptions()

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

func TestFieldsChars(t *testing.T) {
	var lines []string
	var err error
	var options Options
	optionsDefault := GetDefaultOptions()

	options = optionsDefault
	options.NumFields = 1
	lines, err = Uniq(
		[]string {
			"I love music",
			"I love music",
			" ",
			"I LOVE MUSIC",
		}, options,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, lines,
		[]string {
			"I love music",
			"I love music",
			" ",
			"I LOVE MUSIC",
		}, "they should be equal")
	}

	options = optionsDefault
	options.NumFields = 1
	options.NumChars = 1
	lines, err = Uniq(
		[]string {
			"A love music",
			"B love music",
			" ",
			"C I LOVE MUSIC",
			"D I LOVE MUSIC",
		}, options,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, lines,
		[]string {
			"A love music",
			"B love music",
			" ",
			"C I LOVE MUSIC",
		}, "they should be equal")
	}
}

func TestOutputUniq(t *testing.T) {
	var lines []string
	var err error
	var options Options
	optionsDefault := GetDefaultOptions()

	options = optionsDefault
	options.OutputUnique = true
	lines, err = Uniq(
		[]string {
			"I love music",
			"I do not love music",
			"I love music",
			"I LOVE MUSIC",
		}, options,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, lines,
		[]string {
			"I do not love music",
			"I LOVE MUSIC",
		}, "they should be equal")
	}
}

func TestOutputRepeated(t *testing.T) {
	var lines []string
	var err error
	var options Options
	optionsDefault := GetDefaultOptions()

	options = optionsDefault
	options.OutputRepeated = true
	lines, err = Uniq(
		[]string {
			"I love music",
			"I do not love music",
			"I love music",
			"I LOVE MUSIC",
		}, options,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, lines,
		[]string {
			"I love music",
			"I love music",
		}, "they should be equal")
	}
}

func TestCount(t *testing.T) {
	var lines []string
	var err error
	var options Options
	optionsDefault := GetDefaultOptions()

	options = optionsDefault
	options.UseCount = true
	lines, err = Uniq(
		[]string {
			"I love music",
			"I do not love music",
			"I love music",
			"I LOVE MUSIC",
		}, options,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, lines,
		[]string {
			"2 I love music",
			"1 I do not love music",
			"2 I love music",
			"1 I LOVE MUSIC",
		}, "they should be equal")
	}

	options = optionsDefault
	options.UseCount = true
	options.IgnoreCase = true
	lines, err = Uniq(
		[]string {
			"I love music",
			"I do not love music",
			"I love music",
			"I LOVE MUSIC",
		}, options,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, lines,
		[]string {
			"3 I love music",
			"1 I do not love music",
			"3 I love music",
			"3 I LOVE MUSIC",
		}, "they should be equal")
	}
}

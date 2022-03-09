package uniq

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Test struct {
	input []string
	want []string
	options Options
}

func TestTableAll(t *testing.T) {
	optionsDefault := GetDefaultOptions()

	optIgnoreCase := optionsDefault
	optIgnoreCase.IgnoreCase = true

	optFields := optionsDefault
	optFields.NumFields = 1

	optFieldsChars := optionsDefault
	optFieldsChars.NumFields = 1
	optFieldsChars.NumChars = 1

	optOutputUnique := optionsDefault
	optOutputUnique.OutputUnique = true

	optRepeated := optionsDefault
	optRepeated.OutputRepeated = true

	optCount := optionsDefault
	optCount.UseCount = true

	tests := []Test {
		// no options
		{
			input: []string {
				"I love music",
				"I love music",
			},
			want: []string {
				"I love music",
			},
			options: GetDefaultOptions(),
		},
		// ignore case
		{
			input: []string {
				"I LOVE MUSIC",
				"I love music",
			},
			want: []string {
				"I LOVE MUSIC",
			},
			options: optIgnoreCase,
		},
		// fileds and chars
		{
			input:[]string {
				"I love music",
				"I love music",
				" ",
				"I LOVE MUSIC",
			},
			want: []string {
				"I love music",
				"I love music",
				" ",
				"I LOVE MUSIC",
			},
			options: optFields,
		},
		// fields and chars
		{
			input: []string {
				"A love music",
				"B love music",
				" ",
				"C I LOVE MUSIC",
				"D I LOVE MUSIC",
			},
			want: []string {
				"A love music",
				"B love music",
				" ",
				"C I LOVE MUSIC",
			},
			options: optFieldsChars,
		},
		// output unique
		{
			input: []string {
				"I love music",
				"I do not love music",
				"I love music",
				"I LOVE MUSIC",
			},
			want: []string {
				"I do not love music",
				"I LOVE MUSIC",
			},
			options: optOutputUnique,
		},
		// output repeated
		{
			input: []string {
				"I love music",
				"I do not love music",
				"I love music",
				"I LOVE MUSIC",
			},
			want: []string {
				"I love music",
				"I love music",
			},
			options: optRepeated,
		},
		// count
		{
			input: []string {
				"I love music",
				"I do not love music",
				"I love music",
				"I LOVE MUSIC",
			},
			want: []string {
				"2 I love music",
				"1 I do not love music",
				"2 I love music",
				"1 I LOVE MUSIC",
			},
			options: optCount,
		},
	}

	for _, test := range tests {
		res, err := Uniq(test.input, test.options)
		if assert.Nil(t, err) {
			assert.Equal(t, res, test.want)
		}
	}
}

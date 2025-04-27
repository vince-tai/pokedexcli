package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ooga  booga             ",
			expected: []string{"ooga", "booga"},
		},
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    " 2398u  4938!!!!   ,  ",
			expected: []string{"2398u", "4938!!!!", ","},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word not as expected")
			}
		}
	}
}

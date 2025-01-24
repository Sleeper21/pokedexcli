package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELlo  WOrld ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "All pokemon are GREAT!",
			expected: []string{"all", "pokemon", "are", "great!"},
		},
	}

	// loop over the cases and run the tests
	for _, c := range cases {
		actual, err := cleanInput(c.input)
		if err != nil {
			fmt.Println(err)
		}

		if len(actual) != len(c.expected) {
			t.Errorf(`---------------------------
			Input: %s
			Expecting: %s
			Actual: %v
			Failed
			`, c.input, c.expected, actual)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Fail. Found word: %s; want: %s", word, expectedWord)
			}
		}
	}
}

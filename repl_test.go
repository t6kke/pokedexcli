package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "one_word",
			expected: []string{"one_word"},
		},
		{
			input:    "  one_word_with_spaces    ",
			expected: []string{"one_word_with_spaces"},
		},
		{
			input:    "",
			expected: []string{""}, //TODO how do I actually want to handle empty input?
		},
		{
			input:    "one TWO threE 4 FiVe se7en",
			expected: []string{"one", "two", "three", "4", "five", "se7en"},
		},
		{
			input:    "    ",
			expected: []string{}, //TODO how do I actually want to handle input with just spaces in them?
		},
	}

	for _,c := range cases {
		actual := cleanInput(c.input)
		len_actual := len(actual)
		len_expected := len(c.expected)
		if len_actual != len_expected {
			t.Errorf("Result elements count '%d' does not match expected elements count '%d'", len_actual, len_expected)
			//TODO fail test?
		}
		for i := range actual {
			word := actual[i]
			expected_word := c.expected[i]
			if word != expected_word {
				t.Errorf("Result word '%s' does not match expected word '%s'", word, expected_word)
				//TODO fail test?
			}
		}
	}
}

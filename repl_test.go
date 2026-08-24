package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello  World ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HELLO   Wo rld",
			expected: []string{"hello", "wo", "rld"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for tn, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Fatalf("Fail of test number %d. Length of actual different from expected. Expected length %d. Got %d", tn, len(c.expected), len(actual))
		}

		for i, v := range actual {
			if v != c.expected[i] {
				t.Fatalf("Different value in index %d. Expected value of %s. Got %s", i, c.expected[i], v)
			}
		}
	}
}

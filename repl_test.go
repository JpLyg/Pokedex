package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "  HeLLo   WORLD  ", expected: []string{"hello", "world"}},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("len mismatch: got %d, want %d (input: %q)", len(actual), len(c.expected), c.input)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("word %d mismatch: got %q, want %q (input: %q)", i, actual[i], c.expected[i], c.input)
			}
		}
	}
}

package code

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		//{"Hello World", 5},
		//{"Hello", 5},
		//{"", 0},
		//{"   ", 0},
		//{"This is a test", 4},
		{"   fly me   to   the moon  ", 4},
	}

	for _, tc := range cases {
		result := LengthOfLastWord(tc.input)
		if result != tc.expected {
			t.Errorf("LengthOfLastWord(%q) = %d; expected %d", tc.input, result, tc.expected)
		}
	}
}

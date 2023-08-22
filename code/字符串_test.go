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

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"hello", "", -1},
		{"", "world", -1},
		{"hello world", " ", 5},
		{"sadbutsad", "buts", 3},
	}

	for _, tc := range tests {
		got := StrStr(tc.haystack, tc.needle)
		if got != tc.want {
			t.Errorf("StrStr(%q,%q) = %v; want %v", tc.haystack, tc.needle, got, tc.want)
		}
	}
}

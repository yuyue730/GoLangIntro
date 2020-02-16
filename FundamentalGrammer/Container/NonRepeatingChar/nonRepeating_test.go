package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"pwwkew", 3},
		{"abcabcbb", 3},
		{"", 0},
		{"bbb", 1},
		{"abcabcabcd", 4},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubstr(tt.s)
		if actual != tt.ans {
			t.Errorf("Got %d for input %s expect %d", actual, tt.s, tt.ans)
		}
	}
}

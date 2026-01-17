package palindromic_substrings

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1 - no palindromes except single chars",
			s:    "abc",
			want: 3,
		},
		{
			name: "example 2 - all same characters",
			s:    "aaa",
			want: 6,
		},
		{
			name: "example 3 - word palindrome",
			s:    "racecar",
			want: 10,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "two same characters",
			s:    "aa",
			want: 3,
		},
		{
			name: "two different characters",
			s:    "ab",
			want: 2,
		},
		{
			name: "palindrome at start",
			s:    "aab",
			want: 4,
		},
		{
			name: "palindrome at end",
			s:    "baa",
			want: 4,
		},
		{
			name: "palindrome in middle",
			s:    "baaab",
			want: 9,
		},
		{
			name: "multiple palindromes",
			s:    "abba",
			want: 6,
		},
		{
			name: "long repeated characters",
			s:    "aaaa",
			want: 10,
		},
		{
			name: "complex palindrome",
			s:    "abcba",
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountSubstrings(tt.s)
			if got != tt.want {
				t.Errorf("CountSubstrings(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

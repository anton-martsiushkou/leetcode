package longest_substring

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "example 3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "all unique characters",
			s:    "abcdef",
			want: 6,
		},
		{
			name: "with spaces and symbols",
			s:    "a b c a",
			want: 3,
		},
		{
			name: "long string with repeat at end",
			s:    "abcdefga",
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

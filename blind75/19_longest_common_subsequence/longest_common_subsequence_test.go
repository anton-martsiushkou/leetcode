package longest_common_subsequence

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			name:  "example 1",
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			name:  "example 2",
			text1: "abc",
			text2: "abc",
			want:  3,
		},
		{
			name:  "example 3",
			text1: "abc",
			text2: "def",
			want:  0,
		},
		{
			name:  "one character match",
			text1: "abc",
			text2: "axc",
			want:  2,
		},
		{
			name:  "single character strings",
			text1: "a",
			text2: "a",
			want:  1,
		},
		{
			name:  "single character no match",
			text1: "a",
			text2: "b",
			want:  0,
		},
		{
			name:  "longer strings",
			text1: "abcdefghij",
			text2: "ecdgi",
			want:  4,
		},
		{
			name:  "reversed strings",
			text1: "abc",
			text2: "cba",
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestCommonSubsequence(tt.text1, tt.text2)
			if got != tt.want {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

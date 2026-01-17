package word_break

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			name:     "example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			name:     "single word exact match",
			s:        "apple",
			wordDict: []string{"apple"},
			want:     true,
		},
		{
			name:     "single word no match",
			s:        "apple",
			wordDict: []string{"orange"},
			want:     false,
		},
		{
			name:     "empty string",
			s:        "",
			wordDict: []string{"a"},
			want:     true,
		},
		{
			name:     "reuse word multiple times",
			s:        "aaaaaaa",
			wordDict: []string{"aa", "aaa"},
			want:     true,
		},
		{
			name:     "complex segmentation",
			s:        "cars",
			wordDict: []string{"car", "ca", "rs"},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordBreak(tt.s, tt.wordDict)
			if got != tt.want {
				t.Errorf("WordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

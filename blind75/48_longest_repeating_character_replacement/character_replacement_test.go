package character_replacement

import (
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "example 1",
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "example 2",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "all same characters",
			s:    "AAAA",
			k:    0,
			want: 4,
		},
		{
			name: "k equals length",
			s:    "ABCD",
			k:    4,
			want: 4,
		},
		{
			name: "single character",
			s:    "A",
			k:    0,
			want: 1,
		},
		{
			name: "no replacements",
			s:    "ABCDE",
			k:    0,
			want: 1,
		},
		{
			name: "alternating characters",
			s:    "ABABAB",
			k:    2,
			want: 5,
		},
		{
			name: "large k",
			s:    "AABA",
			k:    10,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CharacterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("CharacterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

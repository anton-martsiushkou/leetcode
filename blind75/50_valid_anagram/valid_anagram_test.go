package valid_anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "example 1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "example 2",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "single character same",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "single character different",
			s:    "a",
			t:    "b",
			want: false,
		},
		{
			name: "different lengths",
			s:    "abc",
			t:    "abcd",
			want: false,
		},
		{
			name: "same characters different frequencies",
			s:    "aab",
			t:    "abb",
			want: false,
		},
		{
			name: "all same character",
			s:    "aaaa",
			t:    "aaaa",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

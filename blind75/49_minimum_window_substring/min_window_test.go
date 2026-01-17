package min_window

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "example 1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "example 2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "example 3",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "no valid window",
			s:    "abc",
			t:    "xyz",
			want: "",
		},
		{
			name: "entire string is minimum",
			s:    "abc",
			t:    "abc",
			want: "abc",
		},
		{
			name: "duplicate characters in t",
			s:    "aaabbbc",
			t:    "abc",
			want: "abbc",
		},
		{
			name: "multiple valid windows",
			s:    "abcabdebac",
			t:    "abc",
			want: "bac",
		},
		{
			name: "s shorter than t",
			s:    "ab",
			t:    "abc",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("MinWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

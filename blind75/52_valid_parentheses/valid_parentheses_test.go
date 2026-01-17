package valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "()",
			want: true,
		},
		{
			name: "example 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "example 3",
			s:    "(]",
			want: false,
		},
		{
			name: "example 4",
			s:    "([])",
			want: true,
		},
		{
			name: "nested brackets",
			s:    "{[()]}",
			want: true,
		},
		{
			name: "mismatched brackets",
			s:    "([)]",
			want: false,
		},
		{
			name: "only opening brackets",
			s:    "((((",
			want: false,
		},
		{
			name: "only closing brackets",
			s:    "))))",
			want: false,
		},
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "single opening bracket",
			s:    "(",
			want: false,
		},
		{
			name: "single closing bracket",
			s:    ")",
			want: false,
		},
		{
			name: "complex valid",
			s:    "{[()()]}(){}",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.s)
			if got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

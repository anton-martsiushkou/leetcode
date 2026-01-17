package unique_paths

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "example 1",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "example 2",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "1x1 grid",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "1x10 grid",
			m:    1,
			n:    10,
			want: 1,
		},
		{
			name: "10x1 grid",
			m:    10,
			n:    1,
			want: 1,
		},
		{
			name: "square grid 3x3",
			m:    3,
			n:    3,
			want: 6,
		},
		{
			name: "square grid 4x4",
			m:    4,
			n:    4,
			want: 20,
		},
		{
			name: "2x2 grid",
			m:    2,
			n:    2,
			want: 2,
		},
		{
			name: "large grid",
			m:    10,
			n:    10,
			want: 48620,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("UniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

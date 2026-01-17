package container_with_most_water

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "example 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "ascending order",
			height: []int{1, 2, 3, 4, 5},
			want:   6,
		},
		{
			name:   "descending order",
			height: []int{5, 4, 3, 2, 1},
			want:   6,
		},
		{
			name:   "all same height",
			height: []int{5, 5, 5, 5},
			want:   15,
		},
		{
			name:   "two tall ends",
			height: []int{10, 1, 1, 1, 10},
			want:   40,
		},
		{
			name:   "tall in middle",
			height: []int{1, 10, 1},
			want:   2,
		},
		{
			name:   "alternating heights",
			height: []int{1, 3, 2, 4, 3, 5},
			want:   12,
		},
		{
			name:   "with zeros",
			height: []int{0, 2, 0, 3, 0, 4},
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxArea(tt.height)
			if got != tt.want {
				t.Errorf("MaxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

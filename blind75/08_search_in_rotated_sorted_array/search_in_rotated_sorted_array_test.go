package search_in_rotated_sorted_array

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "example 2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "example 3",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			name:   "single element found",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "target at beginning",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 4,
			want:   0,
		},
		{
			name:   "target at end",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 2,
			want:   6,
		},
		{
			name:   "no rotation",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "rotated once",
			nums:   []int{5, 1, 2, 3, 4},
			target: 1,
			want:   1,
		},
		{
			name:   "two elements rotated",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
		{
			name:   "two elements not found",
			nums:   []int{1, 3},
			target: 2,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

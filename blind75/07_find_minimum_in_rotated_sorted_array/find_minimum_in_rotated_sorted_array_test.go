package find_minimum_in_rotated_sorted_array

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "example 3",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "two elements rotated",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "two elements not rotated",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "rotated once",
			nums: []int{2, 3, 4, 5, 1},
			want: 1,
		},
		{
			name: "not rotated",
			nums: []int{1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "negative numbers",
			nums: []int{3, 4, 5, -2, -1, 0, 1, 2},
			want: -2,
		},
		{
			name: "large rotation",
			nums: []int{5, 1, 2, 3, 4},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMin(tt.nums)
			if got != tt.want {
				t.Errorf("FindMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

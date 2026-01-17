package maximum_subarray

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "example 2",
			nums: []int{1},
			want: 1,
		},
		{
			name: "example 3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "all negative",
			nums: []int{-3, -2, -5, -1, -4},
			want: -1,
		},
		{
			name: "all positive",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "single negative",
			nums: []int{-1},
			want: -1,
		},
		{
			name: "alternating signs",
			nums: []int{-2, 3, -1, 4, -3, 2},
			want: 6,
		},
		{
			name: "zeros and negatives",
			nums: []int{-2, 0, -1, 0, -3},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

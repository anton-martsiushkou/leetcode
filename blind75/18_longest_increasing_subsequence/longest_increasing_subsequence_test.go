package longest_increasing_subsequence

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "example 2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "example 3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "strictly decreasing",
			nums: []int{5, 4, 3, 2, 1},
			want: 1,
		},
		{
			name: "strictly increasing",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "alternating",
			nums: []int{1, 3, 2, 4, 3, 5},
			want: 4,
		},
		{
			name: "negative numbers",
			nums: []int{-10, -5, 0, 3, 1, 2},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLIS(tt.nums)
			if got != tt.want {
				t.Errorf("LengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

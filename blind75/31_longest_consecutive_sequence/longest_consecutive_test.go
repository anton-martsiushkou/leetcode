package longest_consecutive

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "example 2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "all duplicates",
			nums: []int{1, 1, 1, 1},
			want: 1,
		},
		{
			name: "no consecutive",
			nums: []int{1, 3, 5, 7, 9},
			want: 1,
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, 0, 1, 2},
			want: 5,
		},
		{
			name: "consecutive at end",
			nums: []int{10, 5, 1, 2, 3, 4},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestConsecutive(tt.nums)
			if got != tt.want {
				t.Errorf("LongestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestConsecutiveSorting(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "example 2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "all duplicates",
			nums: []int{1, 1, 1, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestConsecutiveSorting(tt.nums)
			if got != tt.want {
				t.Errorf("LongestConsecutiveSorting() = %v, want %v", got, tt.want)
			}
		})
	}
}

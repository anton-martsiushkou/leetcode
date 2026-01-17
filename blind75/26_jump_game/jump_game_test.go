package jump_game

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example 1",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "example 2",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "single element",
			nums: []int{0},
			want: true,
		},
		{
			name: "two elements can reach",
			nums: []int{1, 0},
			want: true,
		},
		{
			name: "two elements cannot reach",
			nums: []int{0, 1},
			want: false,
		},
		{
			name: "all zeros except first",
			nums: []int{5, 0, 0, 0, 0, 0},
			want: true,
		},
		{
			name: "large jump at start",
			nums: []int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: true,
		},
		{
			name: "can barely reach",
			nums: []int{1, 1, 1, 1, 1},
			want: true,
		},
		{
			name: "stuck in middle",
			nums: []int{1, 1, 0, 1},
			want: false,
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1},
			want: true,
		},
		{
			name: "decreasing jumps",
			nums: []int{5, 4, 3, 2, 1, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanJump(tt.nums)
			if got != tt.want {
				t.Errorf("CanJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

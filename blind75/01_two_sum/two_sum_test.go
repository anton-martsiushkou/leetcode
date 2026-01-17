package two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "zero in array",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if len(got) != len(tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
				return
			}
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

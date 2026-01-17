package missing_number

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "example 2",
			nums: []int{0, 1},
			want: 2,
		},
		{
			name: "example 3",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
		{
			name: "missing 0",
			nums: []int{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			name: "missing last",
			nums: []int{0, 1, 2, 3, 4},
			want: 5,
		},
		{
			name: "single element missing 0",
			nums: []int{1},
			want: 0,
		},
		{
			name: "single element missing 1",
			nums: []int{0},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MissingNumber(tt.nums)
			if got != tt.want {
				t.Errorf("MissingNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

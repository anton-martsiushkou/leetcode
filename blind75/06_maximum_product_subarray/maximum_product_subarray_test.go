package maximum_product_subarray

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "example 2",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "example 3",
			nums: []int{-2, 3, -4},
			want: 24,
		},
		{
			name: "single element positive",
			nums: []int{5},
			want: 5,
		},
		{
			name: "single element negative",
			nums: []int{-3},
			want: -3,
		},
		{
			name: "all positive",
			nums: []int{2, 3, 4},
			want: 24,
		},
		{
			name: "all negative even count",
			nums: []int{-2, -3, -4, -5},
			want: 120,
		},
		{
			name: "all negative odd count",
			nums: []int{-2, -3, -4},
			want: 12,
		},
		{
			name: "with zeros",
			nums: []int{0, 2, 3},
			want: 6,
		},
		{
			name: "negative at beginning",
			nums: []int{-2, 3, 4},
			want: 12,
		},
		{
			name: "multiple zeros",
			nums: []int{0, 2, 0, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProduct(tt.nums)
			if got != tt.want {
				t.Errorf("MaxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

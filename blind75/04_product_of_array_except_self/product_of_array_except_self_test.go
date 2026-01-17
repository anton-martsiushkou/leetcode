package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "example 2",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "two elements",
			nums: []int{2, 3},
			want: []int{3, 2},
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "with negative numbers",
			nums: []int{-2, -3, 4, -5},
			want: []int{-60, -40, 30, -24},
		},
		{
			name: "zero at beginning",
			nums: []int{0, 2, 3, 4},
			want: []int{24, 0, 0, 0},
		},
		{
			name: "zero at end",
			nums: []int{2, 3, 4, 0},
			want: []int{0, 0, 0, 24},
		},
		{
			name: "multiple zeros",
			nums: []int{0, 0, 1, 2},
			want: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

package spiral_matrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name: "example 1",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "example 2",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "single element",
			matrix: [][]int{
				{1},
			},
			want: []int{1},
		},
		{
			name: "single row",
			matrix: [][]int{
				{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "single column",
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "2x2 matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			want: []int{1, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SpiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

package three_sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "all negative",
			nums: []int{-1, -2, -3},
			want: [][]int{},
		},
		{
			name: "all positive",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "multiple triplets",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name: "with duplicates",
			nums: []int{-4, -1, -1, 0, 1, 2},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "large numbers",
			nums: []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			want: [][]int{{-5, 0, 5}, {-5, 1, 4}, {-5, 2, 3}, {-4, -1, 5}, {-4, 0, 4}, {-4, 1, 3}, {-3, -2, 5}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.nums)
			// Sort both slices for comparison
			sortTriplets(got)
			sortTriplets(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to sort triplets for comparison
func sortTriplets(triplets [][]int) {
	for _, triplet := range triplets {
		sort.Ints(triplet)
	}
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}

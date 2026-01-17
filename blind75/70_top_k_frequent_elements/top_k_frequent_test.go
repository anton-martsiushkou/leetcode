package top_k_frequent

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "example 2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "all same frequency",
			nums: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{1, 2, 3}, // any 3 elements
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, -1, 2, 2, 3},
			k:    2,
			want: []int{-1, 2},
		},
		{
			name: "k equals unique elements",
			nums: []int{4, 4, 4, 5, 5, 6},
			k:    3,
			want: []int{4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.nums, tt.k)
			if len(got) != len(tt.want) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
				return
			}
			// Sort both slices for comparison since order doesn't matter
			sort.Ints(got)
			sort.Ints(tt.want)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

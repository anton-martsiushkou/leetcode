package combination_sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "example 3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			name:       "single candidate exact match",
			candidates: []int{5},
			target:     5,
			want:       [][]int{{5}},
		},
		{
			name:       "single candidate multiple uses",
			candidates: []int{3},
			target:     9,
			want:       [][]int{{3, 3, 3}},
		},
		{
			name:       "multiple combinations",
			candidates: []int{2, 3, 4},
			target:     6,
			want:       [][]int{{2, 2, 2}, {2, 4}, {3, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinationSum(tt.candidates, tt.target)

			// Sort both got and want for comparison
			sortCombinations(got)
			sortCombinations(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to sort combinations for comparison
func sortCombinations(combinations [][]int) {
	for _, combo := range combinations {
		sort.Ints(combo)
	}
	sort.Slice(combinations, func(i, j int) bool {
		for k := 0; k < len(combinations[i]) && k < len(combinations[j]); k++ {
			if combinations[i][k] != combinations[j][k] {
				return combinations[i][k] < combinations[j][k]
			}
		}
		return len(combinations[i]) < len(combinations[j])
	})
}

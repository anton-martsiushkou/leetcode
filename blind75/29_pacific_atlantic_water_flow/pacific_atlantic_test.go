package pacific_atlantic

import (
	"reflect"
	"sort"
	"testing"
)

func sortResult(result [][]int) [][]int {
	sort.Slice(result, func(i, j int) bool {
		if result[i][0] == result[j][0] {
			return result[i][1] < result[j][1]
		}
		return result[i][0] < result[j][0]
	})
	return result
}

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		{
			name: "example 1",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			name:    "example 2 - single cell",
			heights: [][]int{{1}},
			want:    [][]int{{0, 0}},
		},
		{
			name:    "2x2 grid",
			heights: [][]int{{1, 2}, {3, 4}},
			want:    [][]int{{0, 1}, {1, 0}, {1, 1}},
		},
		{
			name:    "all same height",
			heights: [][]int{{1, 1}, {1, 1}},
			want:    [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PacificAtlantic(tt.heights)
			got = sortResult(got)
			want := sortResult(tt.want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("PacificAtlantic() = %v, want %v", got, want)
			}
		})
	}
}

func TestPacificAtlanticBFS(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		{
			name: "example 1",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			name:    "example 2 - single cell",
			heights: [][]int{{1}},
			want:    [][]int{{0, 0}},
		},
		{
			name:    "all same height",
			heights: [][]int{{1, 1}, {1, 1}},
			want:    [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PacificAtlanticBFS(tt.heights)
			got = sortResult(got)
			want := sortResult(tt.want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("PacificAtlanticBFS() = %v, want %v", got, want)
			}
		})
	}
}

package merge_intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "example 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "example 2 - touching intervals",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "example 3 - unsorted intervals",
			intervals: [][]int{{1, 4}, {0, 4}},
			want:      [][]int{{0, 4}},
		},
		{
			name:      "example 4 - nested interval",
			intervals: [][]int{{1, 4}, {2, 3}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "no overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:      [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "all intervals merge into one",
			intervals: [][]int{{1, 4}, {2, 5}, {3, 6}},
			want:      [][]int{{1, 6}},
		},
		{
			name:      "multiple merges",
			intervals: [][]int{{1, 3}, {2, 4}, {5, 7}, {6, 8}},
			want:      [][]int{{1, 4}, {5, 8}},
		},
		{
			name:      "completely nested intervals",
			intervals: [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}},
			want:      [][]int{{1, 10}},
		},
		{
			name:      "unsorted with multiple overlaps",
			intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			want:      [][]int{{1, 10}},
		},
		{
			name:      "same start different end",
			intervals: [][]int{{1, 4}, {1, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "empty intervals",
			intervals: [][]int{},
			want:      [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

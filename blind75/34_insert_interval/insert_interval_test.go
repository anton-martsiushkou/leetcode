package insert_interval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			name:        "example 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "example 3 - empty intervals",
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			want:        [][]int{{5, 7}},
		},
		{
			name:        "example 4 - new interval inside existing",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			want:        [][]int{{1, 5}},
		},
		{
			name:        "example 5 - new interval extends existing",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 7},
			want:        [][]int{{1, 7}},
		},
		{
			name:        "new interval before all",
			intervals:   [][]int{{3, 5}, {6, 9}},
			newInterval: []int{1, 2},
			want:        [][]int{{1, 2}, {3, 5}, {6, 9}},
		},
		{
			name:        "new interval after all",
			intervals:   [][]int{{1, 2}, {3, 5}},
			newInterval: []int{6, 9},
			want:        [][]int{{1, 2}, {3, 5}, {6, 9}},
		},
		{
			name:        "merge all intervals",
			intervals:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			newInterval: []int{0, 7},
			want:        [][]int{{0, 7}},
		},
		{
			name:        "adjacent intervals - no overlap",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{6, 8},
			want:        [][]int{{1, 5}, {6, 8}},
		},
		{
			name:        "adjacent intervals - touching",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{5, 8},
			want:        [][]int{{1, 8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

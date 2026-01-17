package non_overlapping_intervals

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "example 1",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:      1,
		},
		{
			name:      "example 2 - all identical",
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			want:      2,
		},
		{
			name:      "example 3 - no overlaps",
			intervals: [][]int{{1, 2}, {2, 3}},
			want:      0,
		},
		{
			name:      "example 4",
			intervals: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
			want:      2,
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 2}},
			want:      0,
		},
		{
			name:      "two non-overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}},
			want:      0,
		},
		{
			name:      "two overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 4}},
			want:      1,
		},
		{
			name:      "all intervals overlap - nested",
			intervals: [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
			want:      3,
		},
		{
			name:      "multiple groups of overlaps",
			intervals: [][]int{{1, 2}, {1, 3}, {2, 3}, {5, 6}, {5, 7}, {6, 7}},
			want:      3,
		},
		{
			name:      "touching intervals - not overlapping",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			want:      0,
		},
		{
			name:      "reverse sorted intervals",
			intervals: [][]int{{5, 6}, {4, 5}, {3, 4}, {2, 3}, {1, 2}},
			want:      0,
		},
		{
			name:      "complex overlapping case",
			intervals: [][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}},
			want:      2,
		},
		{
			name:      "one interval contains all others",
			intervals: [][]int{{1, 100}, {2, 3}, {4, 5}, {6, 7}},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EraseOverlapIntervals(tt.intervals)
			if got != tt.want {
				t.Errorf("EraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

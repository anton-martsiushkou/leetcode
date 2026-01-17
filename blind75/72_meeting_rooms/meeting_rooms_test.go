package meeting_rooms

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      bool
	}{
		{
			name:      "example 1 - overlapping meetings",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			want:      false,
		},
		{
			name:      "example 2 - non-overlapping meetings",
			intervals: [][]int{{7, 10}, {2, 4}},
			want:      true,
		},
		{
			name:      "empty intervals",
			intervals: [][]int{},
			want:      true,
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 5}},
			want:      true,
		},
		{
			name:      "adjacent meetings - no overlap",
			intervals: [][]int{{1, 5}, {5, 10}},
			want:      true,
		},
		{
			name:      "adjacent meetings with gap",
			intervals: [][]int{{1, 5}, {6, 10}, {11, 15}},
			want:      true,
		},
		{
			name:      "multiple overlaps",
			intervals: [][]int{{1, 10}, {2, 6}, {3, 5}, {7, 9}},
			want:      false,
		},
		{
			name:      "meetings in reverse order",
			intervals: [][]int{{15, 20}, {10, 15}, {5, 10}},
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanAttendMeetings(tt.intervals)
			if got != tt.want {
				t.Errorf("CanAttendMeetings() = %v, want %v", got, tt.want)
			}
		})
	}
}

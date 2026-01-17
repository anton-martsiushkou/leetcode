package meeting_rooms_ii

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "example 1",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			want:      2,
		},
		{
			name:      "example 2",
			intervals: [][]int{{7, 10}, {2, 4}},
			want:      1,
		},
		{
			name:      "empty intervals",
			intervals: [][]int{},
			want:      0,
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 5}},
			want:      1,
		},
		{
			name:      "all overlapping",
			intervals: [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
			want:      4,
		},
		{
			name:      "no overlapping",
			intervals: [][]int{{1, 5}, {6, 10}, {11, 15}},
			want:      1,
		},
		{
			name:      "adjacent meetings",
			intervals: [][]int{{1, 5}, {5, 10}, {10, 15}},
			want:      1,
		},
		{
			name:      "complex scenario",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}, {25, 35}, {30, 40}},
			want:      3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinMeetingRooms(tt.intervals)
			if got != tt.want {
				t.Errorf("MinMeetingRooms() = %v, want %v", got, tt.want)
			}

			// Also test chronological approach
			got2 := MinMeetingRoomsChronological(tt.intervals)
			if got2 != tt.want {
				t.Errorf("MinMeetingRoomsChronological() = %v, want %v", got2, tt.want)
			}
		})
	}
}

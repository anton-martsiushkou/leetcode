package meeting_rooms

import "sort"

// CanAttendMeetings determines if a person can attend all meetings.
// Time: O(n log n), Space: O(1)
func CanAttendMeetings(intervals [][]int) bool {
	if len(intervals) <= 1 {
		return true
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Check for overlaps in consecutive intervals
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false // Overlap found
		}
	}

	return true
}

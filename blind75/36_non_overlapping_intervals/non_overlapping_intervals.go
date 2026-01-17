package non_overlapping_intervals

import "sort"

// EraseOverlapIntervals returns the minimum number of intervals to remove
// to make the rest non-overlapping.
// Time: O(n log n), Space: O(1)
func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort intervals by end time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1 // We can always keep the first interval
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		// If current interval doesn't overlap with the last kept interval
		if intervals[i][0] >= end {
			count++
			end = intervals[i][1]
		}
		// Else: current interval overlaps, so we skip it (remove it)
	}

	// Number of removals = total intervals - intervals we kept
	return len(intervals) - count
}

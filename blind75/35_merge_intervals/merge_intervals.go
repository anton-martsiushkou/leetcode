package merge_intervals

import "sort"

// Merge merges all overlapping intervals.
// Time: O(n log n), Space: O(n)
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		lastIdx := len(result) - 1
		current := intervals[i]
		last := result[lastIdx]

		// Check if current interval overlaps with last merged interval
		if current[0] <= last[1] {
			// Merge: extend the end time of last interval
			result[lastIdx][1] = max(last[1], current[1])
		} else {
			// No overlap: add current interval
			result = append(result, current)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package insert_interval

// Insert inserts a new interval into a sorted list of non-overlapping intervals.
// Merges overlapping intervals if necessary.
// Time: O(n), Space: O(n)
func Insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	i := 0
	n := len(intervals)

	// Phase 1: Add all intervals that come before newInterval
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Phase 2: Merge all overlapping intervals with newInterval
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	// Phase 3: Add all intervals that come after newInterval
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

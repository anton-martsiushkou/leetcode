package combination_sum

import "sort"

// CombinationSum finds all unique combinations that sum to target.
// Uses backtracking with index tracking to avoid duplicates.
func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	// Sort for early termination optimization
	sort.Ints(candidates)

	backtrack(candidates, target, 0, path, &result)
	return result
}

func backtrack(candidates []int, target int, start int, path []int, result *[][]int) {
	// Base case: found valid combination
	if target == 0 {
		// Make a copy of path to add to result
		combination := make([]int, len(path))
		copy(combination, path)
		*result = append(*result, combination)
		return
	}

	// Base case: target is negative
	if target < 0 {
		return
	}

	// Try each candidate starting from 'start'
	for i := start; i < len(candidates); i++ {
		candidate := candidates[i]

		// Early termination: if candidate > target, skip (array is sorted)
		if candidate > target {
			break
		}

		// Add candidate to current path
		path = append(path, candidate)

		// Recurse with same start index (can reuse same number)
		backtrack(candidates, target-candidate, i, path, result)

		// Backtrack: remove last candidate
		path = path[:len(path)-1]
	}
}

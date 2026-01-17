package contains_duplicate

// ContainsDuplicate returns true if any value appears at least twice.
// Uses hash set for O(n) time complexity.
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

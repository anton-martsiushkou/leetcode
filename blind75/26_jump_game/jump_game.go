package jump_game

// CanJump determines if it's possible to reach the last index.
// Uses a greedy algorithm with O(n) time and O(1) space complexity.
func CanJump(nums []int) bool {
	maxReach := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		// If current position is beyond max reachable, can't proceed
		if i > maxReach {
			return false
		}

		// Update the farthest position we can reach
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}

		// If we can reach the last index, return true
		if maxReach >= n-1 {
			return true
		}
	}

	return true
}

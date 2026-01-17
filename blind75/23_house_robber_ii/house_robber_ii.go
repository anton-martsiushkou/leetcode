package house_robber_ii

// Rob returns the maximum amount of money that can be robbed from circular houses.
// Uses dynamic programming with two passes: O(n) time and O(1) space complexity.
func Rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// Two scenarios: rob houses [0...n-2] or [1...n-1]
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

// robRange applies the House Robber I algorithm to a range of houses
func robRange(nums []int, start, end int) int {
	prev2 := 0
	prev1 := 0

	for i := start; i <= end; i++ {
		current := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package house_robber

// Rob returns the maximum amount of money that can be robbed.
// Uses dynamic programming with O(n) time and O(1) space complexity.
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// prev2 = max money robbed up to house i-2
	// prev1 = max money robbed up to house i-1
	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
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

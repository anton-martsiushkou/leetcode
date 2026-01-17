package maximum_subarray

// MaxSubArray finds the contiguous subarray with the largest sum.
// Uses Kadane's Algorithm with O(n) time and O(1) space.
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// Either extend existing subarray or start new one
		currentSum = max(nums[i], currentSum+nums[i])
		// Update global maximum
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

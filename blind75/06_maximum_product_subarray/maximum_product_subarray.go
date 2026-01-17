package maximum_product_subarray

// MaxProduct finds the contiguous subarray with the largest product.
// Uses modified Kadane's algorithm tracking both max and min products.
func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar := nums[0]
	maxEndingHere := nums[0]
	minEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// Calculate candidates: current number, max*num, min*num
		// We need temp variable because maxEndingHere is used in minEndingHere calculation
		tempMax := max(num, max(num*maxEndingHere, num*minEndingHere))
		minEndingHere = min(num, min(num*maxEndingHere, num*minEndingHere))
		maxEndingHere = tempMax

		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

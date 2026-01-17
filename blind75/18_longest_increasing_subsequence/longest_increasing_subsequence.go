package longest_increasing_subsequence

// LengthOfLIS returns the length of the longest increasing subsequence.
// Uses dynamic programming with binary search for O(n log n) time complexity.
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[i] is the smallest tail of all increasing subsequences of length i+1
	tails := make([]int, 0, len(nums))

	for _, num := range nums {
		// Binary search for the position to insert/replace
		left, right := 0, len(tails)

		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// If num is larger than all elements, append it
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// Replace the element at position left
			tails[left] = num
		}
	}

	return len(tails)
}

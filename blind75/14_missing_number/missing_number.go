package missing_number

// MissingNumber finds the missing number in the range [0, n].
// Uses XOR bit manipulation for O(1) space complexity.
func MissingNumber(nums []int) int {
	result := len(nums)

	for i, num := range nums {
		result ^= i ^ num
	}

	return result
}

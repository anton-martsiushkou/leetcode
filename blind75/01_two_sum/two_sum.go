package two_sum

// TwoSum returns indices of two numbers that add up to target.
// Uses a hash map for O(n) time complexity.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}

package product_of_array_except_self

// ProductExceptSelf returns an array where each element is the product of all elements except itself.
// Uses prefix-suffix product approach with O(n) time and O(1) extra space.
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// First pass: fill result with prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// Second pass: multiply by suffix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

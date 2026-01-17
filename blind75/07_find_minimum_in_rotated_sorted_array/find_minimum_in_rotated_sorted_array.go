package find_minimum_in_rotated_sorted_array

// FindMin finds the minimum element in a rotated sorted array.
// Uses modified binary search with O(log n) time complexity.
func FindMin(nums []int) int {
	left, right := 0, len(nums)-1

	// Binary search for the minimum element
	for left < right {
		mid := left + (right-left)/2

		// If mid element is greater than right element,
		// minimum is in the right half
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// Otherwise, minimum is in the left half (including mid)
			right = mid
		}
	}

	// Left and right converge to the minimum element
	return nums[left]
}

package three_sum

import "sort"

// ThreeSum finds all unique triplets that sum to zero.
// Uses sort + two pointers approach with O(n²) time complexity.
func ThreeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)

	// Sort the array to enable two-pointer approach and duplicate handling
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// Skip duplicates for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Two-pointer approach for the remaining elements
		left, right := i+1, n-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

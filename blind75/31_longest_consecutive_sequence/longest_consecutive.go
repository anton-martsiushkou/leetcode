package longest_consecutive

// LongestConsecutive finds the longest consecutive sequence using hash set
// Time: O(n), Space: O(n)
func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Build hash set
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0

	// Check each number
	for num := range numSet {
		// Only start counting if this is the beginning of a sequence
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			// Count consecutive numbers
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

// LongestConsecutiveSorting finds the longest consecutive sequence using sorting
// Time: O(n log n), Space: O(1) or O(n) depending on sort implementation
func LongestConsecutiveSorting(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Sort is not in-place in Go, but we'll sort the input for simplicity
	// In production, you might want to make a copy first
	sorted := make([]int, len(nums))
	copy(sorted, nums)

	// Simple bubble sort for demonstration (in real code, use sort.Ints)
	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	maxLength := 1
	currentLength := 1

	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			// Skip duplicates
			continue
		} else if sorted[i] == sorted[i-1]+1 {
			// Consecutive
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
			}
		} else {
			// Break in sequence
			currentLength = 1
		}
	}

	return maxLength
}

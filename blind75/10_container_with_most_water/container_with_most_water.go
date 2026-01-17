package container_with_most_water

// MaxArea finds the maximum area of water that can be contained.
// Uses two-pointer approach with O(n) time complexity.
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Calculate area with current two lines
		width := right - left
		h := min(height[left], height[right])
		area := width * h

		// Update maximum area
		if area > maxArea {
			maxArea = area
		}

		// Move the pointer with shorter height
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

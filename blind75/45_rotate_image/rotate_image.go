package rotate_image

// Rotate rotates the image by 90 degrees clockwise in-place.
// Uses transpose + reverse approach for O(1) space complexity.
func Rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	n := len(matrix)

	// Transpose the matrix
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}

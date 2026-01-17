package spiral_matrix

// SpiralOrder returns all elements of the matrix in spiral order.
// Uses boundary tracking for O(1) space complexity.
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	top, bottom, left, right := 0, m-1, 0, n-1

	for top <= bottom && left <= right {
		// Move right along top row
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		// Move down along right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Move left along bottom row (if there's a row remaining)
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}

		// Move up along left column (if there's a column remaining)
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

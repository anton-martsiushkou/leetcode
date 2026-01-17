package set_matrix_zeroes

// SetZeroes sets entire row and column to 0 if an element is 0.
// Uses first row and column as markers for O(1) space complexity.
func SetZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	firstRowHasZero := false
	firstColHasZero := false

	// Check if first row has any zeros
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// Check if first column has any zeros
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// Use first row and column as markers
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set zeros based on markers
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Handle first row
	if firstRowHasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Handle first column
	if firstColHasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

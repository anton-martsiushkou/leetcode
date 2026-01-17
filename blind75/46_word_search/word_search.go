package word_search

// Exist returns true if word exists in the grid.
// Uses backtracking with DFS to explore all possible paths.
func Exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		// Found the complete word
		if index == len(word) {
			return true
		}

		// Out of bounds or character mismatch or already visited
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[index] {
			return false
		}

		// Mark as visited by temporarily modifying the cell
		temp := board[i][j]
		board[i][j] = '#'

		// Explore all 4 directions
		found := dfs(i+1, j, index+1) ||
			dfs(i-1, j, index+1) ||
			dfs(i, j+1, index+1) ||
			dfs(i, j-1, index+1)

		// Backtrack: restore the cell
		board[i][j] = temp

		return found
	}

	// Try each cell as a starting point
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

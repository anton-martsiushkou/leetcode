package number_of_islands

// NumIslands counts islands using DFS
// Time: O(m × n), Space: O(m × n)
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		// Check bounds and if it's land
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
			return
		}

		// Mark as visited
		grid[r][c] = '0'

		// Explore 4 directions
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// Scan the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}

// NumIslandsBFS counts islands using BFS
// Time: O(m × n), Space: O(min(m, n))
func NumIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	bfs := func(startR, startC int) {
		queue := [][]int{{startR, startC}}
		grid[startR][startC] = '0'

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			r, c := cell[0], cell[1]

			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '1' {
					grid[nr][nc] = '0'
					queue = append(queue, []int{nr, nc})
				}
			}
		}
	}

	// Scan the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				bfs(r, c)
			}
		}
	}

	return count
}

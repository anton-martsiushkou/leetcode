package pacific_atlantic

// PacificAtlantic finds cells that can flow to both oceans using DFS
// Time: O(m × n), Space: O(m × n)
func PacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)

	for i := range pacific {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	var dfs func(r, c, prevHeight int, visited [][]bool)
	dfs = func(r, c, prevHeight int, visited [][]bool) {
		// Check bounds and visited
		if r < 0 || r >= rows || c < 0 || c >= cols || visited[r][c] {
			return
		}
		// Water can't flow uphill (reverse flow)
		if heights[r][c] < prevHeight {
			return
		}

		visited[r][c] = true

		// Explore 4 directions
		dfs(r+1, c, heights[r][c], visited)
		dfs(r-1, c, heights[r][c], visited)
		dfs(r, c+1, heights[r][c], visited)
		dfs(r, c-1, heights[r][c], visited)
	}

	// DFS from Pacific borders (top and left)
	for c := 0; c < cols; c++ {
		dfs(0, c, heights[0][c], pacific)
		dfs(rows-1, c, heights[rows-1][c], atlantic)
	}
	for r := 0; r < rows; r++ {
		dfs(r, 0, heights[r][0], pacific)
		dfs(r, cols-1, heights[r][cols-1], atlantic)
	}

	// Find intersection
	result := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

// PacificAtlanticBFS finds cells that can flow to both oceans using BFS
// Time: O(m × n), Space: O(m × n)
func PacificAtlanticBFS(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)

	for i := range pacific {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	bfs := func(queue [][]int, visited [][]bool) {
		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			r, c := cell[0], cell[1]

			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
					!visited[nr][nc] && heights[nr][nc] >= heights[r][c] {
					visited[nr][nc] = true
					queue = append(queue, []int{nr, nc})
				}
			}
		}
	}

	// Initialize queues with border cells
	pacificQueue := [][]int{}
	atlanticQueue := [][]int{}

	for c := 0; c < cols; c++ {
		pacific[0][c] = true
		pacificQueue = append(pacificQueue, []int{0, c})
		atlantic[rows-1][c] = true
		atlanticQueue = append(atlanticQueue, []int{rows - 1, c})
	}
	for r := 0; r < rows; r++ {
		pacific[r][0] = true
		pacificQueue = append(pacificQueue, []int{r, 0})
		atlantic[r][cols-1] = true
		atlanticQueue = append(atlanticQueue, []int{r, cols - 1})
	}

	bfs(pacificQueue, pacific)
	bfs(atlanticQueue, atlantic)

	// Find intersection
	result := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

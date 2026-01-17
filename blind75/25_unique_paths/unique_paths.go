package unique_paths

// UniquePaths returns the number of unique paths from top-left to bottom-right.
// Uses dynamic programming with O(m*n) time and O(n) space complexity.
func UniquePaths(m int, n int) int {
	// Create a 1D array to store the number of paths to each cell in current row
	dp := make([]int, n)

	// Initialize first row - only one way to reach any cell in first row
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	// Process each row starting from row 1
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// dp[j] currently holds paths from above (previous row)
			// dp[j-1] holds paths from left (current row)
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}

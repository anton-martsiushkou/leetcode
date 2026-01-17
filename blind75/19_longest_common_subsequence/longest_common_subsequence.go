package longest_common_subsequence

// LongestCommonSubsequence returns the length of the longest common subsequence.
// Uses dynamic programming with O(m*n) time and space complexity.
func LongestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// dp[i][j] = LCS length of text1[0...i-1] and text2[0...j-1]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Build DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// Characters match, extend previous LCS
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// Take maximum of excluding either character
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

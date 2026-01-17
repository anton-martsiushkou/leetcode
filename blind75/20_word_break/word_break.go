package word_break

// WordBreak determines if string s can be segmented into dictionary words.
// Uses dynamic programming with O(n^2) time complexity.
func WordBreak(s string, wordDict []string) bool {
	// Convert wordDict to set for O(1) lookup
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	// dp[i] = true if s[0:i] can be segmented
	dp := make([]bool, n+1)
	dp[0] = true // Empty string can be segmented

	// Build DP table
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// If s[0:j] can be segmented and s[j:i] is in dictionary
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

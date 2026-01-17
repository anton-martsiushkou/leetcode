package longest_palindromic_substring

// LongestPalindrome returns the longest palindromic substring in s.
// Uses expand around center approach for O(n²) time complexity.
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 0

	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes (center at i)
		len1 := expandAroundCenter(s, i, i)

		// Check for even-length palindromes (center between i and i+1)
		len2 := expandAroundCenter(s, i, i+1)

		// Get the maximum length from both cases
		length := max(len1, len2)

		// Update if we found a longer palindrome
		if length > maxLen {
			maxLen = length
			start = i - (length-1)/2
		}
	}

	return s[start : start+maxLen]
}

// expandAroundCenter expands around the center and returns the length of the palindrome.
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

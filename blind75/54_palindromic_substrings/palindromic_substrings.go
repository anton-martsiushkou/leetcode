package palindromic_substrings

// CountSubstrings returns the number of palindromic substrings in s.
// Uses expand around center approach for O(n²) time complexity.
func CountSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0

	for i := 0; i < len(s); i++ {
		// Count odd-length palindromes (center at i)
		count += expandAroundCenter(s, i, i)

		// Count even-length palindromes (center between i and i+1)
		count += expandAroundCenter(s, i, i+1)
	}

	return count
}

// expandAroundCenter expands around the center and counts palindromes.
func expandAroundCenter(s string, left, right int) int {
	count := 0

	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}

	return count
}

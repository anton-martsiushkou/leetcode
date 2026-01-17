package decode_ways

// NumDecodings returns the number of ways to decode the string.
// Uses dynamic programming with O(n) time and O(1) space complexity.
func NumDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	if n == 1 {
		return 1
	}

	// prev2 = number of ways to decode up to i-2
	// prev1 = number of ways to decode up to i-1
	prev2 := 1 // empty string has one way to decode
	prev1 := 1 // first character is valid (already checked)

	for i := 1; i < n; i++ {
		current := 0

		// Check if single digit is valid (1-9)
		if s[i] >= '1' && s[i] <= '9' {
			current += prev1
		}

		// Check if two digits form a valid number (10-26)
		twoDigit := int(s[i-1]-'0')*10 + int(s[i]-'0')
		if twoDigit >= 10 && twoDigit <= 26 {
			current += prev2
		}

		prev2 = prev1
		prev1 = current
	}

	return prev1
}

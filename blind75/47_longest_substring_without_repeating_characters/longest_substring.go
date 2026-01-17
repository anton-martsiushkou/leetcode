package longest_substring

// LengthOfLongestSubstring finds the length of the longest substring without repeating characters.
// Uses sliding window with hash map for O(n) time complexity.
func LengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	maxLength := 0
	left := 0

	for right, char := range s {
		if prevIndex, found := charIndex[char]; found && prevIndex >= left {
			left = prevIndex + 1
		}
		charIndex[char] = right
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

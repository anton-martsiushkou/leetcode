package character_replacement

// CharacterReplacement finds the length of longest substring with same letter after k replacements.
// Uses sliding window with frequency map for O(n) time complexity.
func CharacterReplacement(s string, k int) int {
	freq := make(map[rune]int)
	maxCount := 0
	maxLength := 0
	left := 0

	for right, char := range s {
		freq[char]++
		if freq[char] > maxCount {
			maxCount = freq[char]
		}

		windowLength := right - left + 1
		if windowLength-maxCount > k {
			freq[rune(s[left])]--
			left++
			windowLength--
		}

		if windowLength > maxLength {
			maxLength = windowLength
		}
	}

	return maxLength
}

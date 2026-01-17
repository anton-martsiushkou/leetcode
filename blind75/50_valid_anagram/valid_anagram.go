package valid_anagram

// IsAnagram checks if t is an anagram of s.
// Uses hash map for O(n) time complexity.
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[rune]int)

	for _, char := range s {
		freq[char]++
	}

	for _, char := range t {
		freq[char]--
		if freq[char] < 0 {
			return false
		}
	}

	return true
}

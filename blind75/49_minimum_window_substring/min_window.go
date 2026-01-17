package min_window

import "math"

// MinWindow finds the minimum window substring containing all characters from t.
// Uses sliding window with hash maps for O(m+n) time complexity.
func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := make(map[byte]int)
	required := len(need)
	formed := 0

	left := 0
	minLen := math.MaxInt32
	minLeft := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		have[char]++

		if need[char] > 0 && have[char] == need[char] {
			formed++
		}

		for formed == required && left <= right {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
			}

			leftChar := s[left]
			have[leftChar]--
			if need[leftChar] > 0 && have[leftChar] < need[leftChar] {
				formed--
			}
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}

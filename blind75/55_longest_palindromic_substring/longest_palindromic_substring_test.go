package longest_palindromic_substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name          string
		s             string
		validOutputs  []string // Multiple valid outputs possible
		expectedLen   int
		checkFunction func(string, string) bool
	}{
		{
			name:         "example 1",
			s:            "babad",
			validOutputs: []string{"bab", "aba"},
			expectedLen:  3,
		},
		{
			name:         "example 2",
			s:            "cbbd",
			validOutputs: []string{"bb"},
			expectedLen:  2,
		},
		{
			name:         "example 3 - single character",
			s:            "a",
			validOutputs: []string{"a"},
			expectedLen:  1,
		},
		{
			name:         "example 4 - multiple single chars",
			s:            "ac",
			validOutputs: []string{"a", "c"},
			expectedLen:  1,
		},
		{
			name:         "all same characters",
			s:            "aaaa",
			validOutputs: []string{"aaaa"},
			expectedLen:  4,
		},
		{
			name:         "no palindrome except singles",
			s:            "abcd",
			validOutputs: []string{"a", "b", "c", "d"},
			expectedLen:  1,
		},
		{
			name:         "palindrome at start",
			s:            "abacabad",
			validOutputs: []string{"abacaba"},
			expectedLen:  7,
		},
		{
			name:         "palindrome at end",
			s:            "xyracecar",
			validOutputs: []string{"racecar"},
			expectedLen:  7,
		},
		{
			name:         "even length palindrome",
			s:            "abba",
			validOutputs: []string{"abba"},
			expectedLen:  4,
		},
		{
			name:         "odd length palindrome",
			s:            "racecar",
			validOutputs: []string{"racecar"},
			expectedLen:  7,
		},
		{
			name:         "multiple palindromes same length",
			s:            "abacdfgdcaba",
			validOutputs: []string{"aba", "aca"},
			expectedLen:  3,
		},
		{
			name:         "two character palindrome",
			s:            "bb",
			validOutputs: []string{"bb"},
			expectedLen:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestPalindrome(tt.s)

			// Check length is correct
			if len(got) != tt.expectedLen {
				t.Errorf("LongestPalindrome(%q) returned length %d, want %d", tt.s, len(got), tt.expectedLen)
				return
			}

			// Check if result is a palindrome
			if !isPalindrome(got) {
				t.Errorf("LongestPalindrome(%q) = %q, which is not a palindrome", tt.s, got)
				return
			}

			// Check if result is in valid outputs
			valid := false
			for _, validOut := range tt.validOutputs {
				if got == validOut {
					valid = true
					break
				}
			}

			if !valid && len(tt.validOutputs) > 0 {
				t.Logf("Note: LongestPalindrome(%q) = %q, expected one of %v (but result is still valid if it's a palindrome of correct length)", tt.s, got, tt.validOutputs)
			}
		})
	}
}

// isPalindrome checks if a string is a palindrome
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

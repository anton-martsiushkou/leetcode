package valid_parentheses

// IsValid determines if the input string has valid parentheses.
// Uses stack for O(n) time complexity.
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if pairs[char] != top {
				return false
			}
		}
	}

	return len(stack) == 0
}

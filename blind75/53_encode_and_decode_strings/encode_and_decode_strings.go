package encode_and_decode_strings

import (
	"strconv"
	"strings"
)

// Encode encodes a list of strings to a single string.
// Uses length-prefix encoding for O(n) time complexity.
func Encode(strs []string) string {
	var builder strings.Builder

	for _, s := range strs {
		builder.WriteString(strconv.Itoa(len(s)))
		builder.WriteByte('#')
		builder.WriteString(s)
	}

	return builder.String()
}

// Decode decodes a single string back to a list of strings.
// Uses length-prefix decoding for O(n) time complexity.
func Decode(s string) []string {
	if s == "" {
		return []string{}
	}

	result := []string{}
	i := 0

	for i < len(s) {
		// Find the delimiter '#'
		j := i
		for j < len(s) && s[j] != '#' {
			j++
		}

		// Extract the length
		length, _ := strconv.Atoi(s[i:j])

		// Move past the '#' delimiter
		j++

		// Extract the string of the specified length
		result = append(result, s[j:j+length])

		// Move index to the next encoded string
		i = j + length
	}

	return result
}

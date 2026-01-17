package codec

import (
	"strconv"
	"strings"
)

// Codec provides methods to encode and decode strings
type Codec struct{}

// Encode encodes a list of strings to a single string.
// Time: O(n), Space: O(n)
func (c *Codec) Encode(strs []string) string {
	var builder strings.Builder

	for _, s := range strs {
		// Format: length + "#" + string
		builder.WriteString(strconv.Itoa(len(s)))
		builder.WriteString("#")
		builder.WriteString(s)
	}

	return builder.String()
}

// Decode decodes a single string to a list of strings.
// Time: O(n), Space: O(n)
func (c *Codec) Decode(s string) []string {
	result := []string{}
	i := 0

	for i < len(s) {
		// Find the delimiter '#'
		j := i
		for s[j] != '#' {
			j++
		}

		// Parse the length
		length, _ := strconv.Atoi(s[i:j])

		// Extract the string using the length
		i = j + 1
		result = append(result, s[i:i+length])

		// Move pointer past the extracted string
		i += length
	}

	return result
}

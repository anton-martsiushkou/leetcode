package group_anagrams

import (
	"sort"
	"strings"
)

// GroupAnagrams groups anagrams together from the input array.
// Uses hash map with sorted string as key for O(n*k*log(k)) time complexity.
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := sortString(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "example 2",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "example 3",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "all same",
			strs: []string{"abc", "bca", "cab"},
			want: [][]string{{"abc", "bca", "cab"}},
		},
		{
			name: "all different",
			strs: []string{"a", "b", "c"},
			want: [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name: "with empty string",
			strs: []string{"", "b", ""},
			want: [][]string{{"", ""}, {"b"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.strs)
			if !equalAnagramGroups(got, tt.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalAnagramGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	sortGroups := func(groups [][]string) {
		for _, group := range groups {
			sort.Strings(group)
		}
		sort.Slice(groups, func(i, j int) bool {
			if len(groups[i]) != len(groups[j]) {
				return len(groups[i]) < len(groups[j])
			}
			return groups[i][0] < groups[j][0]
		})
	}

	sortGroups(a)
	sortGroups(b)

	return reflect.DeepEqual(a, b)
}

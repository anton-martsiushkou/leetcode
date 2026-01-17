package word_dictionary

import (
	"testing"
)

func TestWordDictionary(t *testing.T) {
	wd := Constructor()

	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")

	tests := []struct {
		name string
		word string
		want bool
	}{
		{"exact match exists", "bad", true},
		{"exact match doesn't exist", "pad", false},
		{"wildcard at start", ".ad", true},
		{"wildcard at end", "ba.", true},
		{"multiple wildcards", "b..", true},
		{"all wildcards", "...", true},
		{"wildcard no match", ".a.", true},
		{"wrong length", "b", false},
		{"longer word", "badder", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wd.Search(tt.word)
			if got != tt.want {
				t.Errorf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestWordDictionaryComplex(t *testing.T) {
	wd := Constructor()

	// Add various words
	words := []string{"a", "ab", "abc", "abcd", "abcde"}
	for _, word := range words {
		wd.AddWord(word)
	}

	tests := []struct {
		name string
		word string
		want bool
	}{
		{"single char", "a", true},
		{"single char wildcard", ".", true},
		{"two chars exact", "ab", true},
		{"two chars wildcard", ".b", true},
		{"two chars wildcard both", "..", true},
		{"five chars exact", "abcde", true},
		{"five chars with wildcards", "a.c.e", true},
		{"wrong pattern", "a.d", false},
		{"too long", "abcdef", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wd.Search(tt.word)
			if got != tt.want {
				t.Errorf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestWordDictionaryNoBranch(t *testing.T) {
	wd := Constructor()

	wd.AddWord("test")

	tests := []struct {
		name string
		word string
		want bool
	}{
		{"exact match", "test", true},
		{"wildcard match", "t..t", true},
		{"all wildcards", "....", true},
		{"no match - different char", "best", false},
		{"no match - wildcard wrong length", "t..", false},
		{"no match - too long", "tests", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wd.Search(tt.word)
			if got != tt.want {
				t.Errorf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestWordDictionaryEmpty(t *testing.T) {
	wd := Constructor()

	if wd.Search("anything") {
		t.Error("Empty dictionary should not match any word")
	}

	if wd.Search(".") {
		t.Error("Empty dictionary should not match wildcard")
	}
}

package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	// Test insert and search
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Error("Expected to find 'apple'")
	}

	if trie.Search("app") {
		t.Error("Did not expect to find 'app'")
	}

	if !trie.StartsWith("app") {
		t.Error("Expected 'app' to be a prefix")
	}

	trie.Insert("app")
	if !trie.Search("app") {
		t.Error("Expected to find 'app' after insertion")
	}
}

func TestTrieMultipleWords(t *testing.T) {
	trie := Constructor()

	words := []string{"apple", "app", "application", "apply", "banana", "band"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test all inserted words can be found
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected to find word '%s'", word)
		}
	}

	// Test non-existent words
	nonExistent := []string{"appl", "ban", "banan", "apps", "bandana"}
	for _, word := range nonExistent {
		if trie.Search(word) {
			t.Errorf("Did not expect to find word '%s'", word)
		}
	}

	// Test prefixes
	prefixes := []string{"app", "appl", "appli", "ban", "band"}
	for _, prefix := range prefixes {
		if !trie.StartsWith(prefix) {
			t.Errorf("Expected '%s' to be a valid prefix", prefix)
		}
	}

	// Test invalid prefixes
	invalidPrefixes := []string{"orange", "grape", "car"}
	for _, prefix := range invalidPrefixes {
		if trie.StartsWith(prefix) {
			t.Errorf("Did not expect '%s' to be a valid prefix", prefix)
		}
	}
}

func TestTrieEmptyAndSingleChar(t *testing.T) {
	trie := Constructor()

	// Test single character
	trie.Insert("a")
	if !trie.Search("a") {
		t.Error("Expected to find 'a'")
	}

	if !trie.StartsWith("a") {
		t.Error("Expected 'a' to be a prefix")
	}

	// Test that empty search doesn't match
	if trie.Search("") {
		t.Error("Did not expect to find empty string")
	}
}

func TestTrieOverlappingWords(t *testing.T) {
	trie := Constructor()

	// Test words that are prefixes of each other
	trie.Insert("test")
	trie.Insert("testing")
	trie.Insert("tester")

	if !trie.Search("test") {
		t.Error("Expected to find 'test'")
	}

	if !trie.Search("testing") {
		t.Error("Expected to find 'testing'")
	}

	if !trie.Search("tester") {
		t.Error("Expected to find 'tester'")
	}

	if trie.Search("tes") {
		t.Error("Did not expect to find 'tes'")
	}

	if !trie.StartsWith("test") {
		t.Error("Expected 'test' to be a prefix")
	}
}

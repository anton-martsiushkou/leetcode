package trie

// TrieNode represents a node in the trie.
type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

// NewTrieNode creates a new trie node.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children:    make(map[rune]*TrieNode),
		isEndOfWord: false,
	}
}

// Trie represents a prefix tree data structure.
type Trie struct {
	root *TrieNode
}

// Constructor initializes the trie object.
func Constructor() Trie {
	return Trie{
		root: NewTrieNode(),
	}
}

// Insert inserts a word into the trie.
// Time complexity: O(m) where m is the length of the word.
func (t *Trie) Insert(word string) {
	current := t.root

	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = NewTrieNode()
		}
		current = current.children[ch]
	}

	current.isEndOfWord = true
}

// Search returns true if the word is in the trie.
// Time complexity: O(m) where m is the length of the word.
func (t *Trie) Search(word string) bool {
	current := t.root

	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}

	return current.isEndOfWord
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix.
// Time complexity: O(m) where m is the length of the prefix.
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root

	for _, ch := range prefix {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}

	return true
}

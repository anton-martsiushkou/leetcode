package word_dictionary

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

// WordDictionary supports adding words and searching with wildcards.
type WordDictionary struct {
	root *TrieNode
}

// Constructor initializes the WordDictionary object.
func Constructor() WordDictionary {
	return WordDictionary{
		root: NewTrieNode(),
	}
}

// AddWord adds a word to the data structure.
// Time complexity: O(m) where m is the length of the word.
func (wd *WordDictionary) AddWord(word string) {
	current := wd.root

	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = NewTrieNode()
		}
		current = current.children[ch]
	}

	current.isEndOfWord = true
}

// Search returns true if the word is in the data structure.
// The word may contain dots '.' which can match any letter.
// Time complexity: O(m) for exact match, O(26^k × m) with k wildcards.
func (wd *WordDictionary) Search(word string) bool {
	return wd.searchHelper([]rune(word), 0, wd.root)
}

// searchHelper is a recursive helper for wildcard search.
func (wd *WordDictionary) searchHelper(word []rune, index int, node *TrieNode) bool {
	// Base case: reached end of word
	if index == len(word) {
		return node.isEndOfWord
	}

	ch := word[index]

	// Wildcard: try all possible children
	if ch == '.' {
		for _, child := range node.children {
			if wd.searchHelper(word, index+1, child) {
				return true
			}
		}
		return false
	}

	// Regular character: search specific child
	if child, exists := node.children[ch]; exists {
		return wd.searchHelper(word, index+1, child)
	}

	return false
}

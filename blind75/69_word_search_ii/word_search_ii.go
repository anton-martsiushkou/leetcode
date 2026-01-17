package word_search_ii

// TrieNode represents a node in the trie.
type TrieNode struct {
	children map[byte]*TrieNode
	word     string // Stores the complete word at leaf nodes
}

// NewTrieNode creates a new trie node.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		word:     "",
	}
}

// FindWords returns all words from the list that exist on the board.
// Uses Trie + DFS backtracking for O(M×N×4^L) time complexity.
func FindWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return []string{}
	}

	// Build Trie from words
	root := buildTrie(words)

	result := []string{}
	m, n := len(board), len(board[0])

	// Try starting DFS from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, root, &result)
		}
	}

	return result
}

// buildTrie constructs a Trie from the list of words.
func buildTrie(words []string) *TrieNode {
	root := NewTrieNode()

	for _, word := range words {
		current := root
		for i := 0; i < len(word); i++ {
			ch := word[i]
			if _, exists := current.children[ch]; !exists {
				current.children[ch] = NewTrieNode()
			}
			current = current.children[ch]
		}
		current.word = word // Mark end of word
	}

	return root
}

// dfs performs depth-first search with backtracking.
func dfs(board [][]byte, i, j int, node *TrieNode, result *[]string) {
	// Boundary check
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}

	ch := board[i][j]

	// Check if already visited or character not in Trie
	if ch == '#' || node.children[ch] == nil {
		return
	}

	// Move to next Trie node
	node = node.children[ch]

	// Found a word
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = "" // Avoid duplicate results
	}

	// Mark cell as visited
	board[i][j] = '#'

	// Explore all 4 directions
	dfs(board, i-1, j, node, result) // up
	dfs(board, i+1, j, node, result) // down
	dfs(board, i, j-1, node, result) // left
	dfs(board, i, j+1, node, result) // right

	// Backtrack: restore cell
	board[i][j] = ch
}

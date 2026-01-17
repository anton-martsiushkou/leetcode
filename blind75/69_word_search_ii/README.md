# Word Search II

## Problem Description

Given an `m x n` board of characters and a list of strings `words`, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

## Examples

**Example 1:**
```
Input: board = [["o","a","a","n"],
                ["e","t","a","e"],
                ["i","h","k","r"],
                ["i","f","l","v"]],
       words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
```

**Example 2:**
```
Input: board = [["a","b"],
                ["c","d"]],
       words = ["abcb"]
Output: []
```

## Constraints

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 12`
- `board[i][j]` is a lowercase English letter
- `1 <= words.length <= 3 * 10^4`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters
- All the strings of `words` are unique

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Trie + DFS Backtracking | Trie + Grid | O(M × N × 4^L) | O(W × L) |

Where:
- M × N = board dimensions
- L = maximum word length
- W = number of words

### Approach: Trie + DFS Backtracking

The key insight is to combine a Trie data structure with DFS backtracking to efficiently search for multiple words simultaneously.

**Key Insight:** Instead of searching for each word individually (which would be O(W × M × N × 4^L)), build a Trie of all words and perform a single DFS traversal. The Trie allows us to prune searches early when no word has the current prefix.

### Algorithm Steps

1. **Build Trie**: Insert all words into a Trie
2. **DFS Traversal**: For each cell in the board:
   - Start DFS if the cell's character exists in Trie root
   - During DFS:
     - Mark current cell as visited (modify in-place or use visited set)
     - If current Trie node marks end of word, add word to results
     - Explore 4 adjacent cells (up, down, left, right)
     - Only continue DFS if next character exists in Trie
     - Backtrack: restore cell state
3. **Optimization**: Remove found words from Trie to avoid duplicates

### Example Walkthrough

For board:
```
o a a n
e t a e
i h k r
i f l v
```

Words: ["oath", "pea", "eat", "rain"]

**Build Trie:**
```
      root
     /  |  \
    o   p   e   r
    |   |   |   |
    a   e   a   a
    |   |   |   |
    t   a   t   i
    |           |
    h           n
```

**DFS from (0,0) 'o':**
1. 'o' exists in Trie, continue
2. Try adjacent: (0,1) 'a', (1,0) 'e'
3. Follow path o→a→t→h at positions (0,0)→(0,1)→(1,1)→(2,1)
4. Found "oath", add to results

**DFS from (1,2) 'a':**
1. Follow path e→a→t at positions (1,1)→(1,2)→(0,2)
2. Found "eat", add to results

### Why This is Optimal

- **Time Complexity O(M × N × 4^L)**:
  - Visit each cell once as starting point: O(M × N)
  - From each cell, explore up to 4^L paths (4 directions, L levels deep)
  - Trie pruning significantly reduces actual paths explored

- **Space Complexity O(W × L)**:
  - Trie stores all words: O(W × L)
  - DFS recursion stack: O(L)
  - Result storage: O(W × L)

### Optimizations

1. **Trie Pruning**: Remove words from Trie once found to avoid duplicates
2. **Early Termination**: Stop DFS when current prefix doesn't match any word
3. **In-place Modification**: Mark visited cells by modifying board instead of using separate visited set
4. **Leaf Node Removal**: After finding a word, remove that branch from Trie if it's a leaf

### Comparison to Naive Approach

**Naive (DFS for each word separately):**
- Time: O(W × M × N × 4^L)
- Searches board W times

**Trie + DFS (this solution):**
- Time: O(M × N × 4^L)
- Searches board once, simultaneously looking for all words
- Speedup factor: approximately W times faster

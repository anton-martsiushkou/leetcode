# Word Search

## Problem Description

Given an `m x n` grid of characters `board` and a string `word`, return `true` if `word` exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

## Examples

**Example 1:**
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
```

**Example 2:**
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
```

**Example 3:**
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
```

## Constraints

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 6`
- `1 <= word.length <= 15`
- `board` and `word` consists of only lowercase and uppercase English letters.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Backtracking with DFS | Recursion Stack + Matrix | O(m × n × 4^L) | O(L) |

### Approach: Backtracking with DFS

The optimal solution uses depth-first search (DFS) with backtracking to explore all possible paths in the grid. We try each cell as a starting point and recursively search for the word by exploring all four adjacent directions.

**Key Insight:** This is a path-finding problem that requires exploring all possible paths. Backtracking allows us to mark cells as visited during exploration and unmark them when backtracking, enabling us to reuse cells for different paths.

### Algorithm Steps

1. **Iterate through each cell** in the board as a potential starting point
2. **For each cell**, call a DFS function with parameters:
   - Current position (i, j)
   - Current index in the word we're matching
3. **DFS Base Cases**:
   - If index equals word length, we found the complete word - return `true`
   - If out of bounds, or cell doesn't match word[index], or cell already visited - return `false`
4. **DFS Recursive Case**:
   - Mark current cell as visited (temporarily modify it or use a visited set)
   - Recursively search all 4 adjacent cells (up, down, left, right)
   - If any direction returns `true`, return `true`
   - Unmark the cell (backtrack) before returning `false`
5. **Return** `true` if any starting position finds the word, otherwise `false`

### Example Walkthrough

For `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]` and `word = "ABCCED"`:

1. **Start at (0,0) 'A'**: Matches word[0]
   - Mark visited, search neighbors

2. **Move to (0,1) 'B'**: Matches word[1]
   - Mark visited, search neighbors

3. **Move to (0,2) 'C'**: Matches word[2]
   - Mark visited, search neighbors

4. **Move to (1,2) 'C'**: Matches word[3]
   - Mark visited, search neighbors

5. **Move to (2,2) 'E'**: Matches word[4]
   - Mark visited, search neighbors

6. **Move to (2,3) 'D'**: Doesn't match word[5] 'D'
   - Backtrack and try (2,1) 'D'

7. **Move to (2,1) 'D'**: Matches word[5]
   - Index equals word length - return `true`

### Why This is Optimal

- **Time Complexity O(m × n × 4^L)**: For each of the m×n cells as starting points, we might explore 4 directions at each step for up to L characters (where L is word length). In practice, pruning makes it much faster.
- **Space Complexity O(L)**: The recursion depth is at most L (word length), and we modify the board in-place for marking visited cells
- This is optimal for the general case as we must potentially explore all paths to determine if the word exists

### Alternative Approaches

1. **Trie + DFS**: O(m × n × 4^L) time, O(K) space where K is total characters in dictionary - Better when searching for multiple words simultaneously
2. **BFS**: O(m × n × 4^L) time, O(m × n) space - Uses more memory for the queue, no advantage over DFS for this problem

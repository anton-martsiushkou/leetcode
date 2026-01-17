# Number of Islands

## Problem Description

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## Examples

**Example 1:**
```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

**Example 2:**
```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
```

## Constraints

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS | Recursion Stack | O(m × n) | O(m × n) |
| BFS | Queue | O(m × n) | O(min(m, n)) |
| Union Find | Disjoint Set | O(m × n × α(m×n)) | O(m × n) |

### Approach: DFS/BFS Island Traversal

The key insight is to iterate through the grid, and whenever we find a land cell ('1') that hasn't been visited, we increment our island count and perform a DFS/BFS to mark all connected land cells as visited.

**Key Insight:** Each time we encounter an unvisited land cell ('1'), we've found a new island. We then explore all connected land cells (horizontally and vertically adjacent) and mark them as visited to avoid counting the same island multiple times.

### Algorithm Steps (DFS)

1. Initialize island count to 0
2. Iterate through each cell in the grid
3. When we find a land cell ('1'):
   - Increment island count
   - Perform DFS to mark all connected land cells:
     - Mark current cell as visited (change '1' to '0' or use a visited set)
     - Recursively visit all 4 adjacent cells (up, down, left, right)
     - Only continue if the adjacent cell is land ('1') and within bounds
4. Return the island count

### Algorithm Steps (BFS)

1. Initialize island count to 0
2. Iterate through each cell in the grid
3. When we find a land cell ('1'):
   - Increment island count
   - Add cell to queue and perform BFS:
     - While queue is not empty:
       - Remove cell from queue
       - Mark as visited
       - Add all adjacent land cells to queue
4. Return the island count

### Example Walkthrough

For the grid:
```
["1","1","0"]
["1","0","0"]
["0","0","1"]
```

1. Start at (0,0) which is '1' → island count = 1
   - DFS visits: (0,0) → (0,1) → (1,0)
   - All three cells marked as visited

2. Continue scanning, next unvisited '1' is at (2,2) → island count = 2
   - DFS visits: (2,2)
   - Mark as visited

3. No more unvisited land cells
4. Return 2

### Why This is Optimal

- **Time Complexity O(m × n)**: We visit each cell at most once
- **Space Complexity**:
  - DFS: O(m × n) in worst case (if entire grid is land, recursion depth)
  - BFS: O(min(m, n)) for the queue in worst case
- This is optimal because we must examine every cell at least once to count all islands

# Pacific Atlantic Water Flow

## Problem Description

There is an `m x n` rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an `m x n` integer matrix `heights` where `heights[r][c]` represents the height above sea level of the cell at coordinate `(r, c)`.

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates `result` where `result[i] = [ri, ci]` denotes that rain water can flow from cell `(ri, ci)` to both the Pacific and Atlantic oceans.

## Examples

**Example 1:**
```
Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
Explanation: The following cells can flow to the Pacific and Atlantic oceans:
[0,4]: [0,4] -> Pacific Ocean
       [0,4] -> Atlantic Ocean
[1,3]: [1,3] -> [0,3] -> Pacific Ocean
       [1,3] -> [1,4] -> Atlantic Ocean
[1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean
       [1,4] -> Atlantic Ocean
[2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean
       [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
[3,0]: [3,0] -> Pacific Ocean
       [3,0] -> [4,0] -> Atlantic Ocean
[3,1]: [3,1] -> [3,0] -> Pacific Ocean
       [3,1] -> [4,1] -> Atlantic Ocean
[4,0]: [4,0] -> Pacific Ocean
       [4,0] -> Atlantic Ocean
```

**Example 2:**
```
Input: heights = [[1]]
Output: [[0,0]]
Explanation: The water can flow from the only cell to both oceans.
```

## Constraints

- `m == heights.length`
- `n == heights[r].length`
- `1 <= m, n <= 200`
- `0 <= heights[r][c] <= 10^5`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS from Ocean Borders | Set + Recursion | O(m × n) | O(m × n) |
| BFS from Ocean Borders | Set + Queue | O(m × n) | O(m × n) |

### Approach: Reverse Flow from Oceans

Instead of checking from each cell whether water can reach both oceans (which would be inefficient), we reverse the problem: start from the ocean borders and find all cells that can reach each ocean.

**Key Insight:** Start DFS/BFS from the borders (Pacific and Atlantic) and mark all cells reachable from each ocean. The answer is the intersection of cells reachable from both oceans. Water flows "uphill" in reverse (from ocean to higher or equal cells).

### Algorithm Steps (DFS)

1. Create two boolean matrices to track cells reachable from Pacific and Atlantic
2. Perform DFS from Pacific borders:
   - Top row (all columns)
   - Left column (all rows)
3. Perform DFS from Atlantic borders:
   - Bottom row (all columns)
   - Right column (all rows)
4. During DFS, water can flow to a cell if:
   - The cell is not yet visited from this ocean
   - The cell's height is >= previous cell's height (reverse flow)
5. Find intersection: cells reachable from both Pacific and Atlantic
6. Return the result as a list of coordinates

### DFS Helper Function

```
dfs(row, col, visited, prevHeight):
    if out of bounds or already visited:
        return
    if heights[row][col] < prevHeight:
        return  // Water can't flow uphill

    mark visited
    explore 4 directions (up, down, left, right)
```

### Example Walkthrough

For the matrix:
```
[1, 2, 2, 3, 5]
[3, 2, 3, 4, 4]
[2, 4, 5, 3, 1]
[6, 7, 1, 4, 5]
[5, 1, 1, 2, 4]
```

**Pacific reachable cells (starting from top and left):**
- All cells in first row and column
- Cells reachable by going to equal or higher heights

**Atlantic reachable cells (starting from bottom and right):**
- All cells in last row and column
- Cells reachable by going to equal or higher heights

**Intersection:** Cells like (0,4), (1,3), (1,4), (2,2), etc. that can reach both oceans

### Why This is Optimal

- **Time Complexity O(m × n)**: Each cell is visited at most twice (once for each ocean)
- **Space Complexity O(m × n)**: Two visited matrices of size m × n
- This is optimal because we must examine each cell to determine if it can reach both oceans

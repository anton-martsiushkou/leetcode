# Unique Paths

## Problem Description

There is a robot on an `m x n` grid. The robot is initially located at the **top-left corner** (i.e., `grid[0][0]`). The robot tries to move to the **bottom-right corner** (i.e., `grid[m - 1][n - 1]`). The robot can only move either down or right at any point in time.

Given the two integers `m` and `n`, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to `2 * 10^9`.

## Examples

**Example 1:**
```
Input: m = 3, n = 7
Output: 28
```

**Example 2:**
```
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
```

## Constraints

- `1 <= m, n <= 100`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming | 1D Array | O(m * n) | O(n) |

### Approach: Dynamic Programming

The optimal solution uses dynamic programming to build up the number of unique paths to each cell. The key insight is that to reach any cell (i, j), we can only come from the cell above (i-1, j) or the cell to the left (i, j-1).

**Key Insight:** For any cell (i, j), the number of unique paths to reach it is the sum of:
- Unique paths to reach cell (i-1, j) (cell above)
- Unique paths to reach cell (i, j-1) (cell to the left)

This gives us the recurrence relation: `dp[i][j] = dp[i-1][j] + dp[i][j-1]`

### Algorithm Steps

1. **Base Cases:**
   - First row: only one way to reach any cell (keep moving right)
   - First column: only one way to reach any cell (keep moving down)

2. **DP Recurrence:**
   - Initialize a 1D array of size n with all 1s (first row)
   - For each row i from 1 to m-1:
     - For each column j from 1 to n-1:
       - dp[j] = dp[j] + dp[j-1]
       - (dp[j] is paths from above, dp[j-1] is paths from left)

3. **Return** dp[n-1] (bottom-right corner)

### Example Walkthrough

For `m = 3, n = 3`:

**Initial grid (paths to each cell):**
```
1  1  1
1  ?  ?
1  ?  ?
```

**After processing row 1:**
```
1  1  1
1  2  3
1  ?  ?
```
- (1,1): 1 (from above) + 1 (from left) = 2
- (1,2): 1 (from above) + 2 (from left) = 3

**After processing row 2:**
```
1  1  1
1  2  3
1  3  6
```
- (2,1): 2 (from above) + 1 (from left) = 3
- (2,2): 3 (from above) + 3 (from left) = 6

**Result:** 6 unique paths

### Why This is Optimal

- **Time Complexity O(m * n)**: We process each cell exactly once
- **Space Complexity O(n)**: We only need a single row array, reusing it for each row calculation
- This is optimal because we must consider every cell to calculate the total number of paths
- We can further optimize to O(min(m, n)) space by iterating over the smaller dimension

### Alternative Approaches

1. **Recursive (Brute Force)**: O(2^(m+n)) time - Too slow, explores all possible paths
2. **2D DP Array**: O(m * n) time, O(m * n) space - Same time complexity but uses more space
3. **Combinatorics**: O(min(m, n)) time, O(1) space - Calculate C(m+n-2, m-1), but may overflow for large inputs
4. **Memoization**: O(m * n) time, O(m * n) space - Top-down approach with similar complexity

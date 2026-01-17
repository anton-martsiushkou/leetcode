# Set Matrix Zeroes

## Problem Description

Given an `m x n` integer matrix `matrix`, if an element is `0`, set its entire row and column to `0`'s.

You must do it **in place**.

## Examples

**Example 1:**
```
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
```

**Example 2:**
```
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
```

## Constraints

- `m == matrix.length`
- `n == matrix[0].length`
- `1 <= m, n <= 200`
- `-2^31 <= matrix[i][j] <= 2^31 - 1`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| First Row/Column as Markers | Matrix (in-place) | O(m × n) | O(1) |

### Approach: First Row/Column as Markers

The optimal solution uses the first row and first column of the matrix itself to store information about which rows and columns need to be zeroed. This achieves O(1) space complexity by avoiding the need for separate arrays.

**Key Insight:** Instead of using additional O(m + n) space to track which rows and columns contain zeros, we can repurpose the first row and first column of the matrix as markers.

### Algorithm Steps

1. **Check if first row/column have zeros**: Use two boolean flags to remember if the first row or first column originally contained any zeros
2. **Mark zeros in first row/column**: Iterate through the matrix (starting from index [1,1]). When we find a zero at position [i,j], set `matrix[i][0] = 0` and `matrix[0][j] = 0` as markers
3. **Set zeros based on markers**: Iterate through the matrix again (excluding first row/column). If `matrix[i][0] == 0` or `matrix[0][j] == 0`, set `matrix[i][j] = 0`
4. **Handle first row/column**: Finally, use the boolean flags from step 1 to zero out the first row and/or first column if needed

### Example Walkthrough

For `matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]`:

1. **Initial state**: First row has zeros, so `firstRowHasZero = true`. First column has zero, so `firstColHasZero = true`

2. **Mark phase**: Scan matrix[1..m-1][1..n-1]
   - All elements are non-zero, so no additional markers needed
   - First row already marked: [0, _, _, 0]

3. **Zero phase**: For each cell, check if its row marker or column marker is zero
   - Row 1: matrix[1][0] = 3 (not zero), but matrix[0][1] = 1, matrix[0][2] = 2, matrix[0][3] = 0
   - Set matrix[1][3] = 0
   - Row 2: matrix[2][0] = 1 (not zero), set matrix[2][3] = 0

4. **First row/column**: Zero them out based on flags
   - Result: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

### Why This is Optimal

- **Time Complexity O(m × n)**: We traverse the entire matrix a constant number of times (3 passes)
- **Space Complexity O(1)**: We only use two boolean variables for tracking, and repurpose the matrix itself for marking
- This is optimal because we must examine every element at least once to determine which rows/columns to zero

### Alternative Approaches

1. **Using Sets**: O(m × n) time, O(m + n) space - Uses additional space to store row and column indices
2. **Using Additional Arrays**: O(m × n) time, O(m + n) space - Similar to sets but with boolean arrays

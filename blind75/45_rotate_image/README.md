# Rotate Image

## Problem Description

You are given an `n x n` 2D matrix representing an image, rotate the image by **90 degrees** (clockwise).

You have to rotate the image **in-place**, which means you have to modify the input 2D matrix directly. **DO NOT** allocate another 2D matrix and do the rotation.

## Examples

**Example 1:**
```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
```

**Example 2:**
```
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
```

## Constraints

- `n == matrix.length == matrix[i].length`
- `1 <= n <= 20`
- `-1000 <= matrix[i][j] <= 1000`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Transpose + Reverse | Matrix (in-place) | O(n²) | O(1) |

### Approach: Transpose + Reverse

The optimal solution achieves 90-degree clockwise rotation through two simple transformations: first transpose the matrix (swap rows and columns), then reverse each row.

**Key Insight:** A 90-degree clockwise rotation can be decomposed into two simpler operations:
1. Transpose: `matrix[i][j]` becomes `matrix[j][i]`
2. Reverse each row: `matrix[i]` becomes `matrix[i].reverse()`

This equivalence: `rotate_90_clockwise = transpose + reverse_rows`

### Algorithm Steps

1. **Transpose the matrix**: Swap elements across the diagonal
   - For each `i` from 0 to n-1:
     - For each `j` from i+1 to n-1:
       - Swap `matrix[i][j]` with `matrix[j][i]`
2. **Reverse each row**: Reverse the elements in each row
   - For each row `i` from 0 to n-1:
     - Reverse `matrix[i]` in place

### Example Walkthrough

For `matrix = [[1,2,3],[4,5,6],[7,8,9]]`:

1. **Original**:
   ```
   1 2 3
   4 5 6
   7 8 9
   ```

2. **After transpose** (swap matrix[i][j] with matrix[j][i]):
   ```
   1 4 7
   2 5 8
   3 6 9
   ```

3. **After reverse each row**:
   ```
   7 4 1
   8 5 2
   9 6 3
   ```

### Why This is Optimal

- **Time Complexity O(n²)**: We visit each element twice (once in transpose, once in reverse), but this is still O(n²) since the matrix has n² elements
- **Space Complexity O(1)**: We only swap elements in place without using additional data structures
- This is optimal because we must modify every element at least once to rotate the matrix

### Alternative Approaches

1. **Rotate Layer by Layer**: O(n²) time, O(1) space - Rotates elements in groups of 4, starting from outer layer to inner, more complex to implement
2. **Using Extra Matrix**: O(n²) time, O(n²) space - Creates a new matrix and copies elements to rotated positions, violates in-place requirement

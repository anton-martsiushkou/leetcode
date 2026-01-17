# Spiral Matrix

## Problem Description

Given an `m x n` matrix, return all elements of the matrix in spiral order.

## Examples

**Example 1:**
```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

**Example 2:**
```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

## Constraints

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `-100 <= matrix[i][j] <= 100`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Layer-by-Layer (Boundary Tracking) | Array | O(m × n) | O(1) |

### Approach: Layer-by-Layer (Boundary Tracking)

The optimal solution traverses the matrix in a spiral pattern by maintaining four boundaries (top, bottom, left, right) and shrinking them inward after completing each edge of the spiral.

**Key Insight:** A spiral traversal consists of moving right along the top edge, down the right edge, left along the bottom edge, and up the left edge. After each complete edge, we shrink the corresponding boundary inward.

### Algorithm Steps

1. **Initialize boundaries**: Set `top = 0`, `bottom = m-1`, `left = 0`, `right = n-1`
2. **Traverse in spiral order**:
   - Move **right**: From `left` to `right` on row `top`, then increment `top`
   - Move **down**: From `top` to `bottom` on column `right`, then decrement `right`
   - Move **left**: From `right` to `left` on row `bottom` (if `top <= bottom`), then decrement `bottom`
   - Move **up**: From `bottom` to `top` on column `left` (if `left <= right`), then increment `left`
3. **Repeat** until all boundaries cross each other
4. **Return** the collected elements

### Example Walkthrough

For `matrix = [[1,2,3],[4,5,6],[7,8,9]]`:

1. **Initial**: top=0, bottom=2, left=0, right=2, result=[]

2. **First spiral**:
   - Right: [1,2,3], top becomes 1
   - Down: [3,6,9], right becomes 1
   - Left: [9,8,7], bottom becomes 1
   - Up: [7,4], left becomes 1

3. **Second spiral**:
   - Right: [4,5], top becomes 2
   - Down: [5,6], right becomes 0
   - Boundaries crossed, stop

4. **Result**: [1,2,3,6,9,8,7,4,5]

### Why This is Optimal

- **Time Complexity O(m × n)**: We visit each element exactly once
- **Space Complexity O(1)**: We only use a constant amount of extra space for boundary variables (output array doesn't count toward space complexity)
- This is optimal because we must visit every element at least once to include it in the result

### Alternative Approaches

1. **Direction Vectors with Visited Array**: O(m × n) time, O(m × n) space - Uses a visited matrix to track which cells have been traversed, less efficient in space
2. **Recursive**: O(m × n) time, O(min(m,n)) space - Recursively processes outer layer then inner matrix, uses call stack space

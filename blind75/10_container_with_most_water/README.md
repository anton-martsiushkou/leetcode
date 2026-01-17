# Container With Most Water

## Problem Description

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `i`th line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

**Notice** that you may not slant the container.

## Examples

**Example 1:**
```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
In this case, the max area of water (blue section) the container can contain is 49.
The lines at index 1 and 8 form the container.
```

**Example 2:**
```
Input: height = [1,1]
Output: 1
```

## Constraints

- `n == height.length`
- `2 <= n <= 10^5`
- `0 <= height[i] <= 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Two Pointers | Array | O(n) | O(1) |

### Approach: Two Pointers

The optimal solution uses a two-pointer approach starting from both ends of the array and moving inward.

The area of water between two lines is determined by:
- **Width**: Distance between the two lines (right - left)
- **Height**: Minimum of the two line heights (min(height[left], height[right]))
- **Area**: width × height

The key insight is that we want to maximize the area, and we can do this greedily:
- Start with the widest container (maximum width)
- Move the pointer with the shorter height inward, because:
  - Moving the taller line can only decrease the area (width decreases, height stays same or decreases)
  - Moving the shorter line might increase the area (if we find a taller line, height could increase enough to compensate for width decrease)

**Key Insight:** The area is limited by the shorter line. To potentially find a larger area, we must move the pointer at the shorter line, hoping to find a taller line that compensates for the reduced width.

### Algorithm Steps

1. Initialize two pointers: `left = 0`, `right = n - 1`
2. Initialize `max_area = 0`
3. While `left < right`:
   - Calculate current area: `area = (right - left) × min(height[left], height[right])`
   - Update `max_area` if current area is larger
   - Move the pointer with the shorter height:
     - If `height[left] < height[right]`: increment `left`
     - Otherwise: decrement `right`
4. Return `max_area`

### Example Walkthrough

For `height = [1, 8, 6, 2, 5, 4, 8, 3, 7]`:

| Iteration | left | right | height[left] | height[right] | width | min_height | area | max_area | Move |
|-----------|------|-------|--------------|---------------|-------|------------|------|----------|------|
| 1 | 0 | 8 | 1 | 7 | 8 | 1 | 8 | 8 | left++ (1 < 7) |
| 2 | 1 | 8 | 8 | 7 | 7 | 7 | 49 | 49 | right-- (8 > 7) |
| 3 | 1 | 7 | 8 | 3 | 6 | 3 | 18 | 49 | right-- (8 > 3) |
| 4 | 1 | 6 | 8 | 8 | 5 | 8 | 40 | 49 | left++ (8 = 8) |
| 5 | 2 | 6 | 6 | 8 | 4 | 6 | 24 | 49 | left++ (6 < 8) |
| 6 | 3 | 6 | 2 | 8 | 3 | 2 | 6 | 49 | left++ (2 < 8) |
| 7 | 4 | 6 | 5 | 8 | 2 | 5 | 10 | 49 | left++ (5 < 8) |
| 8 | 5 | 6 | 4 | 8 | 1 | 4 | 4 | 49 | left++ (4 < 8) |
| - | 6 | 6 | - | - | - | - | - | - | left == right, stop |

Final answer: 49 (container formed by lines at indices 1 and 8)

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the array once with two pointers moving toward each other
- **Space Complexity O(1)**: We only use a constant amount of extra space
- This is optimal because we must examine the configuration to find the maximum, and O(n) is the best we can do
- The greedy approach works because:
  - We start with maximum width
  - Moving the shorter line is the only way to potentially increase area
  - We never miss the optimal solution because we consider all viable candidates

### Why the Greedy Approach Works

Proof by contradiction: Suppose the optimal solution is at indices (i, j) where i < j. Our algorithm will find it because:
1. When pointers are at positions (a, b) where a ≤ i and j ≤ b, one of these must be true:
   - If we're at (i, b) and height[i] < height[b], we'll keep i and move b toward j
   - If we're at (a, j) and height[a] < height[j], we'll keep j and move a toward i
2. We won't skip (i, j) because we only move the shorter line, which is the correct greedy choice

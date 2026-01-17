# Maximum Subarray

## Problem Description

Given an integer array `nums`, find the subarray with the largest sum, and return its sum.

A subarray is a contiguous non-empty sequence of elements within an array.

## Examples

**Example 1:**
```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
```

**Example 2:**
```
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
```

**Example 3:**
```
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
```

## Constraints

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

**Follow up:** If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Kadane's Algorithm | Array | O(n) | O(1) |

### Approach: Kadane's Algorithm

The optimal solution uses Kadane's Algorithm, a dynamic programming approach that efficiently finds the maximum subarray sum in a single pass.

The core idea is to maintain two values:
- `current_sum`: The maximum sum of a subarray ending at the current position
- `max_sum`: The overall maximum sum found so far

At each position, we decide whether to extend the existing subarray or start a new one.

**Key Insight:** For any element, the maximum subarray ending at that position is either:
1. The element itself (start a new subarray)
2. The element plus the maximum subarray ending at the previous position (extend existing subarray)

We take the maximum of these two options: `current_sum = max(num, current_sum + num)`

### Algorithm Steps

1. Initialize `max_sum` to the first element (or negative infinity)
2. Initialize `current_sum` to 0
3. For each element `num` in the array:
   - Add `num` to `current_sum`
   - Update `max_sum` if `current_sum` is greater
   - If `current_sum` becomes negative, reset it to 0 (start fresh)
4. Return `max_sum`

**Alternative formulation (more common):**
1. Initialize both `max_sum` and `current_sum` to first element
2. For each element from index 1 onwards:
   - `current_sum = max(num, current_sum + num)`
   - `max_sum = max(max_sum, current_sum)`
3. Return `max_sum`

### Example Walkthrough

For `nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]`:

| Index | num | current_sum | max_sum | Explanation |
|-------|-----|-------------|---------|-------------|
| 0 | -2 | -2 | -2 | Start with first element |
| 1 | 1 | 1 | 1 | Start new subarray (1 > -2+1) |
| 2 | -3 | -2 | 1 | Extend (1 + -3 = -2) |
| 3 | 4 | 4 | 4 | Start new subarray (4 > -2+4) |
| 4 | -1 | 3 | 4 | Extend (4 + -1 = 3) |
| 5 | 2 | 5 | 5 | Extend (3 + 2 = 5) |
| 6 | 1 | 6 | 6 | Extend (5 + 1 = 6) |
| 7 | -5 | 1 | 6 | Extend (6 + -5 = 1) |
| 8 | 4 | 5 | 6 | Extend (1 + 4 = 5) |

Final answer: 6 (subarray [4, -1, 2, 1])

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the array exactly once
- **Space Complexity O(1)**: We only use two variables regardless of input size
- This is optimal because we must examine every element at least once to determine the maximum sum, giving us a lower bound of O(n)
- Kadane's Algorithm is a classic example of dynamic programming with space optimization

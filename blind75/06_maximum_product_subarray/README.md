# Maximum Product Subarray

## Problem Description

Given an integer array `nums`, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

A subarray is a contiguous non-empty sequence of elements within an array.

## Examples

**Example 1:**
```
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
```

**Example 2:**
```
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
```

**Example 3:**
```
Input: nums = [-2,3,-4]
Output: 24
Explanation: [-2,3,-4] has the largest product 24.
```

## Constraints

- `1 <= nums.length <= 2 * 10^4`
- `-10 <= nums[i] <= 10`
- The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming (Modified Kadane's) | Array | O(n) | O(1) |

### Approach: Track Both Max and Min Products

This problem is similar to Maximum Subarray but with a crucial difference: negative numbers can flip the sign of the product. A small negative product can become a large positive product when multiplied by another negative number.

The key insight is to track both the maximum AND minimum products ending at each position because:
- The maximum product can become minimum when multiplied by a negative number
- The minimum product can become maximum when multiplied by a negative number

At each position, we calculate three candidates:
1. The current number itself (start a new subarray)
2. Current number times the maximum product ending at previous position
3. Current number times the minimum product ending at previous position

**Key Insight:** Because multiplication by a negative number swaps max and min, we must track both to handle all cases correctly.

### Algorithm Steps

1. Initialize `max_so_far`, `max_ending_here`, and `min_ending_here` to the first element
2. For each element from index 1 onwards:
   - Calculate three candidates:
     - `num` (start fresh)
     - `num * max_ending_here` (extend max product)
     - `num * min_ending_here` (min becomes max if num is negative)
   - Update `max_ending_here` to the maximum of the three candidates
   - Update `min_ending_here` to the minimum of the three candidates
   - Update `max_so_far` if `max_ending_here` is greater
3. Return `max_so_far`

### Example Walkthrough

For `nums = [2, 3, -2, 4]`:

| Index | num | max_ending | min_ending | max_so_far | Explanation |
|-------|-----|------------|------------|------------|-------------|
| 0 | 2 | 2 | 2 | 2 | Initialize with first element |
| 1 | 3 | 6 | 3 | 6 | max(3, 3*2, 3*2) = 6 |
| 2 | -2 | 3 | -12 | 6 | max(-2, -2*6, -2*3) = 3; min = -12 |
| 3 | 4 | 12 | -48 | 12 | max(4, 4*3, 4*-12) = 12 |

Wait, let me recalculate for the actual maximum:

For `nums = [2, 3, -2, 4]`:

| Index | num | Candidates | max_ending | min_ending | max_so_far |
|-------|-----|------------|------------|------------|------------|
| 0 | 2 | - | 2 | 2 | 2 |
| 1 | 3 | 3, 6, 6 | 6 | 3 | 6 |
| 2 | -2 | -2, -12, -6 | -2 | -12 | 6 |
| 3 | 4 | 4, -8, -48 | 4 | -48 | 6 |

The subarray [2, 3] gives the maximum product of 6.

For `nums = [-2, 3, -4]`:

| Index | num | max_ending | min_ending | max_so_far |
|-------|-----|------------|------------|------------|
| 0 | -2 | -2 | -2 | -2 |
| 1 | 3 | 3 | -6 | 3 |
| 2 | -4 | 24 | -12 | 24 |

At index 2: -4 * min_ending(-6) = 24, which is the maximum!

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the array exactly once
- **Space Complexity O(1)**: We only use three variables regardless of input size
- This is optimal because we must examine every element at least once, giving us a lower bound of O(n)
- The algorithm elegantly handles negative numbers, zeros, and sign flips by tracking both maximum and minimum products

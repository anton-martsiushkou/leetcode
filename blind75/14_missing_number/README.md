# Missing Number

## Problem Description

Given an array `nums` containing `n` distinct numbers in the range `[0, n]`, return the only number in the range that is missing from the array.

## Examples

**Example 1:**
```
Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3].
2 is the missing number in the range since it does not appear in nums.
```

**Example 2:**
```
Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2].
2 is the missing number in the range since it does not appear in nums.
```

**Example 3:**
```
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9].
8 is the missing number in the range since it does not appear in nums.
```

## Constraints

- `n == nums.length`
- `1 <= n <= 10^4`
- `0 <= nums[i] <= n`
- All the numbers of `nums` are unique.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| XOR Bit Manipulation | N/A | O(n) | O(1) |

### Approach: XOR Bit Manipulation

The optimal solution uses the XOR operation's special properties. XOR has two key characteristics:
1. `a ^ a = 0` (any number XORed with itself is 0)
2. `a ^ 0 = a` (any number XORed with 0 is itself)
3. XOR is commutative and associative

**Key Insight:** If we XOR all numbers from 0 to n and XOR all numbers in the array, every number that appears will cancel out (since `a ^ a = 0`), leaving only the missing number.

### Algorithm Steps

1. Initialize `result = n` (start with the length of the array)
2. For each index `i` from 0 to n-1:
   - XOR `result` with `i` (to include all indices)
   - XOR `result` with `nums[i]` (to include all array values)
3. Return `result`

This works because we're effectively computing: `0 ^ 1 ^ 2 ^ ... ^ n ^ nums[0] ^ nums[1] ^ ... ^ nums[n-1]`

### Example Walkthrough

For `nums = [3, 0, 1]`, n = 3:

```
result = 3

i=0: result = 3 ^ 0 ^ 3 = 0
i=1: result = 0 ^ 1 ^ 0 = 1
i=2: result = 1 ^ 2 ^ 1 = 2

Return 2
```

Breaking it down:
- All numbers from 0 to 3: `0 ^ 1 ^ 2 ^ 3`
- Numbers in array: `3 ^ 0 ^ 1`
- Combined: `0 ^ 1 ^ 2 ^ 3 ^ 3 ^ 0 ^ 1`
- Rearranged: `(0 ^ 0) ^ (1 ^ 1) ^ (3 ^ 3) ^ 2 = 0 ^ 0 ^ 0 ^ 2 = 2`

### Why This is Optimal

- **Time Complexity O(n)**: We iterate through the array once
- **Space Complexity O(1)**: We only use a single variable for the result
- This is optimal because we must examine each element at least once

### Alternative Approaches

1. **Math (Sum Formula)**:
   - Expected sum: `n * (n + 1) / 2`
   - Actual sum: sum of all elements
   - Missing = Expected - Actual
   - Time: O(n), Space: O(1)
   - **Caveat**: Can overflow for large values of n

2. **Hash Set**:
   - Add all numbers to a set, then check 0 to n
   - Time: O(n), Space: O(n)

3. **Sorting**:
   - Sort and check for gap
   - Time: O(n log n), Space: O(1) or O(n) depending on sort

The XOR approach is preferred because it's O(1) space, O(n) time, and has no overflow concerns.

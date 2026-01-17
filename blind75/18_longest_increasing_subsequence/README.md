# Longest Increasing Subsequence

## Problem Description

Given an integer array `nums`, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements.

## Examples

**Example 1:**
```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

**Example 2:**
```
Input: nums = [0,1,0,3,2,3]
Output: 4
Explanation: The longest increasing subsequence is [0,1,2,3].
```

**Example 3:**
```
Input: nums = [7,7,7,7,7,7,7]
Output: 1
Explanation: The longest increasing subsequence is [7] (any single element).
```

## Constraints

- `1 <= nums.length <= 2500`
- `-10^4 <= nums[i] <= 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming with Binary Search | Array | O(n log n) | O(n) |

### Approach: Dynamic Programming with Binary Search (Patience Sorting)

The optimal solution uses a greedy approach with binary search. We maintain an array `tails` where `tails[i]` is the smallest tail element of all increasing subsequences of length `i+1`. For each number, we use binary search to find where it should be placed in this array.

**Key Insight:** For each position in `tails`, we want to keep the smallest possible value. This maximizes our chances of extending the subsequence with future elements. If a number is larger than all elements in `tails`, it extends the longest subsequence. Otherwise, it replaces the smallest element in `tails` that is >= to it.

### Algorithm Steps

1. **Initialize:**
   - Create empty array `tails`

2. **Process Each Number:**
   - For each `num` in `nums`:
     - Use binary search to find the leftmost position in `tails` where `tails[pos] >= num`
     - If no such position exists (num is larger than all), append `num` to `tails`
     - Otherwise, replace `tails[pos]` with `num`

3. **Return Result:**
   - Return `len(tails)` (the length of the longest increasing subsequence)

### Example Walkthrough

For `nums = [10, 9, 2, 5, 3, 7, 101, 18]`:

1. **num = 10:** tails = [10]

2. **num = 9:** Binary search finds position 0, replace
   - tails = [9]

3. **num = 2:** Binary search finds position 0, replace
   - tails = [2]

4. **num = 5:** Larger than all, append
   - tails = [2, 5]

5. **num = 3:** Binary search finds position 1, replace
   - tails = [2, 3]

6. **num = 7:** Larger than all, append
   - tails = [2, 3, 7]

7. **num = 101:** Larger than all, append
   - tails = [2, 3, 7, 101]

8. **num = 18:** Binary search finds position 3, replace
   - tails = [2, 3, 7, 18]

9. **Result:** Length = 4

Note: The `tails` array doesn't represent the actual subsequence, just the length. The actual LIS is [2, 3, 7, 101] or [2, 3, 7, 18].

### Why This is Optimal

- **Time Complexity O(n log n)**: For each of n elements, we perform a binary search which takes O(log n)
- **Space Complexity O(n)**: The `tails` array can grow to size n in the worst case
- This is optimal for comparison-based algorithms (proven lower bound for this problem)

### Alternative Approaches

1. **Dynamic Programming O(n²)**: O(n²) time, O(n) space - For each element, check all previous elements. Works but slower.
2. **Recursive with Memoization**: O(n²) time, O(n²) space - More space overhead than iterative DP

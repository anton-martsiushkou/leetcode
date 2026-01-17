# Longest Consecutive Sequence

## Problem Description

Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in `O(n)` time.

## Examples

**Example 1:**
```
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
```

**Example 2:**
```
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
```

## Constraints

- `0 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Hash Set (Intelligent Sequence Building) | Hash Set | O(n) | O(n) |
| Sorting | Sort | O(n log n) | O(1) or O(n) |

### Approach: Hash Set with Intelligent Sequence Building

The key insight is to use a hash set for O(1) lookups and only start counting sequences from their beginning. We identify the start of a sequence by checking if `num - 1` exists in the set.

**Key Insight:** Only count sequences starting from their first element. A number is the start of a sequence if `num - 1` is not in the set. This ensures each number is visited at most twice (once in outer loop, once when building a sequence), giving us O(n) time.

### Algorithm Steps

1. Convert the array to a hash set for O(1) lookups
2. For each number in the set:
   - Check if it's the start of a sequence (i.e., `num - 1` not in set)
   - If yes, count consecutive numbers: `num`, `num+1`, `num+2`, ...
   - Keep track of the maximum length found
3. Return the maximum length

### Why Not Sort?

Sorting would give us O(n log n) time, which violates the problem constraint. The hash set approach achieves O(n) by:
- Building the set: O(n)
- Iterating through numbers: O(n)
- Counting sequences: Each number is counted at most once across all sequences, so total is O(n)

### Example Walkthrough

For `nums = [100, 4, 200, 1, 3, 2]`:

1. Create set: `{100, 4, 200, 1, 3, 2}`

2. Check each number:
   - **100**: Is 99 in set? No → Start of sequence
     - Count: 100 (length 1)
   - **4**: Is 3 in set? Yes → Not a start, skip
   - **200**: Is 199 in set? No → Start of sequence
     - Count: 200 (length 1)
   - **1**: Is 0 in set? No → Start of sequence
     - Count: 1, 2, 3, 4 (length 4) ← Maximum!
   - **3**: Is 2 in set? Yes → Not a start, skip
   - **2**: Is 1 in set? Yes → Not a start, skip

3. Maximum length = 4

### Why This is Optimal

- **Time Complexity O(n)**:
  - Building set: O(n)
  - Outer loop: O(n)
  - Inner while loop: Across all iterations, each number is visited at most once (only when we start from the beginning of a sequence)
  - Total: O(n) + O(n) + O(n) = O(n)

- **Space Complexity O(n)**: Hash set stores all unique numbers

- This is optimal because we must examine each number at least once, giving us a lower bound of O(n)

### Alternative Approach (Not Optimal)

**Sorting:** O(n log n) time, O(1) or O(n) space
- Sort the array
- Iterate through and count consecutive sequences
- This violates the O(n) time requirement

# Contains Duplicate

## Problem Description

Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct.

## Examples

**Example 1:**
```
Input: nums = [1,2,3,1]
Output: true
```

**Example 2:**
```
Input: nums = [1,2,3,4]
Output: false
```

**Example 3:**
```
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
```

## Constraints

- `1 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Hash Set | Hash Set | O(n) | O(n) |

### Approach: Hash Set

The optimal solution uses a hash set to track elements we've seen. As we iterate through the array, we check if the current element exists in the set. If it does, we found a duplicate. If not, we add it to the set.

**Key Insight:** Hash set provides O(1) lookup time, allowing us to check for duplicates in a single pass.

### Algorithm Steps

1. Create an empty hash set
2. For each number in the array:
   - If the number exists in the set, return `true` (duplicate found)
   - Otherwise, add the number to the set
3. If we finish iterating without finding duplicates, return `false`

### Example Walkthrough

For `nums = [1, 2, 3, 1]`:

1. **num=1**: set is empty, add 1 → set = {1}
2. **num=2**: 2 not in set, add 2 → set = {1, 2}
3. **num=3**: 3 not in set, add 3 → set = {1, 2, 3}
4. **num=1**: 1 is in set! → return `true`

### Why This is Optimal

- **Time Complexity O(n)**: Single pass through array with O(1) set operations
- **Space Complexity O(n)**: In worst case (all distinct), we store all n elements
- This is optimal because we need to examine all elements to determine if duplicates exist

### Alternative Approaches

1. **Sorting**: O(n log n) time, O(1) space - Slower but saves space
2. **Brute Force**: O(n²) time, O(1) space - Too slow for large inputs

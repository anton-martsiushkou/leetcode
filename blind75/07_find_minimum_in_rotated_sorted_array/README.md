# Find Minimum in Rotated Sorted Array

## Problem Description

Suppose an array of length `n` sorted in ascending order is rotated between `1` and `n` times. For example, the array `nums = [0,1,2,4,5,6,7]` might become:
- `[4,5,6,7,0,1,2]` if it was rotated 4 times.
- `[0,1,2,4,5,6,7]` if it was rotated 7 times.

Notice that rotating an array `[a[0], a[1], a[2], ..., a[n-1]]` 1 time results in the array `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]`.

Given the sorted rotated array `nums` of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

## Examples

**Example 1:**
```
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
```

**Example 2:**
```
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
```

**Example 3:**
```
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times (or 0 times).
```

## Constraints

- `n == nums.length`
- `1 <= n <= 5000`
- `-5000 <= nums[i] <= 5000`
- All the integers of `nums` are unique.
- `nums` is sorted and rotated between `1` and `n` times.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Modified Binary Search | Array | O(log n) | O(1) |

### Approach: Modified Binary Search

The key insight is that in a rotated sorted array, one half is always sorted. We can use this property with binary search to find the minimum element.

When comparing the middle element with the rightmost element:
- If `nums[mid] > nums[right]`: The minimum is in the right half (the rotation point is on the right)
- If `nums[mid] < nums[right]`: The minimum is in the left half including mid (the array is sorted from mid to right)
- If `nums[mid] == nums[right]`: Not possible in this problem since all elements are unique

**Key Insight:** The minimum element is the only element where `nums[i] < nums[i-1]`. In a rotated sorted array, all elements to the left of the minimum are greater than all elements to the right.

### Algorithm Steps

1. Initialize `left = 0` and `right = len(nums) - 1`
2. While `left < right`:
   - Calculate `mid = left + (right - left) / 2`
   - If `nums[mid] > nums[right]`:
     - The minimum is in the right half
     - Set `left = mid + 1`
   - Else:
     - The minimum is in the left half (including mid)
     - Set `right = mid`
3. Return `nums[left]` (left and right will converge to the minimum)

### Example Walkthrough

For `nums = [4, 5, 6, 7, 0, 1, 2]`:

| Iteration | left | right | mid | nums[mid] | nums[right] | Action |
|-----------|------|-------|-----|-----------|-------------|--------|
| 1 | 0 | 6 | 3 | 7 | 2 | 7 > 2, left = 4 |
| 2 | 4 | 6 | 5 | 1 | 2 | 1 < 2, right = 5 |
| 3 | 4 | 5 | 4 | 0 | 1 | 0 < 1, right = 4 |
| - | 4 | 4 | - | - | - | left == right, return nums[4] = 0 |

For `nums = [11, 13, 15, 17]` (no rotation or full rotation):

| Iteration | left | right | mid | nums[mid] | nums[right] | Action |
|-----------|------|-------|-----|-----------|-------------|--------|
| 1 | 0 | 3 | 1 | 13 | 17 | 13 < 17, right = 1 |
| 2 | 0 | 1 | 0 | 11 | 13 | 11 < 13, right = 0 |
| - | 0 | 0 | - | - | - | left == right, return nums[0] = 11 |

### Why This is Optimal

- **Time Complexity O(log n)**: Binary search halves the search space at each step
- **Space Complexity O(1)**: We only use a constant amount of extra space
- This is optimal because the problem explicitly requires O(log n) time, ruling out linear search
- Binary search is the standard approach for logarithmic time complexity on sorted or partially sorted data

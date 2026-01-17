# Search in Rotated Sorted Array

## Problem Description

There is an integer array `nums` sorted in ascending order (with distinct values).

Prior to being passed to your function, `nums` is possibly rotated at an unknown pivot index `k` (`1 <= k < nums.length`) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (0-indexed). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index `3` and become `[4,5,6,7,0,1,2]`.

Given the array `nums` after the possible rotation and an integer `target`, return the index of `target` if it is in `nums`, or `-1` if it is not in `nums`.

You must write an algorithm with O(log n) runtime complexity.

## Examples

**Example 1:**
```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

**Example 2:**
```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

**Example 3:**
```
Input: nums = [1], target = 0
Output: -1
```

## Constraints

- `1 <= nums.length <= 5000`
- `-10^4 <= nums[i] <= 10^4`
- All values of `nums` are unique.
- `nums` is an ascending array that is possibly rotated.
- `-10^4 <= target <= 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Modified Binary Search | Array | O(log n) | O(1) |

### Approach: Modified Binary Search

The key insight is that in a rotated sorted array, at least one half of the array is always sorted. We can determine which half is sorted by comparing the middle element with the left or right boundary.

Once we know which half is sorted, we can:
1. Check if the target is within the sorted half's range
2. If yes, search in that half
3. If no, search in the other half

**Key Insight:** At any point in the binary search, at least one of the two halves [left...mid] or [mid...right] must be sorted. We can use this to determine which side to search.

### Algorithm Steps

1. Initialize `left = 0` and `right = len(nums) - 1`
2. While `left <= right`:
   - Calculate `mid = left + (right - left) / 2`
   - If `nums[mid] == target`, return `mid`
   - Determine which half is sorted:
     - **If left half is sorted** (`nums[left] <= nums[mid]`):
       - If target is in range `[nums[left], nums[mid])`, search left: `right = mid - 1`
       - Otherwise, search right: `left = mid + 1`
     - **If right half is sorted** (`nums[mid] <= nums[right]`):
       - If target is in range `(nums[mid], nums[right]]`, search right: `left = mid + 1`
       - Otherwise, search left: `right = mid - 1`
3. If loop ends without finding target, return `-1`

### Example Walkthrough

For `nums = [4, 5, 6, 7, 0, 1, 2]`, `target = 0`:

| Iteration | left | right | mid | nums[mid] | Left sorted? | Action |
|-----------|------|-------|-----|-----------|--------------|--------|
| 1 | 0 | 6 | 3 | 7 | Yes [4..7] | 0 not in [4,7), search right, left=4 |
| 2 | 4 | 6 | 5 | 1 | No, right sorted [1..2] | 0 not in (1,2], search left, right=4 |
| 3 | 4 | 4 | 4 | 0 | - | Found! Return 4 |

For `nums = [4, 5, 6, 7, 0, 1, 2]`, `target = 3`:

| Iteration | left | right | mid | nums[mid] | Left sorted? | Action |
|-----------|------|-------|-----|-----------|--------------|--------|
| 1 | 0 | 6 | 3 | 7 | Yes [4..7] | 3 not in [4,7), search right, left=4 |
| 2 | 4 | 6 | 5 | 1 | No, right sorted [1..2] | 3 not in (1,2], search left, right=4 |
| 3 | 4 | 4 | 4 | 0 | - | 0 ≠ 3, left=5 |
| 4 | 5 | 4 | - | - | - | left > right, return -1 |

### Why This is Optimal

- **Time Complexity O(log n)**: Binary search halves the search space at each step
- **Space Complexity O(1)**: We only use a constant amount of extra space
- This is optimal because the problem explicitly requires O(log n) time
- The algorithm cleverly uses the property that one half is always sorted to make the correct binary search decision

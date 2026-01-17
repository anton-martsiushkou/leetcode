# 3Sum

## Problem Description

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

## Examples

**Example 1:**
```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
```

**Example 2:**
```
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
```

**Example 3:**
```
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
```

## Constraints

- `3 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Sort + Two Pointers | Array | O(n²) | O(1)* |

*O(1) if we don't count the space for sorting and output. Sorting takes O(log n) to O(n) depending on implementation.

### Approach: Sort + Two Pointers

The optimal solution combines sorting with the two-pointer technique:

1. Sort the array to enable duplicate detection and two-pointer approach
2. For each element (as the first element of the triplet), use two pointers to find pairs that sum to the negative of that element
3. Skip duplicates to avoid duplicate triplets

**Key Insight:** After sorting, we can fix one element and use two pointers to find pairs in the remaining array that sum to the target. This reduces the 3Sum problem to a 2Sum problem for each fixed element.

### Algorithm Steps

1. Sort the array in ascending order
2. For each index `i` from `0` to `n-3`:
   - Skip duplicates: if `i > 0` and `nums[i] == nums[i-1]`, continue
   - Set `target = -nums[i]`
   - Initialize two pointers: `left = i + 1`, `right = n - 1`
   - While `left < right`:
     - Calculate `sum = nums[left] + nums[right]`
     - If `sum == target`:
       - Add triplet `[nums[i], nums[left], nums[right]]` to result
       - Skip duplicates: increment `left` while `nums[left] == nums[left-1]`
       - Skip duplicates: decrement `right` while `nums[right] == nums[right+1]`
       - Move both pointers: `left++`, `right--`
     - If `sum < target`: increment `left`
     - If `sum > target`: decrement `right`
3. Return result

### Example Walkthrough

For `nums = [-1, 0, 1, 2, -1, -4]`:

**After sorting:** `[-4, -1, -1, 0, 1, 2]`

**i = 0, nums[i] = -4, target = 4:**
- left = 1 (-1), right = 5 (2): sum = 1, too small
- left = 2 (-1), right = 5 (2): sum = 1, too small
- left = 3 (0), right = 5 (2): sum = 2, too small
- left = 4 (1), right = 5 (2): sum = 3, too small
- left = 5, left >= right, done

**i = 1, nums[i] = -1, target = 1:**
- left = 2 (-1), right = 5 (2): sum = 1, found! Add [-1, -1, 2]
- Skip duplicate at left (still -1)
- left = 3 (0), right = 4 (1): sum = 1, found! Add [-1, 0, 1]
- left = 4, right = 4, done

**i = 2, nums[i] = -1:** Skip (duplicate of i=1)

**i = 3, nums[i] = 0, target = 0:**
- left = 4 (1), right = 5 (2): sum = 3, too large
- left = 4, right = 4, done

Result: `[[-1, -1, 2], [-1, 0, 1]]`

### Why This is Optimal

- **Time Complexity O(n²)**: Sorting takes O(n log n), then we iterate through n elements, each with O(n) two-pointer scan, giving O(n²) which dominates
- **Space Complexity O(1)**: We only use constant extra space (not counting the output array and sorting space)
- This is optimal because we need to examine all possible triplets, and O(n²) is the best we can do without using additional complex data structures
- The two-pointer technique after sorting is more efficient than using a hash map for each element (which would also be O(n²) but with worse constants)

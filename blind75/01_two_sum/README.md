# Two Sum

## Problem Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## Examples

**Example 1:**
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

## Constraints

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- Only one valid answer exists.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Hash Map (One-pass) | Hash Map/Dictionary | O(n) | O(n) |

### Approach: Hash Map (One-pass)

The optimal solution uses a hash map to store the values we've seen so far along with their indices. As we iterate through the array, for each element, we calculate its complement (target - current element) and check if this complement exists in our hash map.

**Key Insight:** Instead of checking every pair of numbers (which would be O(n²)), we can use a hash map to check in O(1) time whether the complement of the current number exists.

### Algorithm Steps

1. Create an empty hash map to store `{value: index}` pairs
2. Iterate through the array with index `i` and value `num`:
   - Calculate `complement = target - num`
   - Check if `complement` exists in the hash map
   - If yes: return `[hash_map[complement], i]`
   - If no: add the current number to hash map: `hash_map[num] = i`
3. If no solution is found during iteration, the input is invalid (but problem guarantees exactly one solution exists)

### Example Walkthrough

For `nums = [2, 7, 11, 15]` and `target = 9`:

1. **i=0, num=2**: complement = 9-2 = 7
   - Hash map is empty, 7 not found
   - Add to map: `{2: 0}`

2. **i=1, num=7**: complement = 9-7 = 2
   - Check map: 2 exists at index 0
   - Return `[0, 1]`

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the array only once, and hash map operations (insert/lookup) are O(1) on average
- **Space Complexity O(n)**: In the worst case, we might need to store all n elements in the hash map before finding the solution
- This is optimal because we must examine at least every element once to find the solution, giving us a lower bound of O(n)

### Alternative Approaches (Not Optimal)

1. **Brute Force (Two Nested Loops)**: O(n²) time, O(1) space - Too slow for large inputs
2. **Sort + Two Pointers**: O(n log n) time, O(n) space - Requires sorting which changes indices, so we'd need to store original indices first
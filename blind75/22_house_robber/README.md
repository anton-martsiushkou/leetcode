# House Robber

## Problem Description

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and **it will automatically contact the police if two adjacent houses were broken into on the same night**.

Given an integer array `nums` representing the amount of money of each house, return the maximum amount of money you can rob tonight **without alerting the police**.

## Examples

**Example 1:**
```
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
```

**Example 2:**
```
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
```

## Constraints

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 400`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming | Array/Variables | O(n) | O(1) |

### Approach: Dynamic Programming

The optimal solution uses dynamic programming to build up the maximum amount that can be robbed at each house. At each house, we have two choices:
1. Rob the current house + max amount from houses up to i-2
2. Don't rob the current house, take max amount from houses up to i-1

**Key Insight:** For each house i, the maximum amount we can rob is: `max(rob[i-1], rob[i-2] + nums[i])`. This means we either skip the current house (and take the max from the previous house) or rob it (and add it to the max from two houses ago).

### Algorithm Steps

1. **Base Cases:**
   - If there are no houses, return 0
   - If there's only one house, rob it
   - If there are two houses, rob the one with more money

2. **DP Recurrence:**
   - For each house i (starting from index 2):
     - Calculate max money if we rob this house: `prev2 + nums[i]`
     - Calculate max money if we skip this house: `prev1`
     - Current max is the maximum of these two choices
     - Update prev2 and prev1 for next iteration

3. **Return** the maximum amount after considering all houses

### Example Walkthrough

For `nums = [2, 7, 9, 3, 1]`:

1. **i=0**: rob = 2
   - Only one house, rob it

2. **i=1**: rob = max(2, 7) = 7
   - Two houses, rob the one with more money

3. **i=2**: rob = max(7, 2+9) = 11
   - Skip house 2 (total=7) or rob house 2 with house 0 (total=11)

4. **i=3**: rob = max(11, 7+3) = 11
   - Skip house 3 (total=11) or rob house 3 with max from house 1 (total=10)

5. **i=4**: rob = max(11, 11+1) = 12
   - Skip house 4 (total=11) or rob house 4 with max from house 2 (total=12)

**Result:** 12 (rob houses at indices 0, 2, 4)

### Why This is Optimal

- **Time Complexity O(n)**: We iterate through the array once, making a constant-time decision at each house
- **Space Complexity O(1)**: We only need two variables to track the previous two states, rather than storing the entire DP array
- This is optimal because we must examine every house at least once to make an informed decision, giving us a lower bound of O(n)

### Alternative Approaches

1. **Recursive (Brute Force)**: O(2^n) time, O(n) space - Too slow, explores all possible combinations
2. **DP with Array**: O(n) time, O(n) space - Same time complexity but uses more space than necessary
3. **Memoization**: O(n) time, O(n) space - Top-down approach with similar complexity

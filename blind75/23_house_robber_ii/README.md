# House Robber II

## Problem Description

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are **arranged in a circle**. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have security systems connected, and **it will automatically contact the police if two adjacent houses were broken into on the same night**.

Given an integer array `nums` representing the amount of money of each house, return the maximum amount of money you can rob tonight **without alerting the police**.

## Examples

**Example 1:**
```
Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent.
```

**Example 2:**
```
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
```

**Example 3:**
```
Input: nums = [1,2,3]
Output: 3
```

## Constraints

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 1000`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming (Two Pass) | Variables | O(n) | O(1) |

### Approach: Dynamic Programming (Two Pass)

This problem is an extension of House Robber I with a circular constraint. The key insight is that since houses are arranged in a circle, we cannot rob both the first and last house. We can break this into two subproblems:

1. Rob houses from index 0 to n-2 (exclude last house)
2. Rob houses from index 1 to n-1 (exclude first house)

The answer is the maximum of these two scenarios.

**Key Insight:** The circular constraint means we have two mutually exclusive scenarios: either we can rob the first house (and skip the last), or we can rob the last house (and skip the first). We solve both and take the maximum.

### Algorithm Steps

1. **Edge Cases:**
   - If there's only one house, rob it
   - If there are two houses, rob the one with more money

2. **Two Scenarios:**
   - Scenario A: Rob houses [0...n-2] using House Robber I logic
   - Scenario B: Rob houses [1...n-1] using House Robber I logic

3. **Return** the maximum of both scenarios

4. **House Robber I Helper:**
   - Use the same DP approach as House Robber I
   - Track prev2 (max up to i-2) and prev1 (max up to i-1)
   - For each house: current = max(prev1, prev2 + nums[i])

### Example Walkthrough

For `nums = [2, 3, 2]`:

**Scenario A: Rob houses [0...1] = [2, 3]**
- prev2 = 2, prev1 = max(2, 3) = 3
- Result: 3

**Scenario B: Rob houses [1...2] = [3, 2]**
- prev2 = 3, prev1 = max(3, 2) = 3
- Result: 3

**Final Result:** max(3, 3) = 3

For `nums = [1, 2, 3, 1]`:

**Scenario A: Rob houses [0...2] = [1, 2, 3]**
- i=0: rob = 1
- i=1: rob = max(1, 2) = 2
- i=2: rob = max(2, 1+3) = 4
- Result: 4

**Scenario B: Rob houses [1...3] = [2, 3, 1]**
- i=1: rob = 2
- i=2: rob = max(2, 3) = 3
- i=3: rob = max(3, 2+1) = 3
- Result: 3

**Final Result:** max(4, 3) = 4

### Why This is Optimal

- **Time Complexity O(n)**: We make two passes through the array (one for each scenario), each taking O(n) time
- **Space Complexity O(1)**: We only use a constant amount of extra space for variables
- This is optimal because we must examine every house to make an informed decision, giving us a lower bound of O(n)

### Alternative Approaches

1. **Recursive (Brute Force)**: O(2^n) time - Too slow, explores all possible combinations
2. **DP with Arrays**: O(n) time, O(n) space - Same time complexity but uses more space
3. **Single Pass with State Tracking**: More complex implementation with same complexity

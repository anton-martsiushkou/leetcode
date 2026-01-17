# Jump Game

## Problem Description

You are given an integer array `nums`. You are initially positioned at the array's **first index**, and each element in the array represents your maximum jump length at that position.

Return `true` if you can reach the last index, or `false` otherwise.

## Examples

**Example 1:**
```
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2:**
```
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

## Constraints

- `1 <= nums.length <= 10^4`
- `0 <= nums[i] <= 10^5`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Greedy (Max Reach) | Variable | O(n) | O(1) |

### Approach: Greedy Algorithm

The optimal solution uses a greedy approach by tracking the maximum index we can reach as we iterate through the array. At each position, we update the maximum reachable index and check if we can reach the current position.

**Key Insight:** We don't need to track all possible paths or use dynamic programming. We only need to track the farthest position we can reach. If at any point our current position exceeds the maximum reachable position, we know we can't reach the end.

### Algorithm Steps

1. **Initialize:**
   - maxReach = 0 (farthest index we can reach)

2. **Iterate through array:**
   - For each index i from 0 to n-1:
     - If i > maxReach: return false (can't reach current position)
     - Update maxReach = max(maxReach, i + nums[i])
     - If maxReach >= n-1: return true (can reach the end)

3. **Return** true (if we complete the loop, we can reach the end)

### Example Walkthrough

For `nums = [2, 3, 1, 1, 4]`:

1. **i=0, nums[0]=2:**
   - maxReach = max(0, 0+2) = 2
   - Can reach indices 0, 1, 2

2. **i=1, nums[1]=3:**
   - i (1) <= maxReach (2), continue
   - maxReach = max(2, 1+3) = 4
   - Can now reach indices 0, 1, 2, 3, 4
   - maxReach (4) >= last index (4), return true

**Result:** true

For `nums = [3, 2, 1, 0, 4]`:

1. **i=0, nums[0]=3:**
   - maxReach = max(0, 0+3) = 3
   - Can reach indices 0, 1, 2, 3

2. **i=1, nums[1]=2:**
   - i (1) <= maxReach (3), continue
   - maxReach = max(3, 1+2) = 3
   - Still can only reach up to index 3

3. **i=2, nums[2]=1:**
   - i (2) <= maxReach (3), continue
   - maxReach = max(3, 2+1) = 3
   - Still can only reach up to index 3

4. **i=3, nums[3]=0:**
   - i (3) <= maxReach (3), continue
   - maxReach = max(3, 3+0) = 3
   - Stuck at index 3, can't go further

5. **i=4:**
   - i (4) > maxReach (3), return false

**Result:** false

### Why This is Optimal

- **Time Complexity O(n)**: We iterate through the array once, with constant-time operations at each step
- **Space Complexity O(1)**: We only use a single variable to track the maximum reachable index
- This is optimal because we must examine each element to determine if we can reach the end, giving us a lower bound of O(n)
- The greedy approach is correct because if we can reach a position, we can reach all positions before it

### Alternative Approaches

1. **Dynamic Programming**: O(n^2) time, O(n) space - Overkill, checks all possible jumps from each position
2. **BFS/DFS**: O(n^2) time, O(n) space - Treats it as a graph problem, but unnecessary
3. **Backward Iteration**: O(n) time, O(1) space - Start from the end and work backwards, same complexity

# Climbing Stairs

## Problem Description

You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

## Examples

**Example 1:**
```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

**Example 2:**
```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

**Example 3:**
```
Input: n = 4
Output: 5
Explanation: There are five ways to climb to the top.
1. 1 + 1 + 1 + 1
2. 1 + 1 + 2
3. 1 + 2 + 1
4. 2 + 1 + 1
5. 2 + 2
```

## Constraints

- `1 <= n <= 45`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming (Optimized) | Two Variables | O(n) | O(1) |

### Approach: Dynamic Programming with Space Optimization

This problem is essentially calculating Fibonacci numbers. The key insight is that to reach step `n`, you can either come from step `n-1` (with a 1-step) or step `n-2` (with a 2-step). Therefore, the number of ways to reach step `n` is the sum of ways to reach `n-1` and `n-2`.

**Key Insight:** This follows the Fibonacci sequence pattern: `dp[n] = dp[n-1] + dp[n-2]`, where `dp[i]` represents the number of distinct ways to reach step `i`.

### Algorithm Steps

1. **Base Cases:**
   - If `n = 1`: return 1 (only one way: 1 step)
   - If `n = 2`: return 2 (two ways: 1+1 or 2)

2. **Iterative Calculation:**
   - Initialize `prev2 = 1` (ways to reach step 1)
   - Initialize `prev1 = 2` (ways to reach step 2)
   - For each step from 3 to n:
     - Calculate `current = prev1 + prev2`
     - Update `prev2 = prev1`
     - Update `prev1 = current`
   - Return `prev1`

### Example Walkthrough

For `n = 5`:

1. **Base:** prev2 = 1 (step 1), prev1 = 2 (step 2)

2. **Step 3:** current = 1 + 2 = 3
   - prev2 = 2, prev1 = 3

3. **Step 4:** current = 2 + 3 = 5
   - prev2 = 3, prev1 = 5

4. **Step 5:** current = 3 + 5 = 8
   - prev2 = 5, prev1 = 8

5. **Result:** 8 distinct ways

### Why This is Optimal

- **Time Complexity O(n)**: We iterate through steps 3 to n exactly once
- **Space Complexity O(1)**: We only use two variables instead of an array, which is the most space-efficient approach
- This is optimal because we must calculate each Fibonacci number at least once to get the final result

### Alternative Approaches (Not Optimal)

1. **Recursive (Naive)**: O(2^n) time, O(n) space - Exponential time due to redundant calculations
2. **Dynamic Programming with Array**: O(n) time, O(n) space - Same time complexity but uses unnecessary space
3. **Matrix Exponentiation**: O(log n) time, O(1) space - Faster for very large n, but unnecessary complexity for n ≤ 45

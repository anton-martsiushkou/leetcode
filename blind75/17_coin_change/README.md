# Coin Change

## Problem Description

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

You may assume that you have an infinite number of each kind of coin.

## Examples

**Example 1:**
```
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
```

**Example 2:**
```
Input: coins = [2], amount = 3
Output: -1
Explanation: The amount of 3 cannot be made up with coins of 2.
```

**Example 3:**
```
Input: coins = [1], amount = 0
Output: 0
Explanation: No coins are needed to make up amount 0.
```

## Constraints

- `1 <= coins.length <= 12`
- `1 <= coins[i] <= 2^31 - 1`
- `0 <= amount <= 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming (Bottom-Up) | Array | O(amount × n) | O(amount) |

### Approach: Dynamic Programming (Bottom-Up)

The optimal solution uses dynamic programming to build up the minimum number of coins needed for each amount from 0 to the target amount. For each amount, we try using each coin and take the minimum.

**Key Insight:** To find the minimum coins for amount `i`, we check all coins: if we use a coin of value `c`, we need `1 + dp[i-c]` coins total. We take the minimum across all valid coins.

### Algorithm Steps

1. **Initialize DP Array:**
   - Create array `dp` of size `amount + 1`
   - Set `dp[0] = 0` (0 coins needed for amount 0)
   - Set all other values to `amount + 1` (represents infinity/impossible)

2. **Build DP Table:**
   - For each amount `i` from 1 to `amount`:
     - For each coin `c` in coins:
       - If `c <= i` (coin can be used):
         - `dp[i] = min(dp[i], dp[i - c] + 1)`

3. **Return Result:**
   - If `dp[amount] > amount`: return -1 (impossible)
   - Otherwise: return `dp[amount]`

### Example Walkthrough

For `coins = [1, 2, 5]` and `amount = 11`:

1. **Initialize:** dp = [0, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞]

2. **Amount 1:**
   - Coin 1: dp[1] = min(∞, dp[0] + 1) = 1
   - dp = [0, 1, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞]

3. **Amount 2:**
   - Coin 1: dp[2] = min(∞, dp[1] + 1) = 2
   - Coin 2: dp[2] = min(2, dp[0] + 1) = 1
   - dp = [0, 1, 1, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞]

4. **Amount 5:**
   - After all coins: dp[5] = 1 (one 5-coin)
   - dp = [0, 1, 1, 2, 2, 1, ∞, ∞, ∞, ∞, ∞, ∞]

5. **Amount 11:**
   - After all coins: dp[11] = 3 (5 + 5 + 1)
   - dp = [0, 1, 1, 2, 2, 1, 2, 2, 3, 3, 2, 3]

6. **Result:** 3

### Why This is Optimal

- **Time Complexity O(amount × n)**: For each of the `amount` values, we check all `n` coins
- **Space Complexity O(amount)**: We use a 1D array of size `amount + 1`
- This is optimal because we need to compute the minimum coins for each amount up to the target, and we must check each coin type

### Alternative Approaches (Not Optimal)

1. **Greedy (Always Take Largest Coin)**: Doesn't work - counterexample: coins=[1,3,4], amount=6 → greedy gives 3 (4+1+1), optimal is 2 (3+3)
2. **Recursive (Top-Down without Memoization)**: O(n^amount) time - Exponential time due to overlapping subproblems
3. **BFS**: O(amount × n) time, O(amount) space - Same complexity but more overhead than bottom-up DP

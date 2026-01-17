# Counting Bits

## Problem Description

Given an integer `n`, return an array `ans` of length `n + 1` such that for each `i` (0 <= i <= n), `ans[i]` is the number of 1's in the binary representation of `i`.

## Examples

**Example 1:**
```
Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0 (0 ones)
1 --> 1 (1 one)
2 --> 10 (1 one)
```

**Example 2:**
```
Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0 (0 ones)
1 --> 1 (1 one)
2 --> 10 (1 one)
3 --> 11 (2 ones)
4 --> 100 (1 one)
5 --> 101 (2 ones)
```

## Constraints

- `0 <= n <= 10^5`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming (Bit Manipulation) | Array | O(n) | O(n) |

### Approach: Dynamic Programming with Bit Manipulation

The optimal solution uses dynamic programming combined with the observation that the number of 1 bits in a number `i` can be derived from a previously computed value.

**Key Insight:** For any number `i`, the count of 1 bits equals the count of 1 bits in `i >> 1` (i divided by 2) plus the last bit of `i` (i.e., `i & 1`). This is because right-shifting removes the last bit, and we need to add it back if it was a 1.

Formula: `dp[i] = dp[i >> 1] + (i & 1)`

### Algorithm Steps

1. Create an array `dp` of size `n + 1` initialized to 0
2. For each number `i` from 1 to `n`:
   - `dp[i] = dp[i >> 1] + (i & 1)`
3. Return the `dp` array

### How the Formula Works

- `i >> 1`: Right shift by 1 is equivalent to dividing by 2 (removing the last bit)
- `i & 1`: Extracts the last bit (0 or 1)
- If we already know the count for `i >> 1`, we just add whether the last bit is 1

### Example Walkthrough

For `n = 5`:

| i | Binary | i >> 1 | dp[i >> 1] | i & 1 | dp[i] |
|---|--------|--------|------------|-------|-------|
| 0 | 0      | -      | -          | -     | 0     |
| 1 | 1      | 0      | 0          | 1     | 1     |
| 2 | 10     | 1      | 1          | 0     | 1     |
| 3 | 11     | 1      | 1          | 1     | 2     |
| 4 | 100    | 2      | 1          | 0     | 1     |
| 5 | 101    | 2      | 1          | 1     | 2     |

Result: `[0, 1, 1, 2, 1, 2]`

### Why This is Optimal

- **Time Complexity O(n)**: We iterate through numbers from 0 to n once, and each iteration does O(1) work
- **Space Complexity O(n)**: We need an array of size n + 1 to store the results (this is required by the problem)
- This is optimal because we must compute the result for each number from 0 to n, giving us a lower bound of O(n)

### Alternative DP Formulas

There are other valid DP approaches:

1. **Using `i & (i - 1)`**: `dp[i] = dp[i & (i - 1)] + 1`
   - This uses Brian Kernighan's algorithm idea
   - `i & (i - 1)` removes the rightmost set bit

2. **Using power of 2 offset**: Track the last power of 2 and use offset
   - More complex but also O(n)

The `i >> 1` approach is preferred for its simplicity and clarity.

### Follow-up

Can you do it in O(n) time and O(1) space (excluding the output array)? Yes, this solution achieves that.

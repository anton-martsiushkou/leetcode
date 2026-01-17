# Longest Common Subsequence

## Problem Description

Given two strings `text1` and `text2`, return the length of their longest common subsequence. If there is no common subsequence, return `0`.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

- For example, `"ace"` is a subsequence of `"abcde"`.

A common subsequence of two strings is a subsequence that is common to both strings.

## Examples

**Example 1:**
```
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
```

**Example 2:**
```
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
```

**Example 3:**
```
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no common subsequence, so the result is 0.
```

## Constraints

- `1 <= text1.length, text2.length <= 1000`
- `text1` and `text2` consist of only lowercase English characters.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming (2D) | 2D Array | O(m × n) | O(m × n) |

### Approach: Dynamic Programming (Bottom-Up)

The optimal solution uses a 2D DP table where `dp[i][j]` represents the length of the longest common subsequence of `text1[0...i-1]` and `text2[0...j-1]`. We build this table bottom-up by comparing characters.

**Key Insight:** For each pair of characters at positions `i` and `j`:
- If `text1[i-1] == text2[j-1]`: the characters match, so `dp[i][j] = dp[i-1][j-1] + 1`
- Otherwise: take the maximum of excluding either character: `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`

### Algorithm Steps

1. **Initialize DP Table:**
   - Create 2D array `dp` of size `(m+1) × (n+1)` where m = len(text1), n = len(text2)
   - Initialize all values to 0 (base case: empty string has LCS of 0 with any string)

2. **Build DP Table:**
   - For each `i` from 1 to m:
     - For each `j` from 1 to n:
       - If `text1[i-1] == text2[j-1]`:
         - `dp[i][j] = dp[i-1][j-1] + 1`
       - Else:
         - `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`

3. **Return Result:**
   - Return `dp[m][n]`

### Example Walkthrough

For `text1 = "abcde"` and `text2 = "ace"`:

```
    ""  a  c  e
""   0  0  0  0
a    0  1  1  1
b    0  1  1  1
c    0  1  2  2
d    0  1  2  2
e    0  1  2  3
```

1. **i=1, j=1 (a, a):** Match! dp[1][1] = dp[0][0] + 1 = 1

2. **i=1, j=2 (a, c):** No match. dp[1][2] = max(dp[0][2], dp[1][1]) = 1

3. **i=3, j=2 (c, c):** Match! dp[3][2] = dp[2][1] + 1 = 2

4. **i=5, j=3 (e, e):** Match! dp[5][3] = dp[4][2] + 1 = 3

5. **Result:** dp[5][3] = 3

The LCS is "ace" with length 3.

### Why This is Optimal

- **Time Complexity O(m × n)**: We fill each cell of the m × n table exactly once
- **Space Complexity O(m × n)**: We use a 2D array of size m × n
- This is optimal for the general case (can be optimized to O(min(m,n)) space using 1D array)

### Space Optimization

The space can be reduced to O(min(m, n)) by using only two rows (current and previous) instead of the full 2D table, since we only need the previous row to compute the current row.

### Alternative Approaches (Not Optimal)

1. **Recursive (No Memoization)**: O(2^(m+n)) time - Exponential time due to overlapping subproblems
2. **Recursive with Memoization**: O(m × n) time, O(m × n) space - Same complexity but more overhead

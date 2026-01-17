# Longest Palindromic Substring

## Problem Description

Given a string `s`, return the longest palindromic substring in `s`.

A string is a palindrome when it reads the same backward as forward.

## Examples

**Example 1:**
```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

**Example 2:**
```
Input: s = "cbbd"
Output: "bb"
```

**Example 3:**
```
Input: s = "a"
Output: "a"
```

**Example 4:**
```
Input: s = "ac"
Output: "a"
Explanation: Both "a" and "c" are valid answers (same length), we can return either.
```

## Constraints

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Expand Around Center | String | O(n²) | O(1) |

### Approach: Expand Around Center

The optimal solution expands around each possible center to find the longest palindrome. For each position, we check both odd-length palindromes (single character center) and even-length palindromes (two character center), keeping track of the longest palindrome found.

**Key Insight:** A palindrome mirrors around its center. By expanding from each potential center and tracking the maximum length, we can find the longest palindromic substring efficiently.

### Algorithm Steps

1. Initialize variables to track the longest palindrome found (start and end indices)
2. For each index i in the string:
   - Find the longest odd-length palindrome centered at i
   - Find the longest even-length palindrome centered between i and i+1
   - If either palindrome is longer than the current longest, update the tracking variables
3. Return the substring using the tracked start and end indices

**Helper function - expandAroundCenter(s, left, right):**
1. While left >= 0 and right < length and s[left] == s[right]:
   - Expand outward: left--, right++
2. Return the length of the palindrome found: right - left - 1

### Example Walkthrough

For `s = "babad"`:

1. **Center at index 0 ('b'):**
   - Odd: "b" (length 1)
   - Even: no match

2. **Center at index 1 ('a'):**
   - Odd: "bab" (length 3) ← Longest so far
   - Even: no match

3. **Center at index 2 ('b'):**
   - Odd: "aba" (length 3) (same length as current longest)
   - Even: no match

4. **Center at index 3 ('a'):**
   - Odd: "a" (length 1)
   - Even: no match

5. **Center at index 4 ('d'):**
   - Odd: "d" (length 1)
   - Even: N/A (end of string)

**Result: "bab"** (first palindrome of length 3 found)

### Why This is Optimal

- **Time Complexity O(n²)**: For each of n positions, we potentially expand up to n times in the worst case
- **Space Complexity O(1)**: We only use a constant amount of extra space for indices and temporary variables (the returned substring doesn't count toward space complexity)
- This is more efficient than dynamic programming approaches that use O(n²) space
- This is simpler to implement than Manacher's algorithm (O(n) time) while still being efficient for practical purposes

### Alternative Approaches

1. **Brute Force**: Check every substring O(n³) - For each of n² substrings, verify if palindrome in O(n)
2. **Dynamic Programming**: O(n²) time and O(n²) space - Build a 2D table where dp[i][j] indicates if substring [i,j] is palindrome
3. **Manacher's Algorithm**: O(n) time and O(n) space - Most optimal but complex to implement, uses preprocessing and dynamic programming with clever optimizations

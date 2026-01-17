# Palindromic Substrings

## Problem Description

Given a string `s`, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.

## Examples

**Example 1:**
```
Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
```

**Example 2:**
```
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
```

**Example 3:**
```
Input: s = "racecar"
Output: 10
Explanation: Palindromic strings: "r", "a", "c", "e", "c", "a", "r", "aca", "cec", "racecar".
```

## Constraints

- `1 <= s.length <= 1000`
- `s` consists of lowercase English letters.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Expand Around Center | None | O(n²) | O(1) |

### Approach: Expand Around Center

The optimal solution expands around each possible center to find all palindromic substrings. For each position in the string, we check both odd-length palindromes (single character center) and even-length palindromes (two character center).

**Key Insight:** A palindrome mirrors around its center. We can expand outward from each potential center and count palindromes as we go.

### Algorithm Steps

1. Initialize a counter to 0
2. For each index i in the string:
   - Count odd-length palindromes centered at i (expand from i, i)
   - Count even-length palindromes centered between i and i+1 (expand from i, i+1)
3. For each expansion:
   - Start with left and right pointers at the center
   - While left >= 0 and right < length and s[left] == s[right]:
     - Increment the counter (found a palindrome)
     - Move left pointer left and right pointer right
4. Return the total count

### Example Walkthrough

For `s = "aaa"`:

**Odd-length palindromes:**
1. Center at index 0 ('a'):
   - "a" (count = 1)
2. Center at index 1 ('a'):
   - "a" (count = 2)
   - "aaa" (count = 3)
3. Center at index 2 ('a'):
   - "a" (count = 4)

**Even-length palindromes:**
1. Center between index 0 and 1:
   - "aa" (count = 5)
2. Center between index 1 and 2:
   - "aa" (count = 6)

**Total: 6 palindromic substrings**

### Why This is Optimal

- **Time Complexity O(n²)**: For each of n positions, we potentially expand up to n times in the worst case (e.g., "aaaa...")
- **Space Complexity O(1)**: We only use a constant amount of extra space for pointers and counter
- This is optimal for counting palindromes because we need to examine each potential palindrome

### Alternative Approaches

1. **Brute Force**: Check every substring O(n³) - Too slow as we need to check each substring and verify if it's a palindrome
2. **Dynamic Programming**: O(n²) time and O(n²) space - Uses a 2D table to store whether substring [i,j] is a palindrome
3. **Manacher's Algorithm**: O(n) time but complex to implement and primarily for finding the longest palindrome, not counting all palindromes

# Minimum Window Substring

## Problem Description

Given two strings `s` and `t` of lengths `m` and `n` respectively, return the minimum window substring of `s` such that every character in `t` (including duplicates) is included in the window. If there is no such substring, return the empty string `""`.

The testcases will be generated such that the answer is unique.

## Examples

**Example 1:**
```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
```

**Example 2:**
```
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
```

**Example 3:**
```
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
```

## Constraints

- `m == s.length`
- `n == t.length`
- `1 <= m, n <= 10^5`
- `s` and `t` consist of uppercase and lowercase English letters.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Sliding Window | Hash Map | O(m + n) | O(m + n) |

### Approach: Sliding Window with Two Hash Maps

The optimal solution uses a sliding window with two hash maps: one to track the required character frequencies from `t`, and another to track the current window's character frequencies. We expand the window until all characters are matched, then contract it to find the minimum.

**Key Insight:** We need to find the smallest window in `s` that contains all characters from `t` with at least the required frequencies. We can track when we have a valid window by counting matched characters.

### Algorithm Steps

1. Create frequency map `need` for characters in `t`
2. Initialize `have` map for current window, `left = 0`, `formed = 0` (unique chars matched)
3. Track minimum window with `minLen = infinity`, `minLeft = 0`
4. Iterate with `right` pointer through `s`:
   - Add character to `have` map
   - If character frequency matches requirement, increment `formed`
   - While window is valid (`formed == required`):
     - Update minimum window if current is smaller
     - Remove leftmost character, update `have` and `formed`
     - Increment `left`
5. Return substring from `minLeft` to `minLeft + minLen`, or "" if no valid window

### Example Walkthrough

For `s = "ADOBECODEBANC"`, `t = "ABC"`:

need = {'A':1, 'B':1, 'C':1}, required = 3

1. Expand right until valid window found at "ADOBEC" (right=5)
2. Contract left: "DOBEC" invalid, "ADOBEC" is first valid window
3. Continue expanding: find "ODEBANC" (valid)
4. Contract: "DEBANC", "EBANC", "BANC" (valid, smaller)
5. Final answer: "BANC"

### Why This is Optimal

- **Time Complexity O(m + n)**: We process each character in `s` at most twice (once by right, once by left pointer), and we process `t` once to build the need map
- **Space Complexity O(m + n)**: We store frequency maps for characters in both strings
- This is optimal because we must examine every character to find a valid window

### Alternative Approaches (Not Optimal)

1. **Brute Force**: Check all substrings - O(m² × n) time (m² substrings, each takes O(n) to validate)
2. **Sliding Window with Array**: Use fixed-size array instead of map - Same complexity but slightly faster for ASCII

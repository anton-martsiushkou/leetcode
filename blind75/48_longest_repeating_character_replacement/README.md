# Longest Repeating Character Replacement

## Problem Description

You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

## Examples

**Example 1:**
```
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
```

**Example 2:**
```
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exist other ways to achieve this answer too.
```

## Constraints

- `1 <= s.length <= 10^5`
- `s` consists of only uppercase English letters.
- `0 <= k <= s.length`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Sliding Window | Hash Map/Array | O(n) | O(1) |

### Approach: Sliding Window with Character Frequency

The optimal solution uses a sliding window where we track the frequency of characters in the current window. The key insight is that a valid window can be formed if: `window_length - max_frequency <= k`, where `max_frequency` is the count of the most frequent character in the window.

**Key Insight:** We need to find the longest window where the number of characters that need to be replaced (window_length - max_frequency) is at most k.

### Algorithm Steps

1. Initialize `left = 0`, `maxCount = 0` (max frequency in current window), `maxLength = 0`
2. Create a frequency map for characters in the window
3. Iterate through the string with `right` pointer (0 to n-1):
   - Add current character to frequency map
   - Update `maxCount` to be the maximum frequency in current window
   - While window is invalid (`right - left + 1 - maxCount > k`):
     - Remove leftmost character from window
     - Increment `left`
   - Update `maxLength` with current valid window size
4. Return `maxLength`

### Example Walkthrough

For `s = "AABABBA"`, `k = 1`:

1. **right=0, char='A'**: freq={'A':1}, maxCount=1, window=1, valid (1-1=0<=1), max=1
2. **right=1, char='A'**: freq={'A':2}, maxCount=2, window=2, valid (2-2=0<=1), max=2
3. **right=2, char='B'**: freq={'A':2,'B':1}, maxCount=2, window=3, valid (3-2=1<=1), max=3
4. **right=3, char='A'**: freq={'A':3,'B':1}, maxCount=3, window=4, valid (4-3=1<=1), max=4
5. **right=4, char='B'**: freq={'A':3,'B':2}, maxCount=3, window=5, invalid (5-3=2>1)
   - Shrink: remove 'A', left=1, freq={'A':2,'B':2}, window=4, valid, max=4
6. **right=5, char='B'**: freq={'A':2,'B':3}, maxCount=3, window=5, invalid (5-3=2>1)
   - Shrink: remove 'A', left=2, freq={'A':1,'B':3}, window=4, valid, max=4
7. **right=6, char='A'**: freq={'A':2,'B':3}, maxCount=3, window=5, invalid (5-3=2>1)
   - Shrink: remove 'B', left=3, freq={'A':2,'B':2}, window=4, valid, max=4

Result: 4 (substring "BBBB" by replacing one 'A')

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the string once with right pointer, and left pointer moves at most n times total
- **Space Complexity O(1)**: We store at most 26 uppercase English letters in the frequency map
- This is optimal because we must examine every character at least once

### Alternative Approaches (Not Optimal)

1. **Brute Force**: Check all substrings and count replacements needed - O(n²) time
2. **Binary Search + Sliding Window**: Binary search on answer length - O(n log n) time, more complex

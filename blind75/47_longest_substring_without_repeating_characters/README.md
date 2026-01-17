# Longest Substring Without Repeating Characters

## Problem Description

Given a string `s`, find the length of the longest substring without repeating characters.

## Examples

**Example 1:**
```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

**Example 2:**
```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

**Example 3:**
```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## Constraints

- `0 <= s.length <= 5 * 10^4`
- `s` consists of English letters, digits, symbols and spaces.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Sliding Window | Hash Map/Set | O(n) | O(min(m, n)) |

### Approach: Sliding Window with Hash Map

The optimal solution uses a sliding window technique with a hash map to track the most recent index of each character. We maintain a window with no repeating characters and expand it as much as possible.

**Key Insight:** Instead of checking every possible substring (which would be O(n²) or O(n³)), we can use a sliding window that efficiently expands and contracts based on character repetitions.

### Algorithm Steps

1. Initialize `left = 0` (left pointer of window) and `maxLength = 0`
2. Create a hash map to store `{character: last_seen_index}`
3. Iterate through the string with `right` pointer (0 to n-1):
   - If current character exists in map and its index >= left:
     - Move `left` pointer to `map[char] + 1` (skip the repeated character)
   - Update the map with current character's index
   - Calculate current window length: `right - left + 1`
   - Update `maxLength` if current window is longer
4. Return `maxLength`

### Example Walkthrough

For `s = "abcabcbb"`:

1. **right=0, char='a'**: left=0, map={'a':0}, length=1, max=1
2. **right=1, char='b'**: left=0, map={'a':0,'b':1}, length=2, max=2
3. **right=2, char='c'**: left=0, map={'a':0,'b':1,'c':2}, length=3, max=3
4. **right=3, char='a'**: 'a' at index 0, move left=1, map={'a':3,'b':1,'c':2}, length=3, max=3
5. **right=4, char='b'**: 'b' at index 1, move left=2, map={'a':3,'b':4,'c':2}, length=3, max=3
6. **right=5, char='c'**: 'c' at index 2, move left=3, map={'a':3,'b':4,'c':5}, length=3, max=3
7. **right=6, char='b'**: 'b' at index 4, move left=5, map={'a':3,'b':6,'c':5}, length=2, max=3
8. **right=7, char='b'**: 'b' at index 6, move left=7, map={'a':3,'b':7,'c':5}, length=1, max=3

Result: 3

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the string once with the right pointer, and the left pointer only moves forward
- **Space Complexity O(min(m, n))**: Where m is the size of the character set. The hash map stores at most min(m, n) characters
- This is optimal because we must examine every character at least once, giving us a lower bound of O(n)

### Alternative Approaches (Not Optimal)

1. **Brute Force**: Check all substrings - O(n³) time (generating all substrings is O(n²) and checking uniqueness is O(n))
2. **Sliding Window with Set**: Similar approach but requires removal from set - O(2n) = O(n) time, but slower in practice

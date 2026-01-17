# Valid Anagram

## Problem Description

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Examples

**Example 1:**
```
Input: s = "anagram", t = "nagaram"
Output: true
```

**Example 2:**
```
Input: s = "rat", t = "car"
Output: false
```

## Constraints

- `1 <= s.length, t.length <= 5 * 10^4`
- `s` and `t` consist of lowercase English letters.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Hash Map (Frequency Counter) | Hash Map | O(n) | O(1) |

### Approach: Hash Map Frequency Counter

The optimal solution uses a hash map to count the frequency of each character in both strings. Two strings are anagrams if they have the same character frequencies.

**Key Insight:** Instead of sorting (which would be O(n log n)), we can use a hash map to count character frequencies in O(n) time. Since we only have lowercase English letters (26 characters), the space is O(1).

### Algorithm Steps

1. If lengths differ, return false immediately
2. Create frequency map for string `s`
3. For each character in string `t`:
   - Decrement its count in the frequency map
   - If count becomes negative or character doesn't exist, return false
4. If all characters processed successfully, return true

### Alternative Implementation

1. Count frequencies in both strings separately
2. Compare the two frequency maps for equality

### Example Walkthrough

For `s = "anagram"`, `t = "nagaram"`:

1. Length check: both have length 7
2. Build frequency map from s: {'a':3, 'n':1, 'g':1, 'r':1, 'm':1}
3. Process t:
   - 'n': freq['n'] = 0
   - 'a': freq['a'] = 2
   - 'g': freq['g'] = 0
   - 'a': freq['a'] = 1
   - 'r': freq['r'] = 0
   - 'a': freq['a'] = 0
   - 'm': freq['m'] = 0
4. All counts reached 0, return true

### Why This is Optimal

- **Time Complexity O(n)**: We traverse each string once, where n is the length of the strings
- **Space Complexity O(1)**: We store at most 26 lowercase English letters in the hash map, which is constant
- This is optimal because we must examine every character at least once to verify they match

### Alternative Approaches

1. **Sorting**: Sort both strings and compare - O(n log n) time, O(1) or O(n) space depending on sort implementation
2. **Array Counter**: Use fixed-size array[26] instead of hash map - Same complexity, slightly faster in practice

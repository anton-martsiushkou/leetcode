# Decode Ways

## Problem Description

A message containing letters from `A-Z` can be **encoded** into numbers using the following mapping:

```
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
```

To **decode** an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, `"11106"` can be mapped into:

- `"AAJF"` with the grouping `(1 1 10 6)`
- `"KJF"` with the grouping `(11 10 6)`

Note that the grouping `(1 11 06)` is invalid because `"06"` cannot be mapped into `'F'` since `"6"` is different from `"06"`.

Given a string `s` containing only digits, return the **number of ways** to **decode** it.

The test cases are generated so that the answer fits in a **32-bit** integer.

## Examples

**Example 1:**
```
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
```

**Example 2:**
```
Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
```

**Example 3:**
```
Input: s = "06"
Output: 0
Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").
```

## Constraints

- `1 <= s.length <= 100`
- `s` contains only digits and may contain leading zero(s).

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming | Variables | O(n) | O(1) |

### Approach: Dynamic Programming

The optimal solution uses dynamic programming to count the number of ways to decode the string. At each position, we can either decode one digit (if valid) or two digits (if valid).

**Key Insight:** For each position i, the number of ways to decode up to position i is:
- Ways to decode up to i-1 (if s[i] is valid: 1-9)
- Plus ways to decode up to i-2 (if s[i-1:i+1] is valid: 10-26)

This is similar to the Fibonacci sequence, where each position depends on the previous two positions.

### Algorithm Steps

1. **Edge Cases:**
   - If string is empty or starts with '0', return 0
   - If string has length 1 and is not '0', return 1

2. **DP Recurrence:**
   - Initialize prev2 = 1 (empty string has one way to decode)
   - Initialize prev1 = 1 if s[0] != '0', else 0
   - For each position i from 1 to n-1:
     - current = 0
     - If s[i] is '1'-'9': current += prev1 (can decode single digit)
     - If s[i-1:i+1] is '10'-'26': current += prev2 (can decode two digits)
     - Update prev2 = prev1, prev1 = current

3. **Return** prev1 (number of ways to decode entire string)

### Example Walkthrough

For `s = "226"`:

1. **Initialize:**
   - prev2 = 1 (empty string)
   - prev1 = 1 (s[0] = '2' is valid)

2. **i=1 (s[1] = '2'):**
   - Single digit '2' is valid: current = prev1 = 1
   - Two digits '22' is valid (10-26): current += prev2 = 1 + 1 = 2
   - Update: prev2 = 1, prev1 = 2

3. **i=2 (s[2] = '6'):**
   - Single digit '6' is valid: current = prev1 = 2
   - Two digits '26' is valid (10-26): current += prev2 = 2 + 1 = 3
   - Update: prev2 = 2, prev1 = 3

**Result:** 3 ways to decode

**The three decodings are:**
- "2" + "2" + "6" = "BBF"
- "2" + "26" = "BZ"
- "22" + "6" = "VF"

### Why This is Optimal

- **Time Complexity O(n)**: We iterate through the string once, making constant-time decisions at each position
- **Space Complexity O(1)**: We only need two variables to track the previous two states
- This is optimal because we must examine every character to determine the number of valid decodings, giving us a lower bound of O(n)

### Alternative Approaches

1. **Recursive (Brute Force)**: O(2^n) time - Too slow, explores all possible decodings
2. **DP with Array**: O(n) time, O(n) space - Same time complexity but uses more space
3. **Memoization**: O(n) time, O(n) space - Top-down approach with similar complexity

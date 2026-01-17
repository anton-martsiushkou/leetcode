# Encode and Decode Strings

## Problem Description

Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

Implement the `encode` and `decode` methods.

Note: The string may contain any possible characters out of 256 valid ASCII characters. Your algorithm should be generalized enough to work on any possible characters.

## Examples

**Example 1:**
```
Input: strs = ["Hello","World"]
Output: ["Hello","World"]
Explanation:
encode: "Hello" and "World" -> "5#Hello5#World"
decode: "5#Hello5#World" -> ["Hello", "World"]
```

**Example 2:**
```
Input: strs = [""]
Output: [""]
```

## Constraints

- `0 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` contains any possible ASCII characters.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Length Prefix | String Builder | O(n) | O(n) |

### Approach: Length Prefix Encoding

The optimal solution uses length-prefixed encoding where each string is prefixed with its length followed by a delimiter.

**Key Insight:** By encoding the length of each string before the string itself, we can unambiguously decode the result even if the strings contain any characters including the delimiter.

### Algorithm Steps

**Encode:**
1. For each string in the list:
   - Calculate its length
   - Append `length + "#" + string` to the result
2. Return the concatenated result

**Decode:**
1. Initialize an empty result list
2. Use a pointer to traverse the encoded string
3. For each encoded segment:
   - Read characters until we find '#' to get the length
   - Use the length to extract the exact substring
   - Add the substring to result list
   - Move pointer past the extracted string
4. Return the result list

### Example Walkthrough

For `strs = ["Hello", "World"]`:

**Encode:**
- "Hello" → length = 5 → "5#Hello"
- "World" → length = 5 → "5#World"
- Result: "5#Hello5#World"

**Decode "5#Hello5#World":**
1. Read "5" → length = 5
2. Skip "#"
3. Read 5 characters: "Hello"
4. Read "5" → length = 5
5. Skip "#"
6. Read 5 characters: "World"
7. Result: ["Hello", "World"]

### Why This is Optimal

- **Time Complexity O(n)**: Where n is the total length of all strings. We process each character exactly once for encoding and decoding.
- **Space Complexity O(n)**: For storing the encoded string and result
- This is optimal because we must process every character at least once

### Why This Works for Any Characters

The length prefix approach works even if strings contain:
- The delimiter character '#'
- Numbers
- Special characters
- Empty strings

This is because we know exactly how many characters to read after the delimiter.

### Edge Cases

1. **Empty string in list**: `[""]` → `"0#"`
2. **String with delimiter**: `["#"]` → `"1##"`
3. **Empty list**: `[]` → `""`
4. **String with numbers**: `["123"]` → `"3#123"`

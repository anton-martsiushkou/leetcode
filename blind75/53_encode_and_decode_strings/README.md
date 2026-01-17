# Encode and Decode Strings

## Problem Description

Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

Please implement `encode` and `decode` functions.

## Examples

**Example 1:**
```
Input: ["Hello","World"]
Output: ["Hello","World"]
Explanation:
encode: ["Hello","World"] -> "5#Hello5#World"
decode: "5#Hello5#World" -> ["Hello","World"]
```

**Example 2:**
```
Input: [""]
Output: [""]
Explanation:
encode: [""] -> "0#"
decode: "0#" -> [""]
```

**Example 3:**
```
Input: []
Output: []
Explanation:
encode: [] -> ""
decode: "" -> []
```

**Example 4:**
```
Input: ["a#b","c","de#f"]
Output: ["a#b","c","de#f"]
Explanation: The encoding must handle special characters like '#'
encode: ["a#b","c","de#f"] -> "3#a#b1#c4#de#f"
decode: "3#a#b1#c4#de#f" -> ["a#b","c","de#f"]
```

## Constraints

- `0 <= strs.length <= 1000`
- `0 <= strs[i].length <= 200`
- `strs[i]` contains any possible characters including special characters

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Length-Prefix Encoding | String Builder | O(n) | O(n) |

### Approach: Length-Prefix Encoding

The key challenge is to handle strings that may contain any characters, including delimiters. The optimal solution uses length-prefix encoding where each string is prefixed with its length followed by a delimiter.

**Key Insight:** By encoding the length of each string, we can safely decode strings containing any characters, including the delimiter itself.

### Algorithm Steps

**Encoding:**
1. For each string in the list:
   - Get the length of the string
   - Append `length + '#' + string` to the result
2. Return the concatenated result

**Decoding:**
1. Initialize an empty result list and index i = 0
2. While i < length of encoded string:
   - Find the next '#' delimiter starting from i
   - Extract the length (substring from i to '#')
   - Move i to position after '#'
   - Extract the string of that length
   - Add the extracted string to result
   - Move i forward by the extracted length
3. Return the result list

### Example Walkthrough

For `strs = ["Hello", "World"]`:

**Encoding:**
1. First string "Hello": length = 5
   - Append "5#Hello"
2. Second string "World": length = 5
   - Append "5#World"
3. Result: "5#Hello5#World"

**Decoding "5#Hello5#World":**
1. i = 0, find '#' at index 1
   - Length = 5 (substring [0:1])
   - Extract "Hello" (5 characters after '#')
   - i moves to 7
2. i = 7, find '#' at index 8
   - Length = 5 (substring [7:8])
   - Extract "World" (5 characters after '#')
   - i moves to 14
3. Result: ["Hello", "World"]

### Why This is Optimal

- **Time Complexity O(n)**: Where n is the total length of all strings. We process each character exactly once during encoding and decoding
- **Space Complexity O(n)**: We need to store the encoded string and decoded list
- This approach handles any characters including delimiters because we use length-prefix encoding
- The encoding is unambiguous - we always know exactly how many characters to read for each string

### Alternative Approaches

1. **Simple Delimiter**: Using a single character as delimiter (e.g., ',') fails when strings contain that character
2. **Escape Sequences**: Using escape characters is complex and error-prone with nested escaping
3. **JSON/Serialization**: Overkill for this problem and less efficient

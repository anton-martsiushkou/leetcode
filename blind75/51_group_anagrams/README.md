# Group Anagrams

## Problem Description

Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Examples

**Example 1:**
```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

**Example 2:**
```
Input: strs = [""]
Output: [[""]]
```

**Example 3:**
```
Input: strs = ["a"]
Output: [["a"]]
```

## Constraints

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Hash Map with Sorted Key | Hash Map | O(n × k log k) | O(n × k) |
| Hash Map with Character Count | Hash Map | O(n × k) | O(n × k) |

### Approach 1: Hash Map with Sorted Key

The most straightforward solution uses sorted strings as keys in a hash map. All anagrams will produce the same sorted string, so they'll be grouped together.

**Key Insight:** Anagrams are strings with the same characters in different orders. When sorted, all anagrams produce identical strings.

### Algorithm Steps (Sorted Key)

1. Create a hash map where key is sorted string, value is list of original strings
2. For each string in input:
   - Sort the string to create a key
   - Append the original string to the list for this key
3. Return all values (lists of anagrams) from the hash map

### Approach 2: Hash Map with Character Count (Optimal)

A more efficient approach uses character frequency as the key, avoiding the sorting step.

### Algorithm Steps (Character Count)

1. Create a hash map where key is character frequency tuple, value is list of strings
2. For each string in input:
   - Count character frequencies (e.g., [1,0,0,1,1,0,...,0] for "ate")
   - Use this count as key (convert to tuple for hashability)
   - Append the original string to the list for this key
3. Return all values from the hash map

### Example Walkthrough

For `strs = ["eat","tea","tan","ate","nat","bat"]`:

Using sorted key approach:
1. "eat" → sorted: "aet" → map["aet"] = ["eat"]
2. "tea" → sorted: "aet" → map["aet"] = ["eat", "tea"]
3. "tan" → sorted: "ant" → map["ant"] = ["tan"]
4. "ate" → sorted: "aet" → map["aet"] = ["eat", "tea", "ate"]
5. "nat" → sorted: "ant" → map["ant"] = ["tan", "nat"]
6. "bat" → sorted: "abt" → map["abt"] = ["bat"]

Result: [["eat","tea","ate"], ["tan","nat"], ["bat"]]

### Why This is Optimal

- **Time Complexity O(n × k log k)** for sorting approach: n strings, each of length k requires O(k log k) to sort
- **Time Complexity O(n × k)** for character count approach: n strings, each requires O(k) to count characters
- **Space Complexity O(n × k)**: Store all n strings of average length k
- Character count approach is optimal as we must examine every character

### Alternative Approaches (Not Optimal)

1. **Brute Force**: Compare every pair of strings - O(n² × k) time
2. **Prime Number Encoding**: Assign primes to each letter - O(n × k) but risk of overflow with long strings

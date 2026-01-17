# Word Break

## Problem Description

Given a string `s` and a dictionary of strings `wordDict`, return `true` if `s` can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

## Examples

**Example 1:**
```
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
```

**Example 2:**
```
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
```

**Example 3:**
```
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
```

## Constraints

- `1 <= s.length <= 300`
- `1 <= wordDict.length <= 1000`
- `1 <= wordDict[i].length <= 20`
- `s` and `wordDict[i]` consist of only lowercase English letters.
- All the strings of `wordDict` are unique.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Dynamic Programming | Array + Set | O(n² × m) | O(n + k) |

where n = length of s, m = average word length, k = total characters in wordDict

### Approach: Dynamic Programming with Word Dictionary Set

The optimal solution uses dynamic programming where `dp[i]` represents whether the substring `s[0...i-1]` can be segmented into dictionary words. We convert `wordDict` to a set for O(1) lookup, then for each position, we check all possible words ending at that position.

**Key Insight:** A string can be segmented up to position `i` if there exists a position `j < i` such that:
1. The substring `s[0...j-1]` can be segmented (dp[j] is true)
2. The substring `s[j...i-1]` is in the word dictionary

### Algorithm Steps

1. **Initialize:**
   - Convert `wordDict` to a set for O(1) lookup
   - Create boolean array `dp` of size `n + 1` where n = len(s)
   - Set `dp[0] = true` (empty string can always be segmented)

2. **Build DP Table:**
   - For each position `i` from 1 to n:
     - For each position `j` from 0 to i-1:
       - If `dp[j]` is true AND `s[j:i]` is in wordDict:
         - Set `dp[i] = true` and break

3. **Return Result:**
   - Return `dp[n]`

### Example Walkthrough

For `s = "leetcode"` and `wordDict = ["leet", "code"]`:

1. **Initialize:** dp = [T, F, F, F, F, F, F, F, F]
   - wordSet = {"leet", "code"}

2. **i = 4:**
   - j = 0: s[0:4] = "leet" in wordSet, dp[0] = true
   - dp[4] = true
   - dp = [T, F, F, F, T, F, F, F, F]

3. **i = 8:**
   - j = 4: s[4:8] = "code" in wordSet, dp[4] = true
   - dp[8] = true
   - dp = [T, F, F, F, T, F, F, F, T]

4. **Result:** dp[8] = true

### Example of False Case

For `s = "catsandog"` and `wordDict = ["cats", "dog", "sand", "and", "cat"]`:

The DP table will be: [T, F, F, T, F, F, F, F, F, F]
- Position 3: "cat" found
- No valid segmentation reaches the end
- Result: dp[9] = false

### Why This is Optimal

- **Time Complexity O(n² × m)**: For each of n positions, we check up to n previous positions, and each substring check takes O(m) where m is average word length. With set lookup, the actual complexity is closer to O(n² + n×w) where w is average word length.
- **Space Complexity O(n + k)**: We use O(n) for the dp array and O(k) for storing the word dictionary in a set, where k is total characters in wordDict
- This is optimal for the general case (trie optimization can improve to O(n² + k) in practice)

### Optimization with Trie

For very large word dictionaries, using a Trie can optimize the inner loop by avoiding checking all previous positions, but for the constraints given (wordDict.length ≤ 1000), the set-based approach is simpler and efficient enough.

### Alternative Approaches (Not Optimal)

1. **Recursive (No Memoization)**: O(2^n) time - Exponential due to trying all possible breaks
2. **BFS/DFS with Memoization**: O(n²) time, O(n) space - Same complexity but more overhead than DP

# Add and Search Word - Data Structure Design

## Problem Description

Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the `WordDictionary` class:

- `WordDictionary()` Initializes the object.
- `void addWord(word)` Adds `word` to the data structure, it can be matched later.
- `bool search(word)` Returns `true` if there is any string in the data structure that matches `word` or `false` otherwise. `word` may contain dots `'.'` where dots can be matched with any letter.

## Examples

**Example 1:**
```
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]

Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
```

## Constraints

- `1 <= word.length <= 25`
- `word` in `addWord` consists of lowercase English letters
- `word` in `search` consist of `'.'` or lowercase English letters
- There will be at most `2` dots in `word` for `search` queries
- At most `10^4` calls will be made to `addWord` and `search`

## Solution

| Operation | Time Complexity | Space Complexity |
|-----------|-----------------|------------------|
| AddWord | O(m) where m = word length | O(m) |
| Search (no dots) | O(m) where m = word length | O(1) |
| Search (with dots) | O(26^k × m) where k = number of dots | O(m) for recursion |

### Approach: Trie with Wildcard Search

This problem extends the basic Trie implementation to support wildcard matching with the `'.'` character, which can match any single letter.

**Key Insight:** Use a Trie for efficient prefix-based storage, and implement recursive backtracking for wildcard search. When encountering a `'.'`, try all possible child branches.

### Data Structure

Similar to a standard Trie:
- Each node contains a map of children (one for each character)
- A boolean flag marks complete words
- Search function uses recursion to handle wildcards

### Algorithm Steps

**AddWord(word):**
- Same as standard Trie insert
- Time: O(m) where m is word length

**Search(word):**
1. Use recursive helper function
2. For each character in word:
   - If character is `'.'`: recursively search all child branches
   - If character is a letter: search only that specific child
   - If no matching child exists, return false
3. At the end of word, check if current node is marked as end of word

### Example Walkthrough

After adding "bad", "dad", "mad":

```
        root
       / | \
      b  d  m
      |  |  |
      a  a  a
      |  |  |
      d  d  d (all marked as end)
```

**Search ".ad":**
1. First char is `'.'`: Try all children of root (b, d, m)
2. For each branch, continue with "ad"
3. All three branches have 'a' → 'd' → end marker
4. Return true

**Search "b..":**
1. First char is 'b': Move to 'b' node
2. Second char is `'.'`: Try all children of 'b' node ('a')
3. Third char is `'.'`: Try all children of 'a' node ('d')
4. Check if 'd' is end of word: Yes
5. Return true

### Why This is Optimal

- **AddWord O(m)**: Linear in word length, same as basic Trie
- **Search without dots O(m)**: Linear lookup, same as basic Trie
- **Search with dots O(26^k × m)**: In worst case with k dots, we might explore all 26 branches at each dot position. With the constraint of at most 2 dots, this is manageable.
- **Space O(N × M)**: N words of average length M

### Optimization Notes

1. The problem constrains dots to at most 2, preventing exponential blowup
2. Pruning occurs naturally: if a branch doesn't exist, backtracking stops immediately
3. For production use with many dots, consider adding length-based indexing to reduce search space

# Implement Trie (Prefix Tree)

## Problem Description

A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

- `Trie()` Initializes the trie object.
- `void insert(String word)` Inserts the string `word` into the trie.
- `boolean search(String word)` Returns `true` if the string `word` is in the trie (i.e., was inserted before), and `false` otherwise.
- `boolean startsWith(String prefix)` Returns `true` if there is a previously inserted string `word` that has the prefix `prefix`, and `false` otherwise.

## Examples

**Example 1:**
```
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]

Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True
```

## Constraints

- `1 <= word.length, prefix.length <= 2000`
- `word` and `prefix` consist only of lowercase English letters
- At most `3 * 10^4` calls in total will be made to `insert`, `search`, and `startsWith`

## Solution

| Operation | Time Complexity | Space Complexity |
|-----------|-----------------|------------------|
| Insert | O(m) where m = word length | O(m) |
| Search | O(m) where m = word length | O(1) |
| StartsWith | O(m) where m = prefix length | O(1) |

### Approach: Trie Data Structure

A trie is a tree-like data structure where:
- Each node represents a single character
- Paths from root to nodes represent prefixes
- A special marker indicates complete words

**Key Insight:** By organizing strings as a tree of characters, we can share common prefixes between multiple words, enabling efficient prefix-based operations.

### Data Structure

Each Trie node contains:
- An array or map of child nodes (one for each possible character)
- A boolean flag `isEndOfWord` to mark complete words

```
Example: Insert "apple", "app", "application"

        root
         |
         a
         |
         p
         |
         p (isEnd=true for "app")
        / \
       l   l
       |   |
       e   i
       |   |
    (isEnd) c
       |   |
            a
            |
            t
            |
            i
            |
            o
            |
            n (isEnd=true)

"apple" is marked complete at 'e'
"app" is marked complete at second 'p'
```

### Algorithm Steps

**Insert(word):**
1. Start at root node
2. For each character in word:
   - If child node for character doesn't exist, create it
   - Move to child node
3. Mark final node as end of word

**Search(word):**
1. Start at root node
2. For each character in word:
   - If child node for character doesn't exist, return false
   - Move to child node
3. Return true only if final node is marked as end of word

**StartsWith(prefix):**
1. Start at root node
2. For each character in prefix:
   - If child node for character doesn't exist, return false
   - Move to child node
3. Return true (don't check end of word marker)

### Why This is Optimal

- **Insert O(m)**: We traverse/create exactly m nodes for a word of length m
- **Search O(m)**: We traverse at most m nodes
- **StartsWith O(m)**: We traverse at most m nodes
- **Space O(ALPHABET_SIZE × N × M)** where N is number of words and M is average length
  - In practice, common prefixes are shared, significantly reducing space

### Advantages of Trie

1. **Fast prefix queries**: O(m) vs O(n×m) for linear search
2. **Memory efficient for common prefixes**: "app", "apple", "application" share "app" nodes
3. **Enables autocomplete**: StartsWith gives all words with a prefix
4. **Lexicographic ordering**: In-order traversal yields sorted words

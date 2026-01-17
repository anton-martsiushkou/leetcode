# Subtree of Another Tree

## Problem Description

Given the roots of two binary trees `root` and `subRoot`, return `true` if there is a subtree of `root` with the same structure and node values of `subRoot` and `false` otherwise.

A subtree of a binary tree `tree` is a tree that consists of a node in `tree` and all of this node's descendants. The tree `tree` could also be considered as a subtree of itself.

## Examples

**Example 1:**
```
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
Explanation:
Main tree:        Subtree:
      3              4
     / \            / \
    4   5          1   2
   / \
  1   2
```

**Example 2:**
```
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
Explanation:
Main tree:        Subtree:
      3              4
     / \            / \
    4   5          1   2
   / \
  1   2
 /
0
The subtree has extra node 0, so not a match.
```

**Example 3:**
```
Input: root = [1,1], subRoot = [1]
Output: true
```

## Constraints

- The number of nodes in the `root` tree is in the range `[1, 2000]`
- The number of nodes in the `subRoot` tree is in the range `[1, 1000]`
- `-10^4 <= root.val <= 10^4`
- `-10^4 <= subRoot.val <= 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS Recursive | Binary Tree | O(m * n) | O(h) |
| Tree Serialization | Binary Tree + String | O(m + n) | O(m + n) |

### Approach 1: DFS Recursive (Recommended)

The recursive approach traverses the main tree and at each node checks if the subtree rooted at that node matches the subRoot tree.

**Key Insight:** A subtree match requires two checks:
1. Are the current trees identical? (same structure and values)
2. If not, does the subtree exist in left or right child?

### Algorithm Steps

1. Define a helper function `isSameTree(p, q)` that checks if two trees are identical
2. Main function `isSubtree(root, subRoot)`:
   - Base case: If root is `null`, return `false` (can't find subtree in empty tree)
   - Check if current tree matches subRoot using `isSameTree`
   - If yes, return `true`
   - Otherwise, recursively check left and right subtrees
   - Return `true` if found in either subtree

### Example Walkthrough

For root = `[3,4,5,1,2]` and subRoot = `[4,1,2]`:

```
isSubtree(3, [4,1,2]):
  ├─ isSameTree(3, 4) = false (different values)
  ├─ isSubtree(4, [4,1,2]):
  │    ├─ isSameTree(4, 4):
  │    │    ├─ values match (4 == 4) ✓
  │    │    ├─ left subtrees match (1 == 1) ✓
  │    │    └─ right subtrees match (2 == 2) ✓
  │    │    return true
  │    return true ← Found match!
  return true
```

### Why This is Optimal (for simplicity)

- **Time Complexity O(m * n)**: In the worst case, we check if every node in the main tree (m nodes) is the root of a matching subtree, and each check takes O(n) time
  - m = number of nodes in main tree
  - n = number of nodes in subRoot tree
- **Space Complexity O(h)**: Recursion stack depth equals height of main tree
- This approach is simple and intuitive, though not asymptotically optimal

### Approach 2: Tree Serialization (More Efficient)

Serialize both trees to strings and check if the subRoot serialization is a substring of the root serialization.

**Steps:**
1. Serialize both trees using preorder traversal with unique markers
2. Check if subRoot serialization is a substring of root serialization
3. Use KMP or similar algorithm for O(m + n) substring search

**Advantages:**
- Time complexity: O(m + n)
- More efficient for large trees

**Disadvantages:**
- More complex implementation
- Higher space complexity: O(m + n)
- Requires careful handling of edge cases (e.g., using markers like "#" for values and "null" for children)

### Key Points

1. **Tree equality**: Two trees are identical if they have the same structure and node values
2. **Subtree vs. subset**: A subtree must include ALL descendants, not just a subset of nodes
3. **Edge cases**: Handle null trees, single nodes, and trees where subRoot equals root

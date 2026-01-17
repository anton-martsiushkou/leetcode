# Same Tree

## Problem Description

Given the roots of two binary trees `p` and `q`, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

## Examples

**Example 1:**
```
Input: p = [1,2,3], q = [1,2,3]
Output: true
Explanation: Both trees are:
    1           1
   / \         / \
  2   3       2   3
```

**Example 2:**
```
Input: p = [1,2], q = [1,null,2]
Output: false
Explanation:
    1           1
   /             \
  2               2
```

**Example 3:**
```
Input: p = [1,2,1], q = [1,1,2]
Output: false
Explanation:
    1           1
   / \         / \
  2   1       1   2
```

## Constraints

- The number of nodes in both trees is in the range `[0, 100]`
- `-10^4 <= Node.val <= 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS Recursive | Binary Tree | O(n) | O(h) |
| BFS Iterative | Binary Tree + Queue | O(n) | O(w) |

### Approach: DFS Recursive (Recommended)

The recursive approach leverages the natural recursive structure of trees. Two trees are identical if:
1. Both roots have the same value
2. Their left subtrees are identical
3. Their right subtrees are identical

**Key Insight:** The problem has a simple recursive definition - trees are the same if their roots match and their corresponding subtrees are the same.

### Algorithm Steps

1. Base cases:
   - If both nodes are `null`, return `true` (both empty trees are identical)
   - If one node is `null` and the other isn't, return `false` (different structure)
   - If both nodes exist but have different values, return `false`
2. Recursively check if left subtrees are the same
3. Recursively check if right subtrees are the same
4. Return `true` only if both subtree comparisons return `true`

### Example Walkthrough

For trees p = `[1,2,3]` and q = `[1,2,3]`:

```
isSameTree(1, 1):
  ├─ values match (1 == 1) ✓
  ├─ isSameTree(2, 2):
  │    ├─ values match (2 == 2) ✓
  │    ├─ isSameTree(null, null) = true ✓
  │    └─ isSameTree(null, null) = true ✓
  │    return true
  └─ isSameTree(3, 3):
       ├─ values match (3 == 3) ✓
       ├─ isSameTree(null, null) = true ✓
       └─ isSameTree(null, null) = true ✓
       return true
  return true
```

### Why This is Optimal

- **Time Complexity O(n)**: We visit each node in both trees exactly once, where n is the minimum number of nodes in the two trees
- **Space Complexity O(h)**: The recursion stack depth equals the height of the tree
  - Best case (balanced tree): O(log n)
  - Worst case (skewed tree): O(n)
- The recursive solution is simple, elegant, and naturally matches the tree comparison logic

### Alternative Approaches

1. **BFS Iterative**: Use two queues to traverse both trees level by level
   - Same time complexity O(n)
   - Space complexity O(w) where w is maximum width
   - More code but avoids recursion
